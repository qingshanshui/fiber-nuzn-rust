package middleware

import (
	"fiber-nuzn-rust/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Auth(ctx *fiber.Ctx) error {
	var header = ctx.GetReqHeaders()
	var token = header["Authorization"]
	if token == "" {
		return ctx.Status(401).JSON(fiber.Map{
			"code": 401,
			"data": "未传token",
			"msg":  "操作失败",
		})
	}
	MapClaims, err := utils.ParseToken(header["Authorization"], viper.GetString("Jwt.Secret"))
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"code": 401,
			"data": "token失效",
			"msg":  "操作失败",
		})
	}
	ctx.Locals("uid", MapClaims["uid"])
	return ctx.Next()
}
