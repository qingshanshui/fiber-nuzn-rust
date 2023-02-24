package api

import (
	"fiber-nuzn-rust/controllers"
	"fiber-nuzn-rust/initalize"
	apiService "fiber-nuzn-rust/service/api"
	"fiber-nuzn-rust/validator"
	apiForm "fiber-nuzn-rust/validator/form/api"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	controllers.Base
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// Code 发送验证码
func (t *AuthController) Code(c *fiber.Ctx) error {

	// 初始化参数结构体
	CodeRequestForm := &apiForm.CodeRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, CodeRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}

	// 实际业务调用
	err := apiService.NewAuthService().Code(CodeRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.JSON(t.Fail(err, 500))
	}
	return c.JSON(t.Ok("发送成功")) // => ✋ Login
}

// Register 注册
func (t *AuthController) Register(c *fiber.Ctx) error {
	// 初始化参数结构体
	RegisterRequestForm := &apiForm.RegisterRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, RegisterRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	err := apiService.NewAuthService().Register(RegisterRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.JSON(t.Fail(err, 500))
	}
	return c.JSON(t.Ok("api")) // => ✋ Login
}

// Login 登录
func (t *AuthController) Login(c *fiber.Ctx) error {

	// 初始化参数结构体
	LoginRequestForm := &apiForm.LoginRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, LoginRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	rep, err := apiService.NewAuthService().Login(LoginRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.JSON(t.Fail(err, 500))
	}
	return c.JSON(t.Ok(rep)) // => ✋ Login
}
