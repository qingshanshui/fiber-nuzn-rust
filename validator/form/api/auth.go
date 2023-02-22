package apiForm

// CodeRequest 发送验证码
type CodeRequest struct {
	Email string `form:"email" json:"email"  validate:"required,email"` //  邮箱，验证规则：必填，邮箱类型
}

// LoginRequest 详情
type LoginRequest struct {
	Username string `form:"username" json:"username"  validate:"required"` //  账号，验证规则：必填
	Password string `form:"password" json:"password"  validate:"required"` //  密码，验证规则：必填
}

// RegisterRequest 详情
type RegisterRequest struct {
	Username       string `form:"username" json:"username"  validate:"required"`             //  账号，验证规则：必填
	Password       string `form:"password" json:"password"  validate:"required"`             //  密码，验证规则：必填
	VerifyPassword string `form:"verifyPassword" json:"verifyPassword"  validate:"required"` //  校验密码，验证规则：必填
}
