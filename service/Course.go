package service

import (
	"errors"

	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/model"
)

type CourseService struct{}

func (c *CourseService) CreateCourse(course *model.Course) (uint, error) {
	db := database.DB

	_, ok := (*course.HoursTotal).(int)
	if ok {
		*course.HoursTotal = nil
	}

	if err := db.Create(course).Error; err != nil {
		return 0, err
	}

	return *course.ID, nil
}

func (c *CourseService) DeleteCourse(id uint) error {
	db := database.DB

	if err := db.Delete(&model.Course{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (c *CourseService) GetCourse(id uint, course *model.Course) error {
	db := database.DB

	if err := db.Where("id = ?", id).First(course).Error; err != nil {
		return err
	}

	if course.HoursTotal == nil {
		var sum int
		if course.HoursLecture != nil {
			sum += *course.HoursLecture
		}
		if course.HoursPractices != nil {
			sum += *course.HoursPractices
		}
		if course.HoursExperiment != nil {
			sum += *course.HoursExperiment
		}
		if course.HoursComputer != nil {
			sum += *course.HoursComputer
		}
		if course.HoursSelf != nil {
			sum += *course.HoursSelf
		}
		*course.HoursTotal = sum
	}

	return nil
}

func (c *CourseService) UpdateCourse(id uint, course *model.Course) error {
	db := database.DB

	var old_course model.Course
	if err := db.Where("id = ?", id).First(&old_course).Error; err != nil {
		return err
	}

	tx := db.Begin()

	if course.Code != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("code", course.Code).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.Name != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("name", course.Name).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.ForeignName != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("foreign_name", course.ForeignName).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.Credit != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("credit", course.Credit).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if *course.HoursTotal != nil {
		t, ok := (*course.HoursTotal).(int)
		if ok {
			var sum int
			if course.HoursLecture != nil {
				sum += *course.HoursLecture
			}
			if course.HoursPractices != nil {
				sum += *course.HoursPractices
			}
			if course.HoursExperiment != nil {
				sum += *course.HoursExperiment
			}
			if course.HoursComputer != nil {
				sum += *course.HoursComputer
			}
			if course.HoursSelf != nil {
				sum += *course.HoursSelf
			}
			if t != sum {
				tx.Rollback()
				return errors.New("errHoursTotal")
			}

			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_total", course.HoursTotal).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_lecture", course.HoursLecture).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_practices", course.HoursPractices).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_experiment", course.HoursExperiment).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_computer", course.HoursComputer).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_self", course.HoursSelf).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if !(course.HoursLecture == nil && course.HoursPractices == nil && course.HoursExperiment == nil &&
				course.HoursComputer == nil && course.HoursSelf == nil) {
				tx.Rollback()
				return errors.New("errHoursTotal")
			}

			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_total", course.HoursTotal).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_lecture", nil).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_practices", nil).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_experiment", nil).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_computer", nil).Error; err != nil {
				tx.Rollback()
				return err
			}
			if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("hours_self", nil).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if course.Assessment != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("assessment", course.Assessment).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.ShowRemark != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("show_remark", course.ShowRemark).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.Remark != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("remark", course.Remark).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.DepartmentName != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("department_name", course.DepartmentName).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if course.LeaderName != nil {
		if err := tx.Model(&model.Course{}).Where("id = ?", id).Update("leader_name", course.LeaderName).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
