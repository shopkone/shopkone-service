package mPurchase

import (
	"shopkone-service/internal/module/base/orm/mOrm"
	"time"
)

type PaymentTermsType uint8

const (
	None             PaymentTermsType = iota // 无付款条款
	CashOnDelivery                           // 货到付款
	CashOnReceipt                            // 收款时付款
	PaymentOnReceipt                         // 收到后付款
	PaymentInAdvance                         // 预付款
	Net7                                     // 净付款7天
	Net15                                    // 净付款15天
	Net30                                    // 净付款30天
	Net45                                    // 净付款45天
	Net60                                    // 净付款60天
)

type AdjustType uint8

const (
	CustomsDuties AdjustType = 1 // 关税
	Discount      AdjustType = 2 // 折扣
	ForeignFee    AdjustType = 3 // 国外交易费
	FreightFee    AdjustType = 4 // 运费
	Insurance     AdjustType = 5 // 保险
	RushFee       AdjustType = 6 // 加急费
	Surcharge     AdjustType = 7 // 附加费
	Others        AdjustType = 8 // 其他
)

type PurchaseStatus uint8

const (
	PurchaseStatusDraft           PurchaseStatus = 1 // 草稿
	PurchaseStatusOrdered         PurchaseStatus = 2 // 已下单
	PurchaseStatusPartialReceived PurchaseStatus = 3 // 部分收货
	PurchaseStatusReceived        PurchaseStatus = 4 // 已收货
	PurchaseStatusClosed          PurchaseStatus = 5 // 关闭
)

// Purchase 采购记录
type Purchase struct {
	mOrm.Model
	PurchaseNumber   string               `gorm:"size:50;not null"` // 采购单号
	SupplierId       uint                 `gorm:"index;not null"`   // 供应商ID
	DestinationId    uint                 `gorm:"index;not null"`   // 仓库ID
	CarrierId        *uint                `gorm:"index"`            // 快递公司ID
	DeliveryNumber   string               `gorm:"size:200"`         // 快递单号
	CurrencyCode     string               `gorm:"index;not null"`   // 货币代码
	Remarks          string               `gorm:"size:500"`         // 备注
	Adjust           []PurchaseAdjustItem `gorm:"serializer:json"`  // 采购调整
	Status           PurchaseStatus       `gorm:"default:0"`        // 采购状态
	OldStatus        PurchaseStatus       `gorm:"default:0"`        // 旧状态
	EstimatedArrival *time.Time           // 预计到货时间
	PaymentTerms     PaymentTermsType     // 付款条款
}

// PurchaseItem 采购明细
type PurchaseItem struct {
	mOrm.Model
	VariantID  uint    `gorm:"index;not null"`
	Cost       float64 `gorm:"not null"`
	TaxRate    float64 `gorm:"not null"`
	SKU        string  `gorm:"size:255"`
	PurchaseId uint    `gorm:"index;not null"`
	Purchasing int     `gorm:"not null"`     // 需要采购的数量
	Rejected   int     `gorm:"default:0"`    // 拒绝数量
	Received   int     `gorm:"default:0"`    // 已收货数量
	Active     bool    `gorm:"default:true"` // 是否有效
}

// PurchaseAdjustItem 采购调整项
type PurchaseAdjustItem struct {
	Id       uint       `json:"id"`
	Type     AdjustType `json:"type"`
	Price    float64    `json:"price"`
	TypeText string     `json:"type_text"`
}
