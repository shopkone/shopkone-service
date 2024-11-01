package mInStorePickup

import (
	"shopkone-service/internal/module/base/orm/mOrm"
)

type InStorePickupStatus uint8

const (
	InStorePickupStatusOpen  InStorePickupStatus = 1
	InStorePickupStatusClose InStorePickupStatus = 2
)

type InStorePickupTimeUnit uint8

const (
	InStorePickupTimeUnitMinute InStorePickupTimeUnit = 1 // 分钟
	InStorePickupTimeUnitHour   InStorePickupTimeUnit = 2 // 小时
	InStorePickupTimeUnitDay    InStorePickupTimeUnit = 3 // 天
)

type InStorePickup struct {
	mOrm.Model
	Status        InStorePickupStatus   `gorm:"index"` // 是否开启
	IsUnified     bool                  // 是否统一时间
	Start         uint                  // 开始时间
	End           uint                  // 结束时间
	HasPickupETA  bool                  // 是否启用预计可取时间
	PickupETA     *uint                 // 预计可取时间
	PickupETAUnit InStorePickupTimeUnit `gorm:"default:1"`      // 时间单位
	LocationId    uint                  `gorm:"index;not null"` // 地点id
	Timezone      string                `gorm:"size50"`
}

type InStorePickupBusinessHours struct {
	mOrm.Model
	Week            uint8 // 周几
	Start           uint  // 开始时间
	End             uint  // 结束时间
	IsOpen          bool  `gorm:"index"`
	InStorePickupID uint  `gorm:"index;not null"`
}
