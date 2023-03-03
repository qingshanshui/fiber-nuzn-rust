package adminService

import (
	"fiber-nuzn-rust/models"
	adminForm "fiber-nuzn-rust/validator/form/admin"
)

type Article struct {
}

func NewArticleService() *Article {
	return &Article{}
}

func (t *Article) List(c adminForm.ArticleListRequest, uid string) (*adminForm.ArticleListResponse, error) {

	if c.Size == 0 {
		c.Size = 16
	}
	if c.Page == 0 {
		c.Page = 1
	}
	// 文章列表
	ar := models.NewArticle()
	list, err := ar.GetUserList(uid, c.Page, c.Size)
	if err != nil {
		return nil, err
	}
	// 文章总数
	total, err := ar.GetUserTotal(uid)
	if err != nil {
		return nil, err
	}
	return &adminForm.ArticleListResponse{
		Data:  list,
		Total: total,
	}, nil
}

func (t *Article) Add(c adminForm.ArticleAddRequest, uid string) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Del(c adminForm.ArticleDelRequest, uid string) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Edit(c adminForm.ArticleEditRequest, uid string) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Details(c adminForm.ArticleDetailsRequest, uid string) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}
