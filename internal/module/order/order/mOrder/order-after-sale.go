package mOrder

// OrderAfterSale 售后
type OrderAfterSale struct {
	OrderVariantId uint `gorm:"index;not null"`
}
