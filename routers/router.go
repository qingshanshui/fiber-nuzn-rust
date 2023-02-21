package routers

import (
	v1 "fiber-layout/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	main := v1.NewDefaultController()
	group := app.Group("/v1")

	// 登录
	loginGroup := group.Group("/login")
	loginGroup.Get("/code")         // 登录邮箱验证码
	loginGroup.Get("/oauth/email")  // 邮箱 登录/注册
	loginGroup.Get("/oauth/gitee")  // gitee 登录/注册
	loginGroup.Get("/oauth/githup") // githup回调接口
	loginGroup.Get("/logout")       // 退出登录

	// 文章
	articleGroup := group.Group("/article")
	articleGroup.Get("/")    // 获取所以文章 带分页
	articleGroup.Get("/:id") // 通过id获取文章

	// 留言/回复
	commentGroup := group.Group("/comment")
	commentGroup.Get("/") // 查看全部回复

	// 个人中心 (统一需要登录权限才能访问)
	userGroup := group.Group("/user")
	userGroup.Get("/info/:id") // 获取用户详情
	userGroup.Put("/info/:id") // 更新用户详情

	group.Get("/list", main.List)          // 列表
	group.Post("/category", main.Category) // 详情
}
