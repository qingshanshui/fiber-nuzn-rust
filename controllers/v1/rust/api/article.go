package api

import (
	"fiber-nuzn-rust/controllers"
	"fiber-nuzn-rust/initalize"
	apiService "fiber-nuzn-rust/service/api"
	"fiber-nuzn-rust/validator"
	apiForm "fiber-nuzn-rust/validator/form/api"
	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	controllers.Base
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// List 文章列表
func (t *ArticleController) List(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleListRequestForm := &apiForm.ArticleListRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, ArticleListRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := apiService.NewArticleService().List(ArticleListRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}

// Details 文章详情
func (t *ArticleController) Details(c *fiber.Ctx) error {
	// 实际业务调用
	api, err := apiService.NewArticleService().Details(c.Params("id"))
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
