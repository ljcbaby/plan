package service

import (
	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/model"
)

type TagService struct{}

func (s *TagService) GetTags(tags *model.Tag) error {
	db := database.DB

	if err := db.Where("id = ?", 1).First(tags).Error; err != nil {
		return err
	}

	return nil
}

func (s *TagService) UpdateTags(tags *model.Tag) error {
	db := database.DB

	if err := db.Table("tags").Where("id = ?", 1).Updates(tags).Error; err != nil {
		return err
	}

	return nil
}
