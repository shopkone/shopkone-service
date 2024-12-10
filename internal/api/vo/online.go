package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/online/mav/mNav"
)

type OnlineNavListReq struct {
	g.Meta `path:"online/nav/list" method:"post" summary:"获取导航列表" tags:"Nav"`
}
type OnlineNavListRes struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Handle     string   `json:"handle"`
	FirstNames []string `json:"first_names"`
}

type OnlineNavInfoReq struct {
	g.Meta `path:"online/nav/info" method:"post" summary:"获取导航详情" tags:"Nav"`
	ID     uint `json:"id" v:"required" dc:"导航ID"`
}

type OnlineNavInfoRes struct {
	ID     uint           `json:"id"`
	Title  string         `json:"title"`
	Handle string         `json:"handle"`
	Links  []mNav.NavItem `json:"links"`
}

type OnlineNavUpdateReq struct {
	g.Meta `path:"/online/nav/update" method:"post" summary:"更新导航" tags:"Nav"`
	ID     uint           `json:"id" v:"required" dc:"导航ID"`
	Title  string         `json:"title" v:"required" dc:"导航标题"`
	Handle string         `json:"handle" v:"required" dc:"导航链接"`
	Links  []mNav.NavItem `json:"links" v:"required" dc:"导航链接"`
}
type OnlineNavUpdateRes struct {
}

type OnlineNavCreateReq struct {
	g.Meta `path:"/online/nav/create" method:"post" summary:"创建导航" tags:"Nav"`
	Title  string         `json:"title" v:"required" dc:"导航标题"`
	Handle string         `json:"handle" v:"required" dc:"导航链接"`
	Links  []mNav.NavItem `json:"links" v:"required" dc:"导航链接"`
}
type OnlineNavCreateRes struct {
	ID uint `json:"id"`
}
