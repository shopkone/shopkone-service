package mUser

import "gorm.io/gorm"

type UserLoginRecord struct {
	gorm.Model
	Ip          string `gorm:"size:50;not null"`   // ip地址
	Ua          string `gorm:"size:500;not null"`  // 浏览器ua
	Token       string `gorm:"size:64;not null"`   // 登录token
	IsActive    bool   `gorm:"default:1;not null"` // 是否激活
	Fingerprint string `gorm:"size:500;not null"`  // 浏览器指纹
	UserId      uint   `gorm:"index;not null"`     // 用户ID
}
