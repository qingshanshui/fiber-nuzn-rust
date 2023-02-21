package apiForm

// CategoryRequest 详情
type LoginRequest struct {
	Username string `form:"username" json:"username"  validate:"required"` //  账号，验证规则：必填
	Password string `form:"password" json:"password"  validate:"required"` //  密码，验证规则：必填
}

// CategoryRequest 详情
type RegisterRequest struct {
	Username       string `form:"username" json:"username"  validate:"required"`             //  账号，验证规则：必填
	Password       string `form:"password" json:"password"  validate:"required"`             //  密码，验证规则：必填
	VerifyPassword string `form:"verifyPassword" json:"verifyPassword"  validate:"required"` //  校验密码，验证规则：必填
}
