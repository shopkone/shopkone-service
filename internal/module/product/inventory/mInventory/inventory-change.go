package mInventory

import "shopkone-service/internal/module/base/orm/mOrm"

type InventoryType string

const (
	InventoryChangeProduct  InventoryType = "编辑商品"
	InventoryAddProduct                   = "添加商品"
	InventoryChangeAdjust                 = "手动调整库存"
	InventoryChangePurchase               = "采购单入库"
	InventoryChangeSale                   = "销售"
	InventoryChangeReturn                 = "退货"
	InventoryChangeTransfer               = "库存转移"
	InventoryChangeCreate                 = "添加地点"
)

type InventoryStyle uint8

const (
	InventoryStyleIn  InventoryStyle = iota + 1 // 入库
	InventoryStyleOut                           // 出库
)

type InventoryChange struct {
	mOrm.Model
	HandleEmail  string         `gorm:"size:500;not null"` // 经手人
	DiffQuantity int            `gorm:"not null"`          // 变动数量
	Note         string         `gorm:"size:500"`          // 备注
	Type         InventoryType  `gorm:"not null"`          // 变动类型
	Style        InventoryStyle `gorm:"not null"`          // 变动类型
	InventoryId  uint           `gorm:"index;not null"`    // 库存ID
}
