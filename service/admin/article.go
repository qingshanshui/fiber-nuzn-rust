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

// 用户文章列表
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

// 用户添加文章
func (t *Article) Add(c adminForm.ArticleAddRequest, uid string) (*[]models.Article, error) {

	list, err := models.NewArticle().GetUserAdd(&models.Article{}, uid)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 用户删除文章
func (t *Article) Del(c adminForm.ArticleDelRequest, uid string) error {
	err := models.NewArticle().GetUserDel(c.Id, uid)
	if err != nil {
		return err
	}
	return nil
}

// 用户编辑文章
func (t *Article) Edit(c adminForm.ArticleEditRequest, uid string) (*[]models.Article, error) {

	list, err := models.NewArticle().GetUserEdit(&models.Article{}, uid)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 用户文章详情
func (t *Article) Details(id, uid string) (*models.Article, error) {

	list, err := models.NewArticle().GetUserDetails(id, uid)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}
