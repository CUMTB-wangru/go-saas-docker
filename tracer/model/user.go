package model

import (
	"gorm.io/gorm"
)

// UserInfo 用户表

type UserInfo struct {
	gorm.Model
	UserName string `gorm:"type:varchar(32);unique" json:"user_name" label:"用户名"`
	Password string `gorm:"size:60" json:"password" label:"密码"`
	Phone    string `gorm:"size:11;unique" json:"phone" label:"手机号"`
	Email    string `gorm:"size:32;unique" json:"email" label:"邮箱"`
}
