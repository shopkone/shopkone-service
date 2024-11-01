package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"
)

type InStorePickUpListReq struct {
	g.Meta `path:"/delivery/in-store-pick-up/list" method:"post" tags:"Delivery" summary:"获取本地配送列表"`
}
type InStorePickUpListRes struct {
	Id         uint                               `json:"id"`
	Status     mInStorePickup.InStorePickupStatus `json:"status"`
	LocationID uint                               `json:"location_id"`
}

type BaseInStorePickUp struct {
	Id            uint                                 `json:"id" v:"required"`
	Status        mInStorePickup.InStorePickupStatus   `json:"status"`                // 是否开启
	IsUnified     bool                                 `json:"is_unified"`            // 是否统一时间
	Start         uint                                 `json:"start"`                 // 开始时间
	Timezone      string                               `json:"timezone" v:"required"` // 时区
	End           uint                                 `json:"end"`                   // 结束时间
	HasPickupETA  bool                                 `json:"has_pickup_eta"`        // 是否启用预计可取时间
	PickupETA     *uint                                `json:"pickup_eta"`            // 预计可取时间
	PickupETAUnit mInStorePickup.InStorePickupTimeUnit `json:"pickup_eta_unit"`       // 时间单位
	Weeks         []BaseInStorePickUpBusinessHours     `json:"weeks" v:"required"`    // 工作时间
	LocationID    uint                                 `json:"location_id"`
}

type BaseInStorePickUpBusinessHours struct {
	Id     uint  `json:"id" v:"required"`    // ID
	Week   uint8 `json:"week" v:"required"`  // 周几
	Start  uint  `json:"start" v:"required"` // 开始时间
	End    uint  `json:"end" v:"required"`   // 结束时间
	IsOpen bool  `json:"is_open"`            // 是否营业
}

type InStorePickUpUpdateReq struct {
	g.Meta `path:"/delivery/in-store-pick-up/update" method:"post" tags:"Delivery" summary:"更新本地配送"`
	BaseInStorePickUp
}
type InStorePickUpUpdateRes struct {
}

type InStorePickUpInfoReq struct {
	Id     uint `json:"id" v:"required" dc:"ID"`
	g.Meta `path:"/delivery/in-store-pick-up/info" method:"post" tags:"Delivery" summary:"获取本地配送详情"`
}
type InStorePickUpInfoRes struct {
	BaseInStorePickUp
}
