package admin

import (
	"fiber-nuzn-rust/controllers"
	"fiber-nuzn-rust/initalize"
	apiService "fiber-nuzn-rust/service/api"
	"fiber-nuzn-rust/validator"
	apiForm "fiber-nuzn-rust/validator/form/api"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	controllers.Base
}

func NewUserController() *UserController {
	return &UserController{}
}

// Userinfo 用户详情
func (t *UserController) Userinfo(c *fiber.Ctx) error {
	// 初始化参数结构体
	UserInfoRequestForm := apiForm.UserInfoRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &UserInfoRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := apiService.NewUserService().List(UserInfoRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
