package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/consts"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

type BaseZoneCode struct {
	Id          uint     `json:"id"`
	CountryCode string   `json:"country_code" v:"required"` // 国家代码
	ZoneCodes   []string `json:"zone_codes"`                // 省份代码
}

// BaseShippingZoneFeeCondition 表示运费条件
type BaseShippingZoneFeeCondition struct {
	ID       uint    `json:"id"`        // 条件ID
	Fixed    float32 `json:"fixed"`     // 固定费用
	First    float32 `json:"first"`     // 首重/首件
	FirstFee float32 `json:"first_fee"` // 首重/首件费用
	Next     float32 `json:"next"`      // 续重/续件
	NextFee  float32 `json:"next_fee"`  // 续重/续件费用
	Max      float32 `json:"max"`       // 最大值限制
	Min      float32 `json:"min"`       // 最小值限制
}

// BaseShippingZoneFee 表示物流区域方案运费
type BaseShippingZoneFee struct {
	ID           uint                           `json:"id"`                         // 运费ID
	Name         string                         `json:"name" v:"required"`          // 运费名称
	WeightUnit   consts.WeightUnit              `json:"weight_unit" v:"required"`   // 重量单位
	Type         mShipping.ShippingZoneFeeType  `json:"type" v:"required"`          // 运费类型
	CurrencyCode string                         `json:"currency_code" v:"required"` // 货币代码
	Conditions   []BaseShippingZoneFeeCondition `json:"conditions" v:"required"`    // 运费条件
	Rule         mShipping.ShippingZoneFeeRule  `json:"rule"`                       // 匹配规则
	Remark       string                         `json:"remark"`                     // 备注
	Cod          bool                           `json:"cod"`                        // 是否货到付款
}

// BaseShippingZone 表示物流区域方案
type BaseShippingZone struct {
	ShippingID uint                  `json:"-"`                  // 所属方案
	ID         uint                  `json:"id"`                 // 区域ID
	Name       string                `json:"name" v:"required"`  // 区域名称
	Codes      []BaseZoneCode        `json:"codes" v:"required"` // 国家/地区代码
	Fees       []BaseShippingZoneFee `json:"fees" v:"required"`  // 费用
}

// BaseShipping 表示物流方案
type BaseShipping struct {
	ID          uint                   `json:"id"`                 // 物流方案ID
	Name        string                 `json:"name"`               // 物流方案名称
	Type        mShipping.ShippingType `json:"type"`               // 物流方案类型
	ProductIDs  []uint                 `json:"product_ids"`        // 绑定的商品ID
	LocationIDs []uint                 `json:"location_ids"`       // 绑定的地点ID
	Zones       []BaseShippingZone     `json:"zones" v:"required"` // 物流区域方案
}

type ShippingCreateReq struct {
	g.Meta `path:"/shipping/create" method:"post" tags:"物流" summary:"创建物流"`
	BaseShipping
}
type ShippingCreateRes struct {
	g.Meta `mime:"application/json"`
	ID     uint `json:"id"`
}

type ShippingInfoReq struct {
	g.Meta `path:"/shipping/info" method:"post" tags:"物流" summary:"物流详情"`
	ID     uint `json:"id" v:"required"`
}
type ShippingInfoRes struct {
	g.Meta `mime:"application/json"`
	BaseShipping
}

type ShippingListReq struct {
	g.Meta `path:"/shipping/list" method:"post" tags:"物流" summary:"物流列表"`
}
type ShippingListRes struct {
	Id            uint                   `json:"id"`
	Name          string                 `json:"name"`
	Type          mShipping.ShippingType `json:"type"`
	ProductCount  uint                   `json:"product_count"`
	LocationCount uint                   `json:"location_count"`
	ZoneCount     uint                   `json:"zone_count"`
}

type ShippingUpdateReq struct {
	g.Meta `path:"/shipping/update" method:"post" tags:"物流" summary:"更新物流"`
	BaseShipping
	Id uint `json:"id" v:"required"`
}
type ShippingUpdateRes struct {
}

type ShippingZoneListByCountriesReq struct {
	g.Meta       `path:"/shipping/zone/list-by-countries" method:"post" tags:"物流" summary:"物流区域列表"`
	CountryCodes []string `json:"country_codes" v:"required"`
}
type ShippingZoneListByCountriesRes struct {
	BaseShippingZone
	ShippingID   uint   `json:"shipping_id"`
	ShippingName string `json:"shipping_name"`
}
