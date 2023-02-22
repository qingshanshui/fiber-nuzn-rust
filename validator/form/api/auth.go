package apiForm

// CodeRequest 发送验证码--入参
type CodeRequest struct {
	Email string `form:"email" json:"email"  validate:"required,email,max=50,min=3"` //  邮箱，验证规则：必填，邮箱类型
}

// RegisterRequest 注册--入参
type RegisterRequest struct {
	NickName string `form:"nickName" json:"nickName"  validate:"required,max=50,min=2"` //  昵称，验证规则：必填
	Email    string `form:"email" json:"email"  validate:"required,email,max=50,min=2"` //  邮箱，验证规则：必填，邮箱类型
	Code     string `form:"code" json:"code"  validate:"required,max=10,min=6"`         //  验证码，验证规则：必填
}

// LoginRequest 登录--入参
type LoginRequest struct {
	Email string `form:"email" json:"email"  validate:"required,email,max=50,min=2"` //  邮箱，验证规则：必填，邮箱类型
	Code  string `form:"code" json:"code"  validate:"required,max=10,min=6"`         //  验证码，验证规则：必填
}
