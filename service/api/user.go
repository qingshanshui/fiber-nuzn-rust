package apiService

import (
	"fiber-nuzn-rust/models"
	apiForm "fiber-nuzn-rust/validator/form/api"
)

type User struct {
}

func NewUserService() *User {
	return &User{}
}

func (t *User) List(c apiForm.UserInfoRequest) (*models.Course, error) {
	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}
