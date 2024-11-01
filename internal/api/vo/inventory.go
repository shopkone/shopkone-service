package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/utility/handle"
)

type InventoryListReq struct {
	g.Meta     `path:"/inventory/list" method:"post" tags:"Inventory" summary:"获取库存列表"`
	LocationId uint `json:"location_id" v:"required" dc:"仓库id"`
	handle.PageReq
}

type InventoryListRes struct {
	Id          uint   `json:"id"`
	Quantity    uint   `json:"quantity"`
	ProductName string `json:"product_name"`
	VariantId   uint   `json:"variant_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Sku         string `json:"sku"`
}

type InventoryMoveReq struct {
	g.Meta `path:"/inventory/move" method:"post" tags:"Inventory" summary:"库存转移"`
	From   uint `json:"from" v:"required" dc:"来源仓库id"`
	To     uint `json:"to" v:"required" dc:"目标仓库id"`
}
type InventoryMoveRes struct {
}

type InventoryHistoryReq struct {
	g.Meta `path:"/inventory/history" method:"post" tags:"Inventory" summary:"库存历史记录"`
	Id     uint `json:"id" v:"required" dc:"库存id"`
}
type InventoryHistoryRes struct {
	Id           uint                     `json:"id"`
	DiffQuantity int                      `json:"diff_quantity"`
	Date         int64                    `json:"date"`
	Activity     mInventory.InventoryType `json:"activity"`
}

type InventoryListUnByVariantIdsReq struct {
	g.Meta `path:"/inventory/list-un-by-variant-ids" method:"post" tags:"Inventory" summary:"获取被删除的库存列表"`
	Ids    []uint `json:"ids" v:"required" dc:"变体id列表"`
}
type InventoryListUnByVariantIdsRes struct {
	LocationId uint `json:"location_id"`
	Quantity   uint `json:"quantity"`
	VariantId  uint `json:"variant_id"`
	Id         uint `json:"id"`
}
