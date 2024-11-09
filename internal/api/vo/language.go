package vo

import "github.com/gogf/gf/v2/frame/g"

type LanguageListReq struct {
	g.Meta `path:"/setting/language/list" method:"post" tags:"Setting" summary:"获取语言列表"`
}

type LanguageListRes struct {
	ID          uint     `json:"id"`
	Language    string   `json:"language"`
	IsDefault   bool     `json:"is_default"`
	MarketNames []string `json:"market_names"`
}

type LanguageCreateReq struct {
	g.Meta `path:"/setting/language/create" method:"post" tags:"Setting" summary:"创建语言"`
	Codes  []string `json:"codes" v:"required"`
}
type LanguageCreateRes struct {
}
