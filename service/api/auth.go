package apiService

import (
	"context"
	"errors"
	"fiber-nuzn-rust/initalize"
	"fiber-nuzn-rust/models"
	"fiber-nuzn-rust/pkg/utils"
	apiForm "fiber-nuzn-rust/validator/form/api"
	"github.com/spf13/viper"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Auth struct {
}

func NewAuthService() *Auth {
	return &Auth{}
}

// Code 发送验证码
func (t *Auth) Code(c apiForm.CodeRequest) error {

	// 判断 邮箱是否已经注册
	//i, errua := models.NewUserAuth().UsernameIsMl("2", c.Email)
	//if errua != nil {
	//	return errua
	//}
	//if len(i) != 0 {
	//	return errors.New("邮箱已注册")
	//}

	// 生成 验证码 存入 redis 1分钟
	r := utils.RandString(8)
	if err := initalize.Rdb.Set(context.Background(), c.Email, r, time.Minute*100).Err(); err != nil {
		return err
	}
	// 发送验证码
	if err := initalize.Email.Send(c.Email, "<h1>来自Rust中文网验证码:"+r+"</h1>"); err != nil {
		return errors.New("邮箱验证码发送失败")
	}
	return nil
}

// Register 注册
func (t *Auth) Register(c apiForm.RegisterRequest) error {

	// 判断 邮箱是否已经注册
	i, errua := models.NewUserAuth().UsernameIsMl("2", c.Email)
	if errua != nil {
		return errua
	}
	if len(i) != 0 {
		return errors.New("邮箱已注册")
	}
	// 取redis 验证码 对比传入验证码
	val, rdberr := initalize.Rdb.Get(context.Background(), c.Email).Result()
	if rdberr != nil {
		return rdberr
	}
	if val != c.Code {
		return errors.New("邮箱验证码错误")
	}

	// 组装数据
	uid, _ := gonanoid.New()
	u := models.NewUser()
	u.NickName = c.NickName
	u.RegisterSource = 2
	u.CreatedAt = time.Now()
	err := u.Register(uid, c.Email)
	if err != nil {
		return err
	}
	return nil
}

// Login 登录
func (t *Auth) Login(c apiForm.LoginRequest) (*apiForm.LoginResponse, error) {

	// 判断 邮箱是否存在
	i, errua := models.NewUserAuth().UsernameIsMl("2", c.Email)
	if errua != nil {
		return nil, errua
	}
	if len(i) == 0 {
		return nil, errors.New("邮箱未注册")
	}
	// 取redis 验证码 对比传入验证码
	val, rdberr := initalize.Rdb.Get(context.Background(), c.Email).Result()
	if rdberr != nil {
		return nil, rdberr
	}
	if val != c.Code {
		return nil, errors.New("邮箱验证码错误")
	}
	// 获取 用户资料信息
	u := models.NewUser()
	info, err := u.GetUserInfo(i[0].UserUid)
	if err != nil {
		return nil, err
	}
	// 生成token
	token, err := utils.CreateToken(info.UserUid, viper.GetString("Jwt.Secret"))
	if err != nil {
		return nil, err
	}
	// 组装返回数据
	return &apiForm.LoginResponse{
		Token:     token,
		UserUid:   info.UserUid,
		NickName:  info.NickName,
		Signature: info.Signature,
		Sex:       info.Sex,
		Age:       info.Age,
		Birthday:  info.Birthday,
		Face:      info.Face,
	}, nil
}
