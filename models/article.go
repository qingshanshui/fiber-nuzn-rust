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

// ----------admin----------

// GetUserList 获取用户文章列表
func (t *Article) GetUserList(uid string, page, size int) (*[]Article, error) {
	var r []Article
	if err := initalize.DB.Raw("SELECT * FROM articles WHERE articles.user_id = ?  LIMIT ?,?", uid, (page-1)*size, size).Find(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

// GetUserTotal 获取用户文章列表总条数
func (t *Article) GetUserTotal(uid string) (*int, error) {
	var count int
	if err := initalize.DB.Raw("SELECT COUNT(*) FROM articles WHERE articles.user_id = ? ", uid).Find(&count).Error; err != nil {
		return nil, err
	}
	return &count, nil
}

// GetUserDetails 获取用户文章详情
func (t *Article) GetUserDetails(id, uid string) ([]Article, error) {
	var result []Article
	if err := initalize.DB.Raw("SELECT * FROM articles WHERE articles.article_uid = ? and articles.user_id = ?  LIMIT 1", id, uid).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserDetails 删除用户文章
func (t *Article) GetUserDel(id, uid string) error {
	if err := initalize.DB.Exec("DELETE FROM articles WHERE articles.article_uid = ? AND articles.user_id = ?", id, uid).Error; err != nil {
		return err
	}
	return nil
}

// GetUserDetails 添加用户文章
func (t *Article) GetUserAdd(c *Article, uid string) (*[]Article, error) {
	var result []Article
	if err := initalize.DB.Raw("INSERT INTO articles (articles.article_uid,articles.content,articles.created_at,articles.hot,articles.`show`,articles.title,articles.top,articles.user_id) VALUES (?,?,?,?,?,?,?,?)",
		c.ArticleUid, c.Content, c.CreatedAt, c.Hot, c.Show, c.Title, c.Top, uid).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserDetails 编辑用户文章
func (t *Article) GetUserEdit(c *Article, uid string) (*[]Article, error) {
	var result []Article
	if err := initalize.DB.Raw("UPDATE articles SET articles.content = ? articles.hot = ? articles.`show` =? articles.title = ? articles.top =? articles.updated_at=? WHERE articles.article_uid = ? AND articles.user_id = ?",
		c.Content, c.Hot, c.Show, c.Title, c.Top, c.UpdatedAt, c.ArticleUid, uid).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
