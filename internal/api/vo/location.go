package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
)

type LocationListReq struct {
	g.Meta `path:"/location/list" method:"post" tags:"Location" summary:"获取位置列表"`
	Active *bool `json:"active"`
}
type LocationListRes struct {
	Id      uint             `json:"id"`
	Name    string           `json:"name"`
	Address mAddress.Address `json:"address"`
	Active  bool             `json:"active"`
	Default bool             `json:"default"`
	Order   uint             `json:"order"`
}

type LocationAddReq struct {
	g.Meta  `path:"/location/add" method:"post" tags:"Location" summary:"添加位置"`
	Name    string           `v:"required" json:"name"`
	Address mAddress.Address `json:"address" v:"required"`
}
type LocationAddRes struct {
	Id uint `json:"id"`
}

type LocationInfoReq struct {
	g.Meta `path:"/location/info" method:"post" tags:"Location" summary:"获取位置信息"`
	Id     uint `json:"id" v:"required" dc:"位置ID"`
}
type LocationInfoRes struct {
	Id                 uint             `json:"id"`
	Name               string           `json:"name"`
	Address            mAddress.Address `json:"address"`
	Active             bool             `json:"active"`
	Default            bool             `json:"default"`
	FulfillmentDetails bool             `json:"fulfillment_details"`
}

type LocationUpdateReq struct {
	g.Meta             `path:"/location/update" method:"post" tags:"Location" summary:"更新位置信息"`
	Id                 uint             `json:"id" v:"required" dc:"位置ID"`
	Name               string           `json:"name" v:"required" dc:"位置名称"`
	Address            mAddress.Address `json:"address" v:"required" dc:"位置地址"`
	Active             bool             `json:"active"`
	FulfillmentDetails bool             `json:"fulfillment_details"`
}
type LocationUpdateRes struct {
}

type DeleteLocationReq struct {
	g.Meta `path:"/location/delete" method:"post" tags:"Location" summary:"删除位置"`
	Id     uint `json:"id" v:"required" dc:"位置ID"`
}
type DeleteLocationRes struct {
}

type LocationExistInventoryReq struct {
	g.Meta `path:"/location/exist-inventory" method:"post" tags:"Location" summary:"位置是否有库存"`
	Id     uint `json:"id" v:"required" dc:"位置ID"`
}
type LocationExistInventoryRes struct {
	Exist bool `json:"exist"`
}

type SetDefaultLocationReq struct {
	g.Meta `path:"/location/set-default" method:"post" tags:"Location" summary:"设置默认位置"`
	Id     uint `json:"id" v:"required" dc:"位置ID"`
}
type SetDefaultLocationRes struct {
}

type SetLocationOrderItem struct {
	LocationId uint `json:"location_id" v:"required" dc:"位置ID"`
	Order      uint `json:"order" v:"required" dc:"排序"`
}
type SetLocationOrderReq struct {
	g.Meta `path:"/location/set-order" method:"post" tags:"Location" summary:"设置位置排序"`
	Items  []SetLocationOrderItem `json:"items" v:"required" dc:"排序项"`
}
type SetLocationOrderRes struct {
}
