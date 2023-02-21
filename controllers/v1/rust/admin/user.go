package admin

import (
	"fiber-nuzn-rust/controllers"
	"fiber-nuzn-rust/initalize"
	adminService "fiber-nuzn-rust/service/admin"
	"fiber-nuzn-rust/validator"
	adminForm "fiber-nuzn-rust/validator/form/admin"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	controllers.Base
}

func NewUserController() *UserController {
	return &UserController{}
}

// 用户详情
func (t *UserController) Userinfo(c *fiber.Ctx) error {
	// 初始化参数结构体
	UserInfoRequestForm := adminForm.UserInfoRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &UserInfoRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := adminService.NewUserService().List(UserInfoRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
