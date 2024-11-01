package sShippingZoneFee

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

type FeeUpdateItem struct {
	FeeCreateItem
	Id uint `json:"id"`
}

func (s *sShippingZoneFee) FeeUpdateBatch(fees []mShipping.ShippingZoneFee) (err error) {
	if len(fees) == 0 {
		return nil
	}
	fees = slice.Map(fees, func(index int, item mShipping.ShippingZoneFee) mShipping.ShippingZoneFee {
		item.CanCreateId = true
		return item
	})
	in := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query: []string{
			"name",
			"weight_unit",
			"type",
			"currency_code",
			"rule",
			"remark",
			"cod",
		},
	}
	return handle.BatchUpdateById(in, &fees)
}
