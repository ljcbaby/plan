package service

import (
	"encoding/json"

	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/model"
)

type ProgramService struct{}

func (s *ProgramService) CreateProgram(program *model.Program) (uint, error) {
	var Content model.Node
	Content.ID = new(string)
	*Content.ID = "root"
	Content.Title = new(model.Title)
	Content.Title.Type = new(string)
	*Content.Title.Type = "node"
	Content.Title.Name = new(string)
	*Content.Title.Name = *program.Name
	Content.Content = new([]model.Node)
	*Content.Content = []model.Node{}

	program.Content = new(json.RawMessage)
	*program.Content, _ = json.Marshal(Content)

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

		if *node.Title.Type == "course" && node.Title.CourseCode != nil {
			cs := &CourseService{}

			node.Title.Course = new(json.RawMessage)
			*node.Title.Course, err = cs.GetCourseByCode(*node.Title.CourseCode)
			if err != nil {
				if err.Error() == "record not found" {
					*node.Title.Course = json.RawMessage("{\"name\":\"课程已删除\",\"foreignName\":\"\",\"credit\":0,\"assessment\":\"\",\"departmentName\":\"\",\"leaderName\":\"\"}")
				} else {
					return err
				}
			}

			continue
		}

		if *node.Title.Type == "node" && node.Content != nil {
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

	var content model.Node
	if err := json.Unmarshal(*program.Content, &content); err != nil {
		return err
	}

	queue := []*model.Node{&content}
	seenTags := make(map[string]bool)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Title != nil && node.Title.Tags != nil {
			for _, tag := range *node.Title.Tags {
				if !seenTags[tag] {
					*program.Tags = append(*program.Tags, tag)
					seenTags[tag] = true
				}
			}
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
