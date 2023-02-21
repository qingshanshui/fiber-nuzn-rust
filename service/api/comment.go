package apiService

import (
	"fiber-nuzn-rust/models"
	apiForm "fiber-nuzn-rust/validator/form/api"
)

type Comment struct {
}

func NewCommentService() *Comment {
	return &Comment{}
}

func (t *Comment) List(c apiForm.CommentListRequest) (*models.Course, error) {

	list, err := models.NewCourse().Category(c.Username)
	if err != nil {
		return nil, err
	}
	return list, nil
}
