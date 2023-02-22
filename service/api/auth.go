package apiService

import (
	"errors"
	"fiber-nuzn-rust/initalize"
	"fiber-nuzn-rust/models"
	apiForm "fiber-nuzn-rust/validator/form/api"
)

type Auth struct {
}

func NewAuthService() *Auth {
	return &Auth{}
}

// 发送验证码
func (t *Auth) Code(c apiForm.CodeRequest) error {
	if err := initalize.Email.Send(c.Email, "<h1>来自Rust中文网验证码:654230</h1>"); err != nil {
		return errors.New("邮箱验证码发送失败")
	}
	return nil
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
