package models

import (
	"fiber-nuzn-rust/initalize"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	UserUid        string `gorm:"not null;comment:用户uid" json:"user_uid"`                                        // '用户ID,唯一uid'
	RegisterSource int8   `gorm:"not null;comment:注册来源:1:账号 2:邮箱 3:手机号 4:gitee 5:githup" json:"register_source"` // '注册来源：1:账号 2:邮箱 3:手机号 4:gitee 5:githup'
	NickName       string `gorm:"comment:用户昵称" json:"nick_name"`                                                 // '用户昵称'
	Signature      string `gorm:"comment:用户个人签名" json:"signature"`                                               // '用户个人签名'
	Sex            uint8  `gorm:"default:3;comment:性别:1:男 2:女 3:未知" json:"sex"`                                  // 性别，1:男 2:女 3:未知
	Age            string `gorm:"comment:年龄" json:"age"`                                                         // 年龄
	Birthday       string `gorm:"comment:生日" json:"birthday"`                                                    // 生日
	Status         int8   `gorm:"default:1;comment:用户状态:1正常用户 2禁言用户" json:"status"`                              // '用户状态；1正常用户 2禁言用户'
	Face           string `gorm:"comment:头像" json:"face"`                                                        // '头像'
}

func NewUser() *User {
	return &User{}
}

// Register 注册用户
func (t *User) Register(user_uid, email string) error {

	// 事务写入两张表
	if err := initalize.DB.Transaction(func(tx *gorm.DB) error {
		// 插入主用户表 （user表）
		if err := tx.Debug().Exec("INSERT INTO users (users.user_uid,users.register_source,users.nick_name,users.created_at) VALUES(?,?,?,?)", user_uid, t.RegisterSource, t.NickName, t.CreatedAt).Error; err != nil {
			return err
		}
		// 插入账号用户表（user_auth表）
		if err := tx.Debug().Exec("INSERT INTO user_auths (user_auths.user_uid,user_auths.identity_type,user_auths.identifier,user_auths.created_at)VALUES(?,?,?,?)", user_uid, t.RegisterSource, email, t.CreatedAt).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// GetUserInfo 获取用户信息
func (t *User) GetUserInfo(uid string) (*User, error) {
	if err := initalize.DB.Raw("SELECT * FROM users WHERE users.user_uid = ? LIMIT 1", uid).Find(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
