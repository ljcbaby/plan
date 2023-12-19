package service

import (
	"encoding/json"

	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/model"
)

type ProgramService struct{}

func (s *ProgramService) CreateProgram(program *model.Program) (uint, error) {
	content := json.RawMessage("{}")
	program.Content = &content

	db := database.DB

	if err := db.Create(program).Error; err != nil {
		return 0, err
	}

	return *program.ID, nil
}

func (s *ProgramService) DeleteProgram(id uint) error {
	db := database.DB

	if err := db.Delete(&model.Program{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (s *ProgramService) GetProgram(id uint, program *model.Program) error {
	db := database.DB

	if err := db.Where("id = ?", id).First(program).Error; err != nil {
		return err
	}

	err := s.GenerateProgramTags(program)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProgramService) GetProgramWithNoContent(id uint, program *model.Program) error {
	err := s.GetProgram(id, program)
	if err != nil {
		return err
	}

	program.Content = nil

	return nil
}

func (s *ProgramService) GetProgramWithContent(id uint, program *model.Program) error {
	err := s.GetProgram(id, program)
	if err != nil {
		return err
	}

	queue := []*model.Node{}

	var content model.Node
	if err := json.Unmarshal(*program.Content, &content); err != nil {
		return err
	}
	queue = append(queue, &content)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if *node.Type == "course" && node.CourseCode != nil {
			cs := &CourseService{}

			node.Course = new(json.RawMessage)
			*node.Course, err = cs.GetCourseByCode(*node.CourseCode)
			if err != nil {
				return err
			}

			continue
		}

		if *node.Type == "node" && node.Content != nil {
			for i := range *node.Content {
				queue = append(queue, &(*node.Content)[i])
			}

			continue
		}
	}

	*program.Content, err = json.Marshal(content)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProgramService) GenerateProgramTags(program *model.Program) error {
	program.Tags = &[]string{}

	queue := []*model.Node{}

	var content model.Node
	if err := json.Unmarshal(*program.Content, &content); err != nil {
		return err
	}
	queue = append(queue, &content)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Tags != nil {
			*program.Tags = append(*program.Tags, *node.Tags...)
		}

		if node.Content != nil {
			for i := range *node.Content {
				queue = append(queue, &(*node.Content)[i])
			}
		}
	}

	return nil
}

func (s *ProgramService) GetProgramList(programs *[]model.Program) error {
	db := database.DB

	if err := db.Find(programs).Error; err != nil {
		return err
	}

	for i := range *programs {
		err := s.GenerateProgramTags(&(*programs)[i])
		if err != nil {
			return err
		}

		(*programs)[i].Content = nil
	}

	return nil
}
