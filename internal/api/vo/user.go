package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/shop/user/mUser"
)

// 获取登录信息
type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"post" summary:"获取当前登录信息" tags:"Auth"`
}
type UserInfoRes struct {
	Email    string `json:"email"`
	Language string `json:"language"`
	IsMaster bool   `json:"is_master"`
}

// SetColumns 设置列
type SetColumnsReq struct {
	g.Meta  `path:"/user/set-columns" method:"post" tags:"表格列" summary:"设置列"`
	Columns []mUser.UserColumnItem `json:"columns" v:"required"`
	Type    mUser.UserColumnType   `json:"type" v:"required"`
}
type SetColumnsRes struct {
	ID uint `json:"id"`
}

// GetColumns 获取列
type GetColumnsReq struct {
	g.Meta `path:"/user/get-columns" method:"post" tags:"表格列" summary:"获取列"`
	Type   mUser.UserColumnType `json:"type" v:"required"`
}
type GetColumnsRes struct {
	Columns []mUser.UserColumnItem `json:"columns"`
}
