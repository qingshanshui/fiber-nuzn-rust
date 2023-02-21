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

func (t *Article) List(c apiForm.ArticleListRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t *Article) Details(c apiForm.ArticleDetailsRequest) ([]models.Course, error) {
	list, err := models.NewCourse().List()
	if err != nil {
		return nil, err
	}
	return list, nil
}
