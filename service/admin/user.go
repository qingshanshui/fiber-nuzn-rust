package adminService

import (
	"fiber-nuzn-rust/models"
	adminForm "fiber-nuzn-rust/validator/form/admin"
)

type User struct {
}

func NewUserService() *User {
	return &User{}
}

func (t *User) List(c adminForm.UserInfoRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}
