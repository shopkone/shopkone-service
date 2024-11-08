package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/setting/market/mMarket"
)

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

type MarketListReq struct {
	g.Meta `path:"/market/list" method:"post" summary:"Market List" tags:"Market"`
}
type MarketListRes struct {
	ID           uint     `json:"id"`
	IsMain       bool     `json:"is_main"`
	Name         string   `json:"name"`
	CountryCodes []string `json:"country_codes"`
}

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

type MarketUpdateReq struct {
	g.Meta `path:"/market/update" method:"post" summary:"Market Update" tags:"Market"`
	MarketCreateReq
	ID uint `json:"id" v:"required"`
}
type MarketUpdateRes struct {
	RemoveNames []string `json:"remove_names"`
}

type MarketOptionsReq struct {
	g.Meta `path:"/market/options" method:"post" summary:"Market Options" tags:"Market"`
}
type MarketOptionsRes struct {
	Label  string `json:"label"`
	Value  uint   `json:"value"`
	IsMain bool   `json:"is_main"`
}

type LanguageBindItem struct {
	LanguageId uint `json:"language_id"`
	MarketId   uint `json:"market_id"`
	IsDefault  bool `json:"is_default"`
}

type MarketBindLangReq struct {
	g.Meta `path:"/market/bind-lang" method:"post" summary:"Market Bind Lang" tags:"Market"`
	Bind   []LanguageBindItem `json:"bind"`
	UnBind []LanguageBindItem `json:"un_bind"`
}
type MarketBindLangRes struct {
}

type MarketUpDomainReq struct {
	g.Meta       `path:"/market/up-domain" method:"post" summary:"Market Up Domain" tags:"Market"`
	ID           uint               `json:"id" v:"required"`
	DomainType   mMarket.DomainType `json:"domain_type" v:"required"`
	DomainSuffix string             `json:"domain_suffix"`
	SubDomainID  uint               `json:"sub_domain_id"`
}
type MarketUpDomainRes struct {
}
