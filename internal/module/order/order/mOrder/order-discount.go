package mOrder

import (
	"shopkone-service/internal/module/base/orm/mOrm"
	"time"
)

type OrderDiscountType uint8

const (
	OrderDiscountTypePercentage OrderDiscountType = 1 // 折扣类型：百分比
	OrderDiscountTypeFixed      OrderDiscountType = 2
)

type OrderDiscountStatus uint8

const (
	OrderDiscountStatusPending   OrderDiscountStatus = 1 // 折扣状态：待使用
	OrderDiscountStatusCompleted OrderDiscountStatus = 2 // 折扣状态：已使用
)

type OrderDiscountFromSource uint8

const (
	OrderDiscountFromSourceOrder OrderDiscountFromSource = 1 // 折扣来源：创建订单时自键
)

type OrderDiscount struct {
	mOrm.Model
	Type         OrderDiscountType       `gorm:"not null"`  // 折扣类型
	Value        float32                 `gorm:"default:0"` // 折扣金额
	CreatorEmail string                  `gorm:"size:255"`  // 创建人邮箱
	FromSource   OrderDiscountFromSource `json:"index"`     // 折扣来源
	FromID       uint                    `gorm:"index"`     // 折扣来源ID
	ToOrderID    uint                    `gorm:"index"`     // 消费订单ID
	CostAt       *time.Time              // 消费时间
	CustomerID   uint                    `gorm:"index; not null"` // 所属客户ID
}
