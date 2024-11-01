package sShippingZoneFee

import (
	"github.com/duke-git/lancet/v2/convertor"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

type ConditionCreateItem struct {
	Fixed             float64 `json:"fixed"`
	First             float64 `json:"first"`
	FirstFee          float64 `json:"first_fee"`
	Next              float64 `json:"next"`
	NextFee           float64 `json:"next_fee"`
	Max               float64 `json:"max"`
	Min               float64 `json:"min"`
	ShippingZoneFeeId uint    `json:"shipping_zone_fee_id"`
}

func (s *sShippingZoneFee) ConditionCreate(in []ConditionCreateItem) (err error) {
	var list []mShipping.ShippingZonFeeCondition
	for _, condition := range in {
		temp := mShipping.ShippingZonFeeCondition{}
		if err = convertor.CopyProperties(&temp, condition); err != nil {
			return
		}
		temp.ShopId = s.shopId
		list = append(list, temp)
	}
	return s.orm.Create(&list).Error
}
