package service

import (
	"errors"
	"strconv"

	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/model"
)

type CourseService struct{}

func (c *CourseService) CreateCourse(course *model.Course) (uint, error) {
	db := database.DB

	t, ok := course.HoursTotal.(int)
	if ok {
		course.HoursTotal = strconv.Itoa(t)
	}

	_, ok = course.HoursTotal.(string)
	if !ok {
		return 0, errors.New("errTypeParse")
	}

	if err := db.Create(course).Error; err != nil {
		return 0, err
	}

	return course.ID, nil
}

func (c *CourseService) DeleteCourse(id uint) error {
	db := database.DB

	if err := db.Delete(&model.Course{}, id).Error; err != nil {
		return err
	}

	return nil
}
