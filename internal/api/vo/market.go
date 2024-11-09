package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/setting/market/mMarket"
)

// 创建市场
type MarketCreateReq struct {
	g.Meta       `path:"/market/create" method:"post" summary:"Create Market" tags:"Market"`
	Name         string   `json:"name" v:"required"`
	CountryCodes []string `json:"countryCodes" v:"required"`
	IsMain       bool     `json:"-"`
	Force        bool     `json:"force"`
}
type MarketCreateRes struct {
	ID          uint     `json:"id"`
	RemoveNames []string `json:"remove_names"`
}

// 获取市场列表
type MarketListReq struct {
	g.Meta `path:"/market/list" method:"post" summary:"Market List" tags:"Market"`
}
type MarketListRes struct {
	ID           uint     `json:"id"`
	IsMain       bool     `json:"is_main"`
	Name         string   `json:"name"`
	CountryCodes []string `json:"country_codes"`
}

// 获取市场详情
type MarketInfoReq struct {
	g.Meta `path:"/market/info" method:"post" summary:"Market Info" tags:"Market"`
	ID     uint `json:"id" v:"required"`
}
type MarketInfoRes struct {
	ID           uint               `json:"id"`
	IsMain       bool               `json:"is_main"`
	Name         string             `json:"name"`
	CountryCodes []string           `json:"country_codes"`
	DomainType   mMarket.DomainType `json:"domain_type"`
	DomainSuffix string             `json:"domain_suffix"`
	SubDomainID  uint               `json:"sub_domain_id"`
}

// 更新市场
type MarketUpdateReq struct {
	g.Meta `path:"/market/update" method:"post" summary:"Market Update" tags:"Market"`
	MarketCreateReq
	ID uint `json:"id" v:"required"`
}
type MarketUpdateRes struct {
	RemoveNames []string `json:"remove_names"`
}

// 获取市场简易列表
type MarketOptionsReq struct {
	g.Meta `path:"/market/options" method:"post" summary:"Market Options" tags:"Market"`
}
type MarketOptionsRes struct {
	Label  string `json:"label"`
	Value  uint   `json:"value"`
	IsMain bool   `json:"is_main"`
}

// 更新市场域名
type MarketUpDomainReq struct {
	g.Meta       `path:"/market/up-domain" method:"post" summary:"Market Up Domain" tags:"Market"`
	ID           uint               `json:"id" v:"required"`
	DomainType   mMarket.DomainType `json:"domain_type" v:"required"`
	DomainSuffix string             `json:"domain_suffix"`
	SubDomainID  uint               `json:"sub_domain_id"`
}
type MarketUpDomainRes struct {
}

// 根据语言id更新语言市场绑定
type BindLangByLangIdReq struct {
	g.Meta     `path:"/market/bind-lang-by-lang-id" method:"post" summary:"Market Bind Lang By Lang Id" tags:"Market"`
	MarketIDs  []uint `json:"market-ids" v:"required"`
	LanguageID uint   `json:"language_id" v:"required"`
}
type BindLangByLangIdRes struct {
}

// 根据市场id更新语言市场绑定
type BindLangByMarketIdReq struct {
	g.Meta            `path:"/market/bind-lang-by-market-id" method:"post" summary:"Market Bind Lang By Market Id" tags:"Market"`
	MarketID          uint   `json:"market_id" v:"required"`
	LanguageIDs       []uint `json:"language_ids" v:"required"`
	DefaultLanguageID uint   `json:"default_language_id"`
}
type BindLangByMarketIdRes struct {
}
