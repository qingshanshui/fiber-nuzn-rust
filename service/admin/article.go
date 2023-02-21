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

func (t *Article) List(c adminForm.ArticleListRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Add(c adminForm.ArticleAddRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Del(c adminForm.ArticleDelRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Edit(c adminForm.ArticleEditRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Details(c adminForm.ArticleDetailsRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}
