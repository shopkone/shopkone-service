package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/seo/mSeo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/handle"
)

type LabelImage struct {
	Value   string `json:"value"`
	Label   string `json:"label"`
	ImageId uint   `json:"image_id"`
}

type BaseProduct struct {
	Id                uint                     `json:"id" dc:"商品id"`
	Title             string                   `json:"title" v:"required" dc:"标题"`
	Description       string                   `json:"description" dc:"描述"`
	Status            mProduct.VariantStatus   `json:"status" dc:"商品状态"`
	Spu               string                   `json:"spu" dc:"商品编码"`
	Vendor            string                   `json:"vendor" dc:"供应商"`
	Tags              []string                 `json:"tags" dc:"标签"`
	Seo               mSeo.Seo                 `json:"seo" v:"required" dc:"seo信息"`
	VariantType       mProduct.VariantType     `json:"variant_type" dc:"商品类型"`
	InventoryPolicy   mProduct.InventoryPolicy `json:"inventory_policy" dc:"库存策略"`
	FileIds           []uint                   `json:"file_ids" dc:"文件id"`
	Variants          []BaseVariant            `json:"variants" v:"required" dc:"商品变体"`
	InventoryTracking bool                     `json:"inventory_tracking"` // 库存跟踪
	ScheduledAt       int64                    `json:"scheduled_at"`
	LabelImages       []LabelImage             `json:"label_images"`
	Category          uint                     `json:"category" dc:"商品分类"`
	Collections       []uint                   `json:"collections"`
}

type ProductCreateReq struct {
	g.Meta             `path:"/product/create" method:"post" tags:"Product" summary:"创建商品"`
	EnabledLocationIds []uint `json:"enabled_location_ids" dc:"启用库存的仓库id"`
	BaseProduct
}
type ProductCreateRes struct {
	Id uint `json:"id" dc:"商品id"`
}

type ProductInfoReq struct {
	g.Meta `path:"/product/info" method:"post" tags:"Product" summary:"获取商品详情"`
	Id     uint `json:"id" v:"required" dc:"商品id"`
}
type ProductInfoRes struct {
	ProductCreateReq
	Id uint `json:"id" v:"required" dc:"商品id"`
}

type ProductListReq struct {
	g.Meta         `path:"/product/list" method:"post" tags:"Product" summary:"获取商品列表"`
	TrackInventory int8    `json:"track_inventory"` // 跟踪库存 1. 跟踪 2. 不跟踪
	ExcludeIds     *[]uint `json:"exclude_ids" dc:"排除商品id列表"`
	IncludeIds     *[]uint `json:"include_ids" dc:"包含商品id列表"`
	handle.PageReq
}
type ProductListRes struct {
	Id                uint                   `json:"id" dc:"商品id"`
	Spu               string                 `json:"spu" dc:"商品编码"`
	Vendor            string                 `json:"vendor" dc:"供应商"`
	CreatedAt         int64                  `json:"created_at"`
	Status            mProduct.VariantStatus `json:"status" dc:"商品状态"`
	Title             string                 `json:"title"`
	Variants          []VariantList          `json:"variants"`
	Image             string                 `json:"image"`
	InventoryTracking bool                   `json:"inventory_tracking"`
}

type ProductUpdateReq struct {
	g.Meta `path:"/product/update" method:"post" tags:"Product" summary:"更新商品"`
	ProductInfoRes
	Variants           []BaseVariantWithId `json:"variants"`
	EnabledLocationIds []uint              `json:"enabled_location_ids" dc:"启用库存的仓库id"`
}
type ProductUpdateRes struct {
}

type ListByIdsReq struct {
	g.Meta `path:"/product/list-by-ids" method:"post" tags:"Product" summary:"根据商品id列表获取商品列表"`
	Ids    []uint `json:"ids" v:"required" dc:"商品id列表"`
}
type ListByIdsRes struct {
	Id       uint                   `json:"id" dc:"商品id"`
	MinPrice float32                `json:"min_price"`
	MaxPrice float32                `json:"max_price"`
	Status   mProduct.VariantStatus `json:"status"`
	Image    string                 `json:"image"`
	Title    string                 `json:"title"`
}

type CreateSupplierReq struct {
	g.Meta  `path:"/product/supplier/create" method:"post" tags:"Product" summary:"创建供应商"`
	Address mAddress.Address `json:"address" v:"required"`
}
type CreateSupplierRes struct {
	Id uint `json:"id"`
}

type SupplierListReq struct {
	g.Meta `path:"/product/supplier/list" method:"post" tags:"Product" summary:"获取供应商列表"`
}
type SupplierListRes struct {
	Id      uint             `json:"id"`
	Address mAddress.Address `json:"address"`
}

type SupplierUpdateReq struct {
	g.Meta  `path:"/product/supplier/update" method:"post" tags:"Product" summary:"更新供应商"`
	Address mAddress.Address `json:"address" v:"required"`
	Id      uint             `json:"id" v:"required"`
}
type SupplierUpdateRes struct {
}
