package sDeliveryArea

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/utility/handle"
)

func (s *sDeliveryArea) DeliveryFeeUpdateBatch(fees []mLocalDelivery.LocalDeliveryFee) error {
	if len(fees) == 0 {
		return nil
	}
	fees = slice.Map(fees, func(index int, item mLocalDelivery.LocalDeliveryFee) mLocalDelivery.LocalDeliveryFee {
		item.CanCreateId = true
		return item
	})
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"fee", "condition"},
	}
	return handle.BatchUpdateById(batchIn, &fees)
}
