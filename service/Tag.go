package service

import (
	"github.com/ljcbaby/plan/database"
	"github.com/ljcbaby/plan/model"
)

type TagService struct{}

func (s *TagService) CreateTag(tag *model.Tag) error {
	db := database.DB

	if err := db.Create(tag).Error; err != nil {
		return err
	}

	return nil
}

func (s *TagService) GetTagList() ([]model.Tag, error) {
	db := database.DB

	var tags []model.Tag
	if err := db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *TagService) DeleteTag(id uint64) error {
	db := database.DB

	if err := db.Delete(&model.Tag{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (s *TagService) UpdateTag(tag *model.Tag) error {
	db := database.DB

	if err := db.Save(tag).Error; err != nil {
		return err
	}

	return nil
}
