package sDeliveryArea

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/utility/handle"
)

func (s *sDeliveryArea) DeliveryAreaUpdateBatch(areas []mLocalDelivery.LocalDeliveryArea) (err error) {
	if len(areas) == 0 {
		return err
	}
	areas = slice.Map(areas, func(_ int, item mLocalDelivery.LocalDeliveryArea) mLocalDelivery.LocalDeliveryArea {
		item.CanCreateId = true
		return item
	})
	in := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"name", "postal_code", "note"},
	}
	return handle.BatchUpdateById(in, &areas)
}
