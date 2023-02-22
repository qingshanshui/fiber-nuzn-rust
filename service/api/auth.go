package apiService

import (
	"context"
	"errors"
	"fiber-nuzn-rust/initalize"
	"fiber-nuzn-rust/models"
	"fiber-nuzn-rust/pkg/utils"
	apiForm "fiber-nuzn-rust/validator/form/api"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Auth struct {
}

func NewAuthService() *Auth {
	return &Auth{}
}

// 发送验证码
func (t *Auth) Code(c apiForm.CodeRequest) error {

	// 生成 验证码 存入 redis 1分钟
	r := utils.RandString(8)
	if err := initalize.Rdb.Set(context.Background(), c.Email, r, time.Minute*1).Err(); err != nil {
		return err
	}
	// 发送验证码
	if err := initalize.Email.Send(c.Email, "<h1>来自Rust中文网验证码:"+r+"</h1>"); err != nil {
		return errors.New("邮箱验证码发送失败")
	}
	return nil
}

func (t *Auth) Login(c apiForm.LoginRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Email)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Auth) Register(c apiForm.RegisterRequest) error {
	// 取redis 验证码 对比传入验证码
	val, rdberr := initalize.Rdb.Get(context.Background(), c.Email).Result()
	if rdberr != nil {
		return rdberr
	}
	if val != c.Code {
		return errors.New("验证码错误")
	}
	uid, _ := gonanoid.New()
	u := models.NewUser()
	u.NickName = c.NickName
	u.RegisterSource = 2
	u.CreatedAt = time.Now()
	err := u.Register(uid, c.Email)
	fmt.Println(err, "===========")
	if err != nil {
		return err
	}
	return nil
}
