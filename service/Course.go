package service

import (
	"database/sql"
	"encoding/json"
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

func (c *CourseService) GetCourseByCode(code string) (json.RawMessage, error) {
	var course model.Course
	db := database.DB

	if err := db.Where("code = ?", code).First(&course).Error; err != nil {
		return nil, err
	}

	var t sql.NullString
	if err := db.Model(&model.Course{}).Where("id = ?", course.ID).Select("hours_total").Scan(&t).Error; err != nil {
		return nil, err
	}
	course.HoursTotal = new(interface{})
	if t.Valid {
		*course.HoursTotal = t.String
	} else {
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

	result, err := json.Marshal(course)
	if err != nil {
		return nil, err
	}

	return result, nil
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

func (c *CourseService) GetCourseList(page *model.Page, r *model.Course, courses *[]model.Course) error {
	db := database.DB

	if r.Code != nil {
		db = db.Where("code LIKE ?", "%"+*r.Code+"%")
	}
	if r.Name != nil {
		db = db.Where("name LIKE ?", "%"+*r.Name+"%")
	}
	if r.ForeignName != nil {
		db = db.Where("foreign_name LIKE ?", "%"+*r.ForeignName+"%")
	}
	if r.Remark != nil {
		db = db.Where("remark LIKE ?", "%"+*r.Remark+"%")
	}
	if r.ShowRemark != nil {
		db = db.Where("show_remark LIKE ?", "%"+*r.ShowRemark+"%")
	}
	if r.DepartmentName != nil {
		db = db.Where("department_name LIKE ?", "%"+*r.DepartmentName+"%")
	}
	if r.LeaderName != nil {
		db = db.Where("leader_name LIKE ?", "%"+*r.LeaderName+"%")
	}
	if r.Assessment != nil {
		db = db.Where("assessment = ?", *r.Assessment)
	}
	if r.Credit != nil {
		db = db.Where("credit = ?", *r.Credit)
	}

	if err := db.Model(&model.Course{}).Count(&page.Total).Error; err != nil {
		return err
	}

	if err := db.Offset((page.Current - 1) * page.PageSize).Limit(page.PageSize).Find(courses).Error; err != nil {
		return err
	}

	for i := 0; i < len(*courses); i++ {
		db := database.DB
		var t sql.NullString
		if err := db.Model(&model.Course{}).Where("id = ?", (*courses)[i].ID).Select("hours_total").Scan(&t).Error; err != nil {
			return err
		}
		(*courses)[i].HoursTotal = new(interface{})
		if t.Valid {
			*(*courses)[i].HoursTotal = t.String
		} else {
			var sum int
			if (*courses)[i].HoursLecture != nil {
				sum += *(*courses)[i].HoursLecture
			}
			if (*courses)[i].HoursPractices != nil {
				sum += *(*courses)[i].HoursPractices
			}
			if (*courses)[i].HoursExperiment != nil {
				sum += *(*courses)[i].HoursExperiment
			}
			if (*courses)[i].HoursComputer != nil {
				sum += *(*courses)[i].HoursComputer
			}
			if (*courses)[i].HoursSelf != nil {
				sum += *(*courses)[i].HoursSelf
			}
			*(*courses)[i].HoursTotal = sum
		}
	}

	return nil
}
