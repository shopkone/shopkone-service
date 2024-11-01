package sDeliveryArea

import (
	"github.com/duke-git/lancet/v2/convertor"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
)

type DeliveryAreaCreateItem struct {
	Name            string                    `json:"name"`        // 区域名称
	PostalCode      string                    `json:"postal_code"` // 邮编
	Note            string                    `json:"note"`        // 备注
	LocalDeliveryID uint                      `json:"local_delivery_id"`
	Fees            []vo.BaseLocalDeliveryFee `json:"fees"`
}

func (s *sDeliveryArea) DeliveryAreaCreate(in []DeliveryAreaCreateItem) (err error) {
	if len(in) == 0 {
		return err
	}
	// 创建区域
	var data []mLocalDelivery.LocalDeliveryArea
	for _, item := range in {
		temp := mLocalDelivery.LocalDeliveryArea{}
		if err = convertor.CopyProperties(&temp, item); err != nil {
			return err
		}
		temp.ShopId = s.shopId
		data = append(data, temp)
	}
	if err = s.orm.Create(&data).Error; err != nil {
		return err
	}

	// 创建配送费用
	var fees []mLocalDelivery.LocalDeliveryFee
	for i, item := range in {
		find := data[i]
		for _, fee := range item.Fees {
			temp := mLocalDelivery.LocalDeliveryFee{}
			temp.ShopId = s.shopId
			temp.Condition = fee.Condition
			temp.Fee = fee.Fee
			temp.LocalDeliveryAreaID = find.ID
			fees = append(fees, temp)
		}
	}
	return s.orm.Create(&fees).Error
}
