package mTransfer

import (
	"shopkone-service/internal/module/base/orm/mOrm"
	"time"
)

type TransferStatus uint8

const (
	TransferStatusDraft           TransferStatus = 1 // 草稿
	TransferStatusOrdered         TransferStatus = 2 // 待收货
	TransferStatusPartialReceived TransferStatus = 3 // 部分收货
	TransferStatusReceived        TransferStatus = 4 // 已完成
)

type Transfer struct {
	mOrm.Model
	DestinationId    uint           `gorm:"index;not null"`
	OriginId         uint           `gorm:"index;not null"`
	TransferNumber   string         `gorm:"index;not null"`
	Status           TransferStatus `gorm:"default:0"`
	CarrierId        *uint          `gorm:"index"`    // 快递公司ID
	DeliveryNumber   string         `gorm:"size:200"` // 快递单号
	EstimatedArrival *time.Time     // 预计到货时间
}

type TransferItem struct {
	mOrm.Model
	Quantity   uint `gorm:"not null"`
	VariantId  uint `gorm:"index;not null"`
	TransferId uint `gorm:"index;not null"`
	Received   uint `gorm:"default:0"`
	Rejected   uint `gorm:"default:0"`
}
