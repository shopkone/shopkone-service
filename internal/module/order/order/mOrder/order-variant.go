package mOrder

import "shopkone-service/internal/module/product/product/mProduct"

type OrderVariantStatus uint8

const (
	OrderVariantStatusPending   OrderVariantStatus = 0 // 待处理
	OrderVariantStatusCompleted OrderVariantStatus = 1 // 已完成
	OrderVariantStatusCanceled  OrderVariantStatus = 2 // 已取消
)

type OrderVariant struct {
	Product     mProduct.Product   `gorm:"serializer:json"` // 变体
	Variant     mProduct.Variant   `gorm:"serializer:json"` // 变体
	Status      OrderVariantStatus `gorm:"not null"`        // 商品状态
	SaleTaxRate float32            `gorm:"default:0"`       // 售税
	SaleTaxName string             `gorm:"size:200"`        // 售税名称
}
