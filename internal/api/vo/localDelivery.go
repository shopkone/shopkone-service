package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
)

type BaseLocalDeliveryFee struct {
	Condition float64 `json:"condition" v:"required" dc:"条件"`
	Fee       float64 `json:"fee" v:"required" dc:"费用"`
	Id        uint    `json:"id"`
}

type BaseLocalDeliverArea struct {
	Id         uint                   `json:"id"`
	Name       string                 `json:"name" v:"required" dc:"名称"`
	PostalCode string                 `json:"postal_code" v:"required" dc:"邮编"`
	Note       string                 `json:"note" dc:"备注"`
	Fees       []BaseLocalDeliveryFee `json:"fees" v:"required" dc:"费用"`
}

type UpdateLocalDeliveryReq struct {
	g.Meta `path:"/delivery/local-delivery/update" method:"post" tags:"Delivery" summary:"更新本地配送"`
	Id     uint                               `json:"id" v:"required" dc:"ID"`
	Status mLocalDelivery.LocalDeliveryStatus `json:"status" v:"required" dc:"状态"`
	Areas  []BaseLocalDeliverArea             `json:"areas" dc:"地区"`
}
type UpdateLocalDeliveryRes struct {
}

type LocalDeliveryListReq struct {
	g.Meta `path:"/delivery/local-delivery/list" method:"post" tags:"Delivery" summary:"获取本地配送列表"`
}
type LocalDeliveryListRes struct {
	Id         uint                               `json:"id"`
	Status     mLocalDelivery.LocalDeliveryStatus `json:"status"`
	LocationId uint                               `json:"location_id"`
}

type LocalDeliveryInfoReq struct {
	g.Meta `path:"/delivery/local-delivery/info" method:"post" tags:"Delivery" summary:"获取本地配送详情"`
	Id     uint `json:"id" v:"required" dc:"ID"`
}
type LocalDeliveryInfoRes struct {
	UpdateLocalDeliveryReq
	LocationId uint `json:"location_id"`
}
