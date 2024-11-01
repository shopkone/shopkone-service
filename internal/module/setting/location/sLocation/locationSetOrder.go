package sLocation

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/location/mLocation"
	"shopkone-service/utility/handle"
)

func (s *sLocation) LocationSetOrder(items []vo.SetLocationOrderItem) (err error) {
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"order_num"},
	}
	locations := slice.Map(items, func(index int, item vo.SetLocationOrderItem) mLocation.Location {
		i := mLocation.Location{}
		i.ID = item.LocationId
		i.OrderNum = item.Order
		i.ShopId = s.shopId
		i.CanCreateId = true
		return i
	})
	return handle.BatchUpdateById(batchIn, &locations)
}
