package models

import "gorm.io/gorm"

// UserAuth 用户授权表
type UserAuth struct {
	gorm.Model
	UserUid      string `gorm:"not null;comment:用户uid" json:"userUid"`                                 // '用户uid'
	IdentityType int    `gorm:"not null;comment:1:账号 2:邮箱 3:手机号 4:gitee 5:githup" json:"identityType"` // '1:账号 2:邮箱 3:手机号 4:gitee 5:githup'
	Identifier   string `gorm:"not null;comment:手机号 邮箱 用户名或第三方应用的唯一标识,类似账号" json:"identifier"`         // '手机号 邮箱 用户名或第三方应用的唯一标识,类似账号'
	Certificate  string `gorm:"comment:密码凭证(站内的保存密码,站外的不保存或保存token),类似密码" json:"certificate"`          // '密码凭证(站内的保存密码,站外的不保存或保存token),类似密码'
}
