package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/resource/iResource"
	"shopkone-service/internal/module/base/resource/mResource"
)

type AddressConfig struct {
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	Company    string `json:"company"`
	Country    string `json:"country"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postal_code"`
	Zone       string `json:"zone"`
}

type AddressFormatting struct {
	Country string `json:"country"`
	Format  string `json:"format"`
}

type CountriesReq struct {
	g.Meta `path:"/base/countries" method:"post" summary:"获取国家列表" tags:"Base"`
}
type CountriesRes struct {
	Code             string                     `json:"code"`       // 国家code
	Name             string                     `json:"name"`       // 国家名称
	Continent        string                     `json:"continent"`  // 大洲
	Flag             mResource.CountryFlag      `json:"flag"`       // 国家旗帜
	Zones            []iResource.ZoneListOut    `json:"zones"`      // 区号列表
	Config           AddressConfig              `json:"config"`     // 地址配置
	Formatting       string                     `json:"formatting"` // 地址格式
	PostalCodeConfig mResource.PostalCodeConfig `json:"postal_code_config"`
}

// PhonePrefixReq 获取手机区号列表
type PhonePrefixReq struct {
	g.Meta `path:"/base/phone-prefix" method:"post" summary:"获取手机区号列表" tags:"Base"`
}
type PhonePrefixRes struct {
	Code   string `json:"code"`   // 国家code
	Prefix int    `json:"prefix"` // 电话前缀
}

// TimezoneListReq 获取时区列表
type TimezoneListReq struct {
	g.Meta `path:"/base/timezone-list" method:"post" summary:"获取时区列表" tags:"Base"`
}
type TimezoneListRes struct {
	OlsonName   string `json:"olson_name"`  // 时区名称
	Description string `json:"description"` // 时区描述
}

// CurrencyListReq 获取货币列表
type CurrencyListReq struct {
	g.Meta `path:"/base/currency-list" method:"post" summary:"获取货币列表" tags:"Base"`
}
type CurrencyListRes struct {
	Code   string `json:"code"`   // 货币代码
	Title  string `json:"title"`  // 货币名称
	Symbol string `json:"symbol"` // 货币符号
}

// CategoryListReq 获取商品分类列表
type CategoryListReq struct {
	g.Meta `path:"/base/category-list" method:"post" summary:"获取商品分类列表" tags:"Base"`
}
type CategoryListRes struct {
	Label string `json:"label"`
	Value uint   `json:"value"`
	Deep  uint   `json:"deep"`
	Pid   uint   `json:"pid"`
}

// CarrierListReq 获取物流商列表
type CarrierListReq struct {
	g.Meta `path:"/base/carrier-list" method:"post" summary:"获取物流商列表" tags:"Base"`
}
type CarrierListRes struct {
	Id                       uint   `json:"id"`
	Name                     string `json:"name"`
	DisplayName              string `json:"display_name"`
	SupportsShipmentTracking bool   `json:"supports_shipment_tracking"`
}

type LanguagesReq struct {
	g.Meta `path:"/base/language-list" method:"post" summary:"获取语言列表" tags:"Base"`
}
type LanguagesRes struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
