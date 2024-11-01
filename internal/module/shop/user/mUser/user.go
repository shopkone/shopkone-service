package mUser

import "gorm.io/gorm"

// 用户状态
type UserStatus int

const (
	UserStatusRegistered  UserStatus = iota + 1 // 已注册
	UserStatusPreRegister                       // 未注册
	UserStatusLocked                            // 冻结
)

/*User 用户*/
type User struct {
	gorm.Model
	Email       string     `gorm:"size:128;uniqueIndex"`         // 邮箱
	Password    string     `gorm:"size:300;index"`               // 密码
	LoginTimes  int        `gorm:"default:0"`                    // 登录次数
	LoginTokens []string   `gorm:"serializer:json"`              // 登录令牌
	Language    string     `gorm:"default:en;size:30"`           // 语言
	Status      UserStatus `gorm:"default:0;type:TINYINT;index"` // 状态
}
