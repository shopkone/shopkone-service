package vo

import "github.com/gogf/gf/v2/frame/g"

type OnlineNavListReq struct {
	g.Meta `path:"online/nav/list" method:"post" summary:"获取导航列表" tags:"Nav"`
}
type OnlineNavListRes struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Handle     string   `json:"handle"`
	FirstNames []string `json:"first_names"`
}
