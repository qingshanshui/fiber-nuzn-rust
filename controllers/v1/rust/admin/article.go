package admin

import (
	"fiber-nuzn-rust/controllers"
	"fiber-nuzn-rust/initalize"
	adminService "fiber-nuzn-rust/service/admin"
	"fiber-nuzn-rust/validator"
	adminForm "fiber-nuzn-rust/validator/form/admin"

	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	controllers.Base
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// List 用户 文章列表
func (t *ArticleController) List(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleListRequestForm := adminForm.ArticleListRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ArticleListRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := adminService.NewArticleService().List(ArticleListRequestForm, c.Locals("uid").(string))
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}

// Add 用户 添加文章
func (t *ArticleController) Add(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleAddRequestForm := adminForm.ArticleAddRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ArticleAddRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := adminService.NewArticleService().Add(ArticleAddRequestForm, c.Locals("uid").(string))
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}

// Edit 用户 编辑文章
func (t *ArticleController) Edit(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleEditRequestForm := adminForm.ArticleEditRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ArticleEditRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := adminService.NewArticleService().Edit(ArticleEditRequestForm, c.Locals("uid").(string))
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}

// Del 用户 删除文章
func (t *ArticleController) Del(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleDelRequestForm := adminForm.ArticleDelRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ArticleDelRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := adminService.NewArticleService().Del(ArticleDelRequestForm, c.Locals("uid").(string))
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}

// Details 用户 文章详情
func (t *ArticleController) Details(c *fiber.Ctx) error {
	// 初始化参数结构体
	ArticleDetailsRequestForm := adminForm.ArticleDetailsRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ArticleDetailsRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := adminService.NewArticleService().Details(ArticleDetailsRequestForm, c.Locals("uid").(string))
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
