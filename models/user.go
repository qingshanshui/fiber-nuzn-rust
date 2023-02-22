package models

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	UserUid        string `gorm:"not null;comment:用户uid" json:"user_uid"`                          // '用户ID,唯一uid'
	RegisterSource int8   `gorm:"not null;comment:注册来源:1邮箱 2gitee 3githup" json:"register_source"` // '注册来源：1邮箱 2gitee 3githup'
	NickName       string `gorm:"comment:用户昵称" json:"nick_name"`                                   // '用户昵称'
	Signature      string `gorm:"comment:用户个人签名" json:"signature"`                                 // '用户个人签名'
	Sex            uint8  `gorm:"default:3;comment:性别:1:男 2:女 3:未知" json:"sex"`                    // 性别，1:男 2:女 3:未知
	Age            string `gorm:"comment:年龄" json:"age"`                                           // 年龄
	Birthday       string `gorm:"comment:生日" json:"birthday"`                                      // 生日
	Status         int8   `gorm:"default:1;comment:用户状态:1正常用户 2禁言用户" json:"status"`                // '用户状态；1正常用户 2禁言用户'
	Face           string `gorm:"comment:头像" json:"face"`                                          // '头像'
}
