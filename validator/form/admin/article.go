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

// ArticleAddRequest 添加文章
type ArticleAddRequest struct {
	*models.Article
}

// ArticleDelRequest 删除文章
type ArticleDelRequest struct {
	Id string `form:"id" json:"id"  validate:"required"` //  账号，验证规则：必填
}

// ArticleEditRequest 编辑
type ArticleEditRequest struct {
	*models.Article
}
