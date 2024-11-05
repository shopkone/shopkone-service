package vo

import "github.com/gogf/gf/v2/frame/g"

type LanguageListReq struct {
	g.Meta `path:"/setting/language/list" method:"post" tags:"Setting" summary:"获取语言列表"`
}
type LanguageListRes struct {
	ID        uint   `json:"id"`
	Language  string `json:"language"`
	IsDefault bool   `json:"is_default"`
	MarketIds []uint `json:"market_ids"`
	IsActive  bool   `json:"is_active"`
}
