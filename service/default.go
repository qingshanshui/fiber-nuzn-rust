package service

import (
	"fiber-layout/models"
	v1 "fiber-layout/validator/form"
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
