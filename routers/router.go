package routers

import (
	"fiber-nuzn-rust/controllers/v1/rust/admin"
	"fiber-nuzn-rust/controllers/v1/rust/api"
	"fiber-nuzn-rust/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {

	group := app.Group("/v1/rust/api")

	// 登录
	authGroup := group.Group("/auth")
	auth := api.NewAuthController()
	authGroup.Post("/code", auth.Code)         // 发送验证码      ***
	authGroup.Post("/register", auth.Register) // 注册      ***
	authGroup.Post("/login", auth.Login)       // 登录      ***

	// 用户
	userGroup := group.Group("/user")
	user := admin.NewUserController()
	userGroup.Post("/userinfo", user.Userinfo) // 获取用户详情

	// authGroup.Post("/logout")   // 退出登录
	// authGroup.Post("/code")         // 登录邮箱验证码
	// authGroup.Post("/oauth/email")  // 邮箱 登录/注册
	// authGroup.Post("/oauth/gitee")  // gitee 登录/注册
	// authGroup.Post("/oauth/githup") // githup回调接口

	// 文章
	articleGroup := group.Group("/article")
	article := api.NewArticleController()
	articleGroup.Post("/list", article.List)           // 获取所以文章列表 带分页      ***
	articleGroup.Post("/details/:id", article.Details) // 通过id获取文章      ***

	// 留言
	commentGroup := group.Group("/comment")
	comment := api.NewCommentController()
	commentGroup.Post("/list/:id", comment.List) // 查看文章全部留言

	//
	//
	//
	// 个人中心 (统一需要登录权限才能访问)
	//
	//
	//

	adminGroup := group.Group("/admin")
	adminGroup.Use(middleware.Auth)

	// 文章
	adminArticleGroup := adminGroup.Group("/article")
	adminArticle := admin.NewArticleController()
	adminArticleGroup.Post("/list", adminArticle.List)   // admin文章列表
	adminArticleGroup.Post("/add", adminArticle.Add)     // admin添加文章
	adminArticleGroup.Post("/edit", adminArticle.Edit)   // admin编辑文章
	adminArticleGroup.Post("/del", adminArticle.Del)     // admin删除文章
	adminArticleGroup.Post("/:id", adminArticle.Details) // admin文章详情

	// userGroup.Post("/userinfo/update/:id") // 更新用户详情

}
