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
	DomainType   mMarket.DomainType `json:"domain_type"`
	ID           uint               `json:"id"`
	IsMain       bool               `json:"is_main"`
	Name         string             `json:"name"`
	CountryCodes []string           `json:"country_codes"`
}

// 获取市场详情
type MarketInfoReq struct {
	g.Meta `path:"/market/info" method:"post" summary:"Market Info" tags:"Market"`
	ID     uint `json:"id" v:"required"`
}

type MarketInfoRes struct {
	ID                uint               `json:"id"`
	IsMain            bool               `json:"is_main"`
	Name              string             `json:"name"`
	CountryCodes      []string           `json:"country_codes"`
	DomainType        mMarket.DomainType `json:"domain_type"`
	DomainSuffix      string             `json:"domain_suffix"`
	SubDomainID       uint               `json:"sub_domain_id"`
	DefaultLanguageId uint               `json:"default_language_id"`
	LanguageIds       []uint             `json:"language_ids"`
	CurrencyCode      string             `json:"currency_code"`
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
	Label             string             `json:"label"`
	Value             uint               `json:"value"`
	IsMain            bool               `json:"is_main"`
	LanguageIds       []uint             `json:"language_ids"`
	DefaultLanguageId uint               `json:"default_language_id"`
	DomainType        mMarket.DomainType `json:"domain_type"`
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

// 更新市场语言
type MarketUpdateLangReq struct {
	g.Meta            `path:"/market/update-lang" method:"post" summary:"Market Update Lang" tags:"Market"`
	ID                uint   `json:"id" v:"required"`
	LanguageIds       []uint `json:"language_ids" v:"required"`
	DefaultLanguageID uint   `json:"default_language_id" v:"required"`
}
type MarketUpdateLangRes struct {
}

// 根据语言id更新市场语言
type MarketUpdateLangByLangIDReq struct {
	g.Meta    `path:"/market/update-lang-by-lang-id" method:"post" summary:"Market Update Lang By Lang ID" tags:"Market"`
	LangId    uint   `json:"lang_id" v:"required"`
	MarketIds []uint `json:"market_ids"`
}
type MarketUpdateLangByLangIDRes struct {
}

type MarketUpdateProductItem struct {
	ID        uint     `json:"id" v:"required"`
	ProductID uint     `json:"product_id" v:"required"`
	Fixed     *float64 `json:"fixed"`
	Exclude   bool     `json:"exclude"`
}

// 更新商品价格调整
type MarketUpdateProductReq struct {
	g.Meta         `path:"/market/update-product" method:"post" summary:"Market Update Product" tags:"Market"`
	MarketID       uint                        `json:"market_id" v:"required"`
	AdjustPercent  float64                     `json:"adjust_percent"`
	AdjustType     mMarket.PriceAdjustmentType `json:"adjust_type"`
	AdjustProducts []MarketUpdateProductItem   `json:"adjust_products"`
	CurrencyCode   string                      `json:"currency_code"`
}
type MarketUpdateProductRes struct {
}

// 或许商品价格调整
type MarketGetProductReq struct {
	g.Meta   `path:"/market/get-product" method:"post" summary:"Market Get Product" tags:"Market"`
	MarketID uint `json:"market_id" v:"required"`
}
type MarketGetProductRes struct {
	AdjustPercent         float64                     `json:"adjust_percent"`
	AdjustType            mMarket.PriceAdjustmentType `json:"adjust_type"`
	AdjustProducts        []MarketUpdateProductItem   `json:"adjust_products"`
	CurrencyCode          string                      `json:"currency_code"`
	ExchangeRate          float64                     `json:"exchange_rate"`
	ExchangeRateTimeStamp int64                       `json:"exchange_rate_time_stamp"`
}

type MarketSimpleReq struct {
	g.Meta `path:"/market/simple" method:"post" summary:"Market Simple" tags:"Market"`
	ID     uint `json:"id"`
}
type MarketSimpleRes struct {
	Name   string `json:"name"`
	IsMain bool   `json:"is_main"`
}
