package mShipping

import (
	"shopkone-service/internal/consts"
	"shopkone-service/internal/module/base/orm/mOrm"
)

// ShippingType 表示物流方案的类型
type ShippingType uint8

// 物流方案类型定义
var (
	GeneralExpressDelivery  ShippingType = 1 // 自定义方案
	CustomerExpressDelivery ShippingType = 2 // 通用方案
)

// ShippingZoneFeeRule 表示物流区域运费规则
type ShippingZoneFeeRule uint8

// 运费规则类型定义
var (
	ShippingZoneFeeRuleOrderPrice   ShippingZoneFeeRule = 1 // 按订单价格
	ShippingZoneFeeRuleProductPrice ShippingZoneFeeRule = 2 // 按商品价格
	ShippingZoneFeeRuleProductCount ShippingZoneFeeRule = 3 // 按商品数量
	ShippingZoneFeeRuleOrderWeight  ShippingZoneFeeRule = 4 // 按订单重量
)

// ShippingZoneFeeType 表示运费计算的类型
type ShippingZoneFeeType int8

// 运费计算类型定义
var (
	ShippingZoneFeeTypeFixed  ShippingZoneFeeType = 1 // 固定运费
	ShippingZoneFeeTypeWeight ShippingZoneFeeType = 2 // 按重量计费
	ShippingZoneFeeTypeCount  ShippingZoneFeeType = 3 // 按商品数量
)

// Shipping 表示物流方案
type Shipping struct {
	mOrm.Model
	Type ShippingType `gorm:"default:1"` // 物流方案类型
	Name string       `gorm:"size:200"`  // 物流方案名称
}

// ShippingProduct 表示物流方案与商品的绑定关系
type ShippingProduct struct {
	mOrm.Model
	ProductId  uint `gorm:"index"` // 商品ID
	ShippingId uint `gorm:"index"` // 物流方案ID
}

// ShippingLocation 表示物流方案与地点的绑定关系
type ShippingLocation struct {
	mOrm.Model
	LocationId uint `gorm:"index"` // 地点ID
	ShippingId uint `gorm:"index"` // 物流方案ID
}

// ShippingZone 表示物流区域方案
type ShippingZone struct {
	mOrm.Model
	Name       string `gorm:"size:200"` // 物流区域名称
	ShippingId uint   `gorm:"index"`    // 物流方案ID
}

// ShippingZoneCode 表示物流区域代码
type ShippingZoneCode struct {
	mOrm.Model
	CountryCode    string   `gorm:"size:50"`         // 国家或地区代码
	ZoneCodes      []string `gorm:"serializer:json"` // 区域代码
	ShippingZoneId uint     `gorm:"index"`           // 物流区域方案ID
}

// ShippingZoneFee 表示物流区域的运费规则
type ShippingZoneFee struct {
	mOrm.Model
	Name           string              `json:"name" gorm:"size:200"`          // 运费规则名称
	WeightUnit     consts.WeightUnit   `json:"weight_unit" gorm:"size:50"`    // 重量单位
	Type           ShippingZoneFeeType `json:"type" gorm:"default:1"`         // 运费类型
	CurrencyCode   string              `json:"currency_code" gorm:"size:10"`  // 货币代码
	Rule           ShippingZoneFeeRule `json:"rule" gorm:"default:0"`         // 匹配规则
	Remark         string              `json:"remark" gorm:"size:200"`        // 备注
	Cod            bool                `json:"cod" gorm:"default:false"`      // 是否支持货到付款
	ShippingZoneId uint                `json:"shipping_zone_id" gorm:"index"` // 物流区域方案ID
}

// ShippingZonFeeCondition 表示物流区域方案运费的条件
type ShippingZonFeeCondition struct {
	mOrm.Model
	Fixed             float64 `json:"fixed" gorm:"default:0"`            // 固定费用
	First             float64 `json:"first" gorm:"default:0"`            // 首重/首件
	FirstFee          float64 `json:"first_fee" gorm:"default:0"`        // 首重/首件的费用
	Next              float64 `json:"next" gorm:"default:0"`             // 续重/续件
	NextFee           float64 `json:"next_fee" gorm:"default:0"`         // 续重/续件的费用
	Max               float64 `json:"max" gorm:"default:0"`              // 最大值限制
	Min               float64 `json:"min" gorm:"default:0"`              // 最小值限制
	ShippingZoneFeeId uint    `json:"shipping_zone_fee_id" gorm:"index"` // 关联物流区域方案运费ID
}
