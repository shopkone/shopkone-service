package iResource

import "shopkone-service/internal/module/base/resource/mResource"

// 区域列表
type ZoneListOut struct {
	Code string `json:"code"` // 区号
	Name string `json:"name"` // 区号名称
}

// 国家列表
type CountryListOut struct {
	Code             string                      `json:"code"`               // 国家code
	Name             string                      `json:"name"`               // 国家名称
	Continent        string                      `json:"continent"`          // 大洲
	Flag             mResource.CountryFlag       `json:"flag"`               // 国家旗帜
	Zones            []ZoneListOut               `json:"zones"`              // 区号列表
	Formatting       mResource.AddressFormatting `json:"formatting"`         // 地址格式
	Config           mResource.AddressLabel      `json:"config"`             // 地址标签
	PostalCodeConfig mResource.PostalCodeConfig  `json:"postal_code_config"` // 邮编配置
}

// 电话前缀列表
type PhonePrefixListOut struct {
	Code   string `json:"code"`   // 国家code
	Prefix int    `json:"prefix"` // 电话前缀
}
