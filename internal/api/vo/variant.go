package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/consts"
	"shopkone-service/internal/module/product/product/mProduct"
)

type VariantInventory struct {
	Id         uint `json:"id" v:"required"`
	Quantity   uint `json:"quantity"`
	LocationId uint `json:"location_id"`
}

type BaseVariant struct {
	Id               uint                   `json:"id" dc:"变体ID"`
	Price            float32                `json:"price" v:"required"` // 价格
	CostPerItem      *float32               `json:"cost_per_item"`      // 每件成本
	CompareAtPrice   *float32               `json:"compare_at_price"`   // 原价
	WeightUnit       consts.WeightUnit      `json:"weight_unit"`        // 重量单位
	Weight           *float32               `json:"weight"`             // 重量
	Sku              string                 `json:"sku"`                // SKU
	Barcode          string                 `json:"barcode"`            // 条形码
	ImageId          uint                   `json:"image_id"`           // 图片ID
	Name             []mProduct.VariantName `json:"name"`               // 名称
	Inventories      []VariantInventory     `json:"inventories"`        // 库存
	TaxRequired      bool                   `json:"tax_required"`       // 是否需要计税
	ShippingRequired bool                   `json:"shipping_required"`  // 是否需要配送
}

type BaseVariantWithId struct {
	BaseVariant
	Id uint `json:"id" v:"required" dc:"变体ID"`
}

type VariantList struct {
	Id       uint                   `json:"id"`
	Price    float32                `json:"price"`
	Quantity uint                   `json:"quantity"`
	Sku      string                 `json:"sku"`
	Name     []mProduct.VariantName `json:"name"`
}

type VariantListByIdsReq struct {
	g.Meta `path:"/product/variant/list-by-ids" method:"post" summary:"根据变体ID列表获取变体列表" tags:"Product"`
	Ids    []uint `json:"ids" v:"required" dc:"变体ID列表"`
}
type VariantListByIdsRes struct {
	Id           uint   `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	ProductTitle string `json:"product_title"`
	IsDeleted    bool   `json:"is_deleted"`
}
