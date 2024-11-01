package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/email/iEmail"
)

// Login 登录
type LoginReq struct {
	g.Meta   `path:"/account/login" method:"post" summary:"登录" tags:"Account"`
	Email    string `v:"required|email"`
	Password string `v:"required|length:8,30"`
}
type LoginRes struct {
	Token string `json:"token"`
}

// 退出登录
type LogoutReq struct {
	g.Meta `path:"/account/logout" method:"post" summary:"退出登录" tags:"Account"`
}
type LogoutRes struct {
}

// 注册
type RegisterReq struct {
	g.Meta   `path:"/account/register" method:"post" summary:"注册" tags:"Account"`
	Email    string `json:"email" v:"required|email" dc:"邮箱"`
	Password string `json:"password" v:"required" dc:"密码"`
	Code     string `json:"code" v:"required" dc:"验证码"`
}
type RegisterRes struct {
}

// 重置密码
type ResetPwdReq struct {
	g.Meta   `path:"/account/reset-pwd" method:"post" summary:"重置密码" tags:"Account"`
	Email    string `json:"email" v:"required|email" dc:"邮箱"`
	Code     string `json:"code" v:"required" dc:"验证码"`
	Password string `json:"password" v:"required" dc:"密码"`
}
type ResetPwdRes struct {
}

// 接受邀请
type AcceptInviteReq struct {
	g.Meta   `path:"/account/accept-invite" method:"post" summary:"接受邀请" tags:"Account"`
	Token    string `json:"token" v:"required" dc:"邀请令牌"`
	Password string `json:"password"  dc:"密码"`
}
type AcceptInviteRes struct {
	ShopUUId string `json:"shop_uuid"`
	Token    string `json:"token"`
}

// 校验邀请
type CheckInviteReq struct {
	g.Meta `path:"/account/check-invite" method:"post" summary:"校验邀请" tags:"Account"`
	Token  string `json:"token" v:"required" dc:"邀请令牌"`
}
type CheckInviteRes struct {
	ShopName string `json:"shop_name"`
	ShopUUID string `json:"shop_uuid"`
	ShopID   uint   `json:"shop_id"`
	UserID   uint   `json:"user_id"`
	StaffID  uint   `json:"staff_id"`
	Email    string `json:"email"`
}

// 发送验证码
type SendCodeReq struct {
	g.Meta `path:"/account/send-code" method:"post" summary:"发送验证码" tags:"Account"`
	Email  string       `v:"required|email"`
	Type   iEmail.Sense `v:"required|in:register,reset-pwd"`
}
type SendCodeRes struct {
}
