package api

import (
	"fiber-nuzn-rust/controllers"
	"fiber-nuzn-rust/initalize"
	apiService "fiber-nuzn-rust/service/api"
	"fiber-nuzn-rust/validator"
	apiForm "fiber-nuzn-rust/validator/form/api"

	"github.com/gofiber/fiber/v2"
)

type CommentController struct {
	controllers.Base
}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// 留言列表
func (t *CommentController) List(c *fiber.Ctx) error {
	// 初始化参数结构体
	CommentListRequestForm := apiForm.CommentListRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &CommentListRequestForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := apiService.NewCommentService().List(CommentListRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
