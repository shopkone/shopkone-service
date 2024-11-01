package sDeliveryArea

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
)

func (s *sDeliveryArea) DeliveryAreaList(localDeliveryId uint) (out []vo.BaseLocalDeliverArea, err error) {
	// 获取区域
	var data []mLocalDelivery.LocalDeliveryArea
	if err = s.orm.Model(&data).Where("local_delivery_id = ?", localDeliveryId).
		Where("shop_id = ?", s.shopId).
		Omit("created_at", "updated_at", "deleted_at").
		Where("shop_id = ?", s.shopId).Find(&data).Error; err != nil {
		return out, err
	}

	if len(data) == 0 {
		return out, nil
	}

	// 获取费用
	areaIds := slice.Map(data, func(index int, item mLocalDelivery.LocalDeliveryArea) uint {
		return item.ID
	})
	var fees []mLocalDelivery.LocalDeliveryFee
	if err = s.orm.Model(&fees).Where("local_delivery_area_id in ?", areaIds).
		Where("shop_id = ?", s.shopId).
		Omit("created_at", "updated_at", "deleted_at").
		Where("shop_id = ?", s.shopId).Find(&fees).Error; err != nil {
		return nil, err
	}

	// 组装数据
	for _, item := range data {
		temp := vo.BaseLocalDeliverArea{}
		if err = convertor.CopyProperties(&temp, item); err != nil {
			return nil, err
		}
		currentFees := slice.Filter(fees, func(index int, i mLocalDelivery.LocalDeliveryFee) bool {
			return i.LocalDeliveryAreaID == item.ID
		})
		temp.Fees = slice.Map(currentFees, func(index int, item mLocalDelivery.LocalDeliveryFee) vo.BaseLocalDeliveryFee {
			t := vo.BaseLocalDeliveryFee{}
			convertor.CopyProperties(&t, item)
			return t
		})
		out = append(out, temp)
	}

	return out, nil
}
