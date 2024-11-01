package mShipping

import (
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/product/inventory/mInventory"
)

type ShippingNoteStatus uint8

const (
	ShippingNoteStatusPending   ShippingNoteStatus = 1 // 待发货
	ShippingNoteStatusShipping  ShippingNoteStatus = 2 // 运输中
	ShippingNoteStatusCompleted ShippingNoteStatus = 3 // 已送达
	ShippingNoteStatusSignedFor ShippingNoteStatus = 4 // 已签收
	ShippingNoteStatusCancelled ShippingNoteStatus = 5 // 已取消
)

// 物流单
type ShippingNote struct {
	FreightPrice    float32            `gorm:"not null"`        // 运费
	FreightTaxRate  float32            `gorm:"not null"`        // 运费税率
	FreightPlanName string             `gorm:"size:200"`        // 运费方案
	Sender          mAddress.Address   `gorm:"serializer:json"` // 寄件人
	Receiver        mAddress.Address   `gorm:"serializer:json"` // 收件人
	HandleEmail     string             `gorm:"size:500"`        // 操作人
	Status          ShippingNoteStatus `gorm:"not null"`        // 状态
	Note            string             `gorm:"size:500"`        // 备注
}

// 物流单明细
type ShippingNoteItem struct {
	ShippingNoteId    uint                 `gorm:"index;not null"`  // 物流单id
	Quantity          uint                 `gorm:"not null"`        // 数量
	Note              string               `gorm:"size:500"`        // 备注
	Inventory         mInventory.Inventory `gorm:"serializer:json"` // 库存
	OrderVariantId    uint                 `gorm:"index;not null"`  // 商品id
	InventoryChangeId uint                 `gorm:"not null"`        // 库存变更 id
}
