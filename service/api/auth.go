package apiService

import (
	"fiber-nuzn-rust/models"
	apiForm "fiber-nuzn-rust/validator/form/api"
)

type Auth struct {
}

func NewAuthService() *Auth {
	return &Auth{}
}

func (t *Auth) Login(c apiForm.LoginRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Auth) Register(c apiForm.RegisterRequest) ([]models.Course, error) {
	list, err := models.NewCourse().List()
	if err != nil {
		return nil, err
	}
	return list, nil
}
