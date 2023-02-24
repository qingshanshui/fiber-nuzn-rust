package apiService

import (
	"fiber-nuzn-rust/models"
	apiForm "fiber-nuzn-rust/validator/form/api"
)

type Article struct {
}

func NewArticleService() *Article {
	return &Article{}
}

// List 文章列表
func (t *Article) List(c *apiForm.ArticleListRequest) (*apiForm.ArticleListResponse, error) {
	if c.Size == 0 {
		c.Size = 16
	}
	if c.Page == 0 {
		c.Page = 1
	}
	// 文章列表
	ar := models.NewArticle()
	list, err := ar.GetList(c.Page, c.Size)
	if err != nil {
		return nil, err
	}
	// 文章总数
	total, err := ar.GetTotal()
	if err != nil {
		return nil, err
	}
	return &apiForm.ArticleListResponse{
		Data:  list,
		Total: total,
	}, nil
}

// Details 文章详情
func (t *Article) Details(id string) (*models.Article, error) {
	list, err := models.NewArticle().GetDetails(id)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}
