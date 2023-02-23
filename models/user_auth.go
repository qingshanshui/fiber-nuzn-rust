package models

import (
	"fiber-nuzn-rust/initalize"

	"gorm.io/gorm"
)

// UserAuth 用户授权表
type UserAuth struct {
	gorm.Model
	UserUid      string `gorm:"not null;comment:用户uid" json:"userUid"`                                 // '用户uid'
	IdentityType int    `gorm:"not null;comment:1:账号 2:邮箱 3:手机号 4:gitee 5:githup" json:"identityType"` // '1:账号 2:邮箱 3:手机号 4:gitee 5:githup'
	Identifier   string `gorm:"not null;comment:手机号 邮箱 用户名或第三方应用的唯一标识,类似账号" json:"identifier"`         // '手机号 邮箱 用户名或第三方应用的唯一标识,类似账号'
	Certificate  string `gorm:"comment:密码凭证(站内的保存密码,站外的不保存或保存token),类似密码" json:"certificate"`          // '密码凭证(站内的保存密码,站外的不保存或保存token),类似密码'
}

func NewUserAuth() *UserAuth {
	return &UserAuth{}
}

// 查询当前 注册账号是否 存在数据库
func (t *UserAuth) UsernameIsMl(types, username string) ([]UserAuth, error) {
	results := []UserAuth{}
	result := initalize.DB.Raw("SELECT user_auths.id FROM user_auths WHERE user_auths.identity_type = ? AND user_auths.identifier = ?  LIMIT 1", types, username).Find(&results)
	if err := result.Error; err != nil {
		return nil, err
	}
	return results, nil
}
