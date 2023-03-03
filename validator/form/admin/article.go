package adminForm

import "fiber-nuzn-rust/models"

// ArticleListRequest 文章列表--入参
type ArticleListRequest struct {
	Size int `form:"size" json:"size"` // 一页显示多少条
	Page int `form:"page" json:"page"` // 当前页
}

// ArticleListResponse 文章列表--出参
type ArticleListResponse struct {
	Data  *[]models.Article `form:"data" json:"data"`   // 文章数据
	Total *int              `form:"total" json:"total"` // 总条数
}

// ArticleAddRequest 详情
type ArticleAddRequest struct {
	Username string `form:"username" json:"username"  validate:"required"` //  账号，验证规则：必填
	Password string `form:"password" json:"password"  validate:"required"` //  密码，验证规则：必填
}

// ArticleDelRequest 详情
type ArticleDelRequest struct {
	Username string `form:"username" json:"username"  validate:"required"` //  账号，验证规则：必填
	Password string `form:"password" json:"password"  validate:"required"` //  密码，验证规则：必填
}

// ArticleEditRequest 详情
type ArticleEditRequest struct {
	Username string `form:"username" json:"username"  validate:"required"` //  账号，验证规则：必填
	Password string `form:"password" json:"password"  validate:"required"` //  密码，验证规则：必填
}

// ArticleDetailsRequest 详情
type ArticleDetailsRequest struct {
	Username string `form:"username" json:"username"  validate:"required"` //  账号，验证规则：必填
	Password string `form:"password" json:"password"  validate:"required"` //  密码，验证规则：必填
}
