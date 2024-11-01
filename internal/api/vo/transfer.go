package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/utility/handle"
)

type BaseTransferItem struct {
	Id        uint `json:"id"`
	VariantId uint `json:"variant_id" v:"required"`
	Quantity  uint `json:"quantity" v:"required"`
	Rejected  uint `json:"rejected"`
	Received  uint `json:"received"`
}

type TransferCreateReq struct {
	g.Meta           `path:"/transfer/create" method:"post" tags:"Transfer" summary:"创建库存转移"`
	OriginId         uint               `json:"origin_id" v:"required"`
	DestinationId    uint               `json:"destination_id" v:"required"`
	Items            []BaseTransferItem `json:"items" v:"required"`
	CarrierId        *uint              `json:"carrier_id"`
	DeliveryNumber   string             `json:"delivery_number"`
	EstimatedArrival int64              `json:"estimated_arrival"`
}

type TransferCreateRes struct {
	Id uint `json:"id"`
}

type TransferListReq struct {
	g.Meta `path:"/transfer/list" method:"post" tags:"Transfer" summary:"获取库存转移列表"`
	handle.PageReq
}
type TransferListRes struct {
	Id               uint                     `json:"id"`
	OriginId         uint                     `json:"origin_id"`
	DestinationId    uint                     `json:"destination_id"`
	EstimatedArrival int64                    `json:"estimated_arrival"`
	Status           mTransfer.TransferStatus `json:"status"`
	Received         uint                     `json:"received"`
	Rejected         uint                     `json:"rejected"`
	Quantity         uint                     `json:"quantity"`
	TransferNumber   string                   `json:"transfer_number"`
}

type TransferInfoReq struct {
	g.Meta `path:"/transfer/info" method:"post" tags:"Transfer" summary:"库存转移详情"`
	Id     uint `json:"id" v:"required"`
}

type TransferInfoRes struct {
	TransferCreateReq
	Id             uint                     `json:"id"`
	Status         mTransfer.TransferStatus `json:"status"`
	TransferNumber string                   `json:"transfer_number"`
}

type TransferMarkReq struct {
	g.Meta `path:"/transfer/mark" method:"post" tags:"Transfer" summary:"标记库存转移状态"`
	Id     uint `json:"id" v:"required"`
}
type TransferMarkRes struct {
}

type TransferAdjustItem struct {
	VariantID uint `json:"variant_id"`
	Rejected  uint `json:"rejected"`
	Received  uint `json:"received"`
}
type TransferAdjustReq struct {
	g.Meta `path:"/transfer/adjust" method:"post" tags:"Transfer" summary:"库存转移调整"`
	Id     uint                 `json:"id" v:"required"`
	Items  []TransferAdjustItem `json:"items" v:"required"`
}
type TransferAdjustRes struct{}

type TransferRemoveReq struct {
	g.Meta `path:"/transfer/remove" method:"post" tags:"Transfer" summary:"��弃库存转移"`
	Id     uint `json:"id" v:"required"`
}
type TransferRemoveRes struct{}

type TransferUpdateReq struct {
	g.Meta `path:"/transfer/update" method:"post" tags:"Transfer" summary:"修改库存转移"`
	TransferCreateReq
	Id uint `json:"id" v:"required"`
}
type TransferUpdateRes struct{}
