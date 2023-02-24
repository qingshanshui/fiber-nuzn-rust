package models

import (
	"fiber-nuzn-rust/initalize"
	"gorm.io/gorm"
)

// Article 文章表
type Article struct {
	gorm.Model
	ArticleUid string `gorm:"not null;comment:文章uid,文章唯一uid" json:"articleUid"`      // '文章uid,文章唯一uid'
	Title      string `gorm:"not null;comment:文章标题" json:"title"`                    // '文章标题'
	Content    string `gorm:"type:longtext;not null;comment:html内容" json:"content"`  // 'html内容'
	UserId     string `gorm:"not null;comment:用户id" json:"user_id"`                  // '用户id'
	Hot        int    `gorm:"default:1;comment:热度" json:"hot"`                       // 热度
	Top        int    `gorm:"comment:热度" json:"top"`                                 // 置顶
	Show       int    `gorm:"default:1;not null;comment:文章是否显示,1显示 2隐藏" json:"show"` // 文章是否显示,1显示 2隐藏
}

func NewArticle() *Article {
	return &Article{}
}

// GetList 获取文章列表
func (t *Article) GetList(page, size int) (*[]Article, error) {
	var r []Article
	if err := initalize.DB.Raw("SELECT * FROM articles LIMIT ?,?", (page-1)*size, size).Find(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

// GetTotal 获取文章列表总条数
func (t *Article) GetTotal() (*int, error) {
	var count int
	if err := initalize.DB.Raw("SELECT COUNT(*) FROM articles").Find(&count).Error; err != nil {
		return nil, err
	}
	return &count, nil
}

// GetDetails 获取文章详情
func (t *Article) GetDetails(id string) ([]Article, error) {
	var result []Article
	if err := initalize.DB.Raw("SELECT * FROM articles WHERE articles.article_uid = ? LIMIT 1", id).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
