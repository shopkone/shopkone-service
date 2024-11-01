package sShippingZoneFee

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

func (s *sShippingZoneFee) ConditionUpdateBatch(conditions []mShipping.ShippingZonFeeCondition) (err error) {
	if len(conditions) == 0 {
		return err
	}
	conditions = slice.Map(conditions, func(index int, item mShipping.ShippingZonFeeCondition) mShipping.ShippingZonFeeCondition {
		item.CanCreateId = true
		return item
	})
	in := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query: []string{
			"fixed",
			"first",
			"first_fee",
			"next",
			"next_fee",
			"max",
			"min",
		},
	}
	return handle.BatchUpdateById(in, &conditions)
}
