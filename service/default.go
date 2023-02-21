package service

import (
	"fiber-nuzn-rust/models"
	v1 "fiber-nuzn-rust/validator/form"
)

type Default struct {
}

func NewDefaultService() *Default {
	return &Default{}
}

func (t *Default) List() ([]models.Course, error) {
	list, err := models.NewCourse().List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Default) Category(c v1.CategoryRequest) (*models.Course, error) {
	list, err := models.NewCourse().Category(c.ID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
