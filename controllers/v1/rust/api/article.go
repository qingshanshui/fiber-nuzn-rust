package api

import (
	"fiber-nuzn-rust/controllers"
	"fiber-nuzn-rust/initalize"
	"fiber-nuzn-rust/service"
	"fiber-nuzn-rust/validator"
	"fiber-nuzn-rust/validator/form"

	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	controllers.Base
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// 文章列表
func (t *ArticleController) List(c *fiber.Ctx) error {
	// 初始化参数结构体
	categoryForm := form.CategoryRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &categoryForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().Category(categoryForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}

// 文章详情
func (t *ArticleController) Details(c *fiber.Ctx) error {
	// 初始化参数结构体
	categoryForm := form.CategoryRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &categoryForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().Category(categoryForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
