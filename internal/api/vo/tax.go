package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/setting/tax/mTax"
)

type TaxListReq struct {
	g.Meta `path:"/tax/list" method:"post" tags:"Tax" summary:"获取税率列表"`
}
type TaxListRes struct {
	Id          uint           `json:"id"`
	TaxRate     float64        `json:"tax_rate"`
	CountryCode string         `json:"country_code"`
	Status      mTax.TaxStatus `json:"status"`
}

type BaseTax struct {
	ID          uint              `json:"id"`
	TaxRate     float64           `json:"tax_rate" v:"required" dc:"税率"`
	HasNote     bool              `json:"has_note"`
	Note        string            `json:"note"`
	Status      mTax.TaxStatus    `json:"status"`
	CountryCode string            `json:"country_code"`
	Zones       []BaseTaxZone     `json:"zones"`
	Customers   []BaseCustomerTax `json:"customers"`
}
type BaseTaxZone struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name" v:"required:true"`
	ZoneCode string  `json:"zone_code" v:"required:true"`
	TaxRate  float64 `json:"tax_rate"`
}
type BaseCustomerTax struct {
	ID           uint                  `json:"id"`
	CollectionID uint                  `json:"collection_id"`
	Zones        []BaseCustomerTaxZone `json:"zones"`
	Type         mTax.CustomerTaxType  `json:"type"`
}

type BaseCustomerTaxZone struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" v:"required:true"`
	CountryCode string  `json:"country_code"`
	ZoneCode    string  `json:"zone_code"`
	TaxRate     float64 `json:"tax_rate"`
}

type TaxInfoReq struct {
	g.Meta `path:"/tax/info" method:"post" tags:"Tax" summary:"获取税率详情"`
	Id     uint `json:"id" v:"required" dc:"ID"`
}
type TaxInfoRes struct {
	BaseTax
}

type TaxUpdateReq struct {
	g.Meta `path:"/tax/update" method:"post" tags:"Tax" summary:"更新税率"`
	BaseTax
	ID uint `json:"id" v:"required" dc:"ID"`
}
type TaxUpdateRes struct {
	TaxInfoRes
}

type TaxCreateReq struct {
	g.Meta       `path:"/tax/create" method:"post" tags:"Tax" summary:"创建税率"`
	CountryCodes []string `json:"country_codes" v:"required" dc:"国家代码"`
}
type TaxCreateRes struct {
}

type TaxRemoveReq struct {
	g.Meta `path:"/tax/remove" method:"post" tags:"Tax" summary:"删除税率"`
	Ids    []uint `json:"ids" v:"required" dc:"ID"`
}
type TaxRemoveRes struct {
}

type TaxActiveReq struct {
	g.Meta `path:"/tax/active" method:"post" tags:"Tax" summary:"启用税率"`
	ID     uint `json:"id" v:"required" dc:"ID"`
	Active bool `json:"active"`
}
type TaxActiveRes struct {
	List []TaxListRes `json:"list"`
}
