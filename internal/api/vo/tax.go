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
	ID          uint           `json:"id"`
	TaxRate     float64        `json:"tax_rate" v:"required" dc:"税率"`
	HasNote     bool           `json:"has_note"`
	Note        string         `json:"note"`
	Status      mTax.TaxStatus `json:"status"`
	CountryCode string         `json:"country_code"`
}
type BaseTaxZone struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ZoneCode string `json:"zone_code"`
	TaxRate  float64
}

type TaxInfoReq struct {
	g.Meta `path:"/tax/info" method:"post" tags:"Tax" summary:"获取税率详情"`
	Id     uint `json:"id" v:"required" dc:"ID"`
}
type TaxInfoRes struct {
	BaseTax
	Zones []BaseTaxZone `json:"zones"`
}
