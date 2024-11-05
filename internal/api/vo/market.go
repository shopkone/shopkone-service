package vo

import "github.com/gogf/gf/v2/frame/g"

type MarketCreateReq struct {
	g.Meta       `path:"/market/create" method:"post" summary:"Create Market" tags:"Market"`
	Name         string   `json:"name" v:"required"`
	CountryCodes []string `json:"countryCodes" v:"required"`
	IsMain       bool     `json:"-"`
}
type MarketCreateRes struct {
	ID uint `json:"id"`
}
