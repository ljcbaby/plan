package service

import (
	"encoding/json"
	"sort"

	"github.com/gin-gonic/gin"
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
					courseJSON, _ := json.Marshal(gin.H{
						"name":           "课程已删除",
						"foreignName":    "",
						"credit":         0,
						"assessment":     "",
						"departmentName": "",
						"leaderName":     "",
					})
					*node.Title.Course = json.RawMessage(courseJSON)
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

	var dfs func(node *model.Node) float64
	dfs = func(node *model.Node) float64 {
		if node.Title != nil && node.Title.Type != nil && *node.Title.Type == "course" {
			var credit float64 = 0

			if node.Title.Course != nil {
				var course model.Course
				if err := json.Unmarshal(*node.Title.Course, &course); err != nil {
					return 0
				}

				if course.Credit != nil {
					credit = *course.Credit
				}
			}

			return credit
		}

		if node.Title != nil && node.Title.Type != nil && *node.Title.Type == "node" {
			if node.Title.Requirement != nil && node.Title.Requirement.MinCredit != nil {
				if node.Content != nil {
					var credit float64 = 0
					for i := range *node.Content {
						credit += dfs(&(*node.Content)[i])
					}
					if credit < *node.Title.Requirement.MinCredit {
						return -1
					}
				}

				return *node.Title.Requirement.MinCredit
			}

			var allCredit float64 = 0

			if node.Content != nil {
				if node.Title.Requirement != nil && node.Title.Requirement.MinCourse != nil {
					var l []float64
					for i := range *node.Content {
						l = append(l, dfs(&(*node.Content)[i]))
					}
					if len(l) < *node.Title.Requirement.MinCourse {
						allCredit = -1
					} else {
						sort.Float64s(l)
						for i := 0; i < *node.Title.Requirement.MinCourse; i++ {
							allCredit += l[i]
						}
					}
				} else {
					for i := range *node.Content {
						allCredit += dfs(&(*node.Content)[i])
					}
				}
			}

			node.Title.AllCredit = new(float64)
			*node.Title.AllCredit = allCredit
		}

		return *node.Title.AllCredit
	}

	dfs(&content)

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

func (s *ProgramService) CalculateProgram(id uint, tags *[]string, credit *float64, hours *int) error {
	var program model.Program
	if err := s.GetProgramWithContent(id, &program); err != nil {
		return err
	}

	*credit = 0
	*hours = 0
	tag := (*tags)[0]

	var dfs func(node *model.Node)
	dfs = func(node *model.Node) {
		if node.Title != nil && node.Title.Type != nil && *node.Title.Type == "course" {
			var flag bool = false

			if node.Title.Tags != nil {
				for _, t := range *node.Title.Tags {
					if t == tag {
						flag = true
						break
					}
				}
			}

			if flag && node.Title.Course != nil {
				var course model.Course
				if err := json.Unmarshal(*node.Title.Course, &course); err != nil {
					return
				}

				if course.Credit != nil {
					*credit += *course.Credit
				}
			}

			return
		}

		if node.Title != nil && node.Title.Type != nil && *node.Title.Type == "node" {
			var flag bool = false

			if node.Title.Tags != nil {
				for _, t := range *node.Title.Tags {
					if t == tag {
						flag = true
						break
					}
				}
			}

			if flag {
				if node.Title.AllCredit != nil {
					*credit += *node.Title.AllCredit
				} else {
					*credit += *node.Title.Requirement.MinCredit
				}
				return
			}

			if node.Content != nil {
				for i := range *node.Content {
					dfs(&(*node.Content)[i])
				}
			}
		}
	}

	var content model.Node
	if err := json.Unmarshal(*program.Content, &content); err != nil {
		return err
	}

	dfs(&content)

	return nil
}

func (s *ProgramService) UpdateProgram(id uint, program *model.Program) error {
	db := database.DB

	var oldProgram model.Program
	if err := db.Where("id = ?", id).First(&oldProgram).Error; err != nil {
		return err
	}

	tx := db.Begin()

	if program.Name != nil {
		if err := tx.Model(&oldProgram).Update("name", program.Name).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if program.Major != nil {
		if err := tx.Model(&oldProgram).Update("major", program.Major).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if program.Department != nil {
		if err := tx.Model(&oldProgram).Update("department", program.Department).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if program.DependencyID != nil {
		if err := tx.Model(&oldProgram).Update("dependency_id", program.DependencyID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if program.Grade != nil {
		if err := tx.Model(&oldProgram).Update("grade", program.Grade).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if program.Content != nil {
		queue := []*model.Node{}

		var content model.Node
		if err := json.Unmarshal(*program.Content, &content); err != nil {
			tx.Rollback()
			return err
		}
		queue = append(queue, &content)

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Title != nil && node.Title.Course != nil {
				node.Title.Course = nil
			}

			if node.Title != nil && node.Title.AllCredit != nil {
				node.Title.AllCredit = nil
			}

			if node.Content != nil {
				for i := range *node.Content {
					queue = append(queue, &(*node.Content)[i])
				}
			}
		}

		*program.Content, _ = json.Marshal(content)

		if err := tx.Model(&oldProgram).Update("content", program.Content).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
