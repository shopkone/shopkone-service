package sShippingZoneFee

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/consts"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/code"
)

type FeeCreateItem struct {
	Name           string                        `json:"name"`
	WeightUnit     consts.WeightUnit             `json:"weight_unit"`
	Type           mShipping.ShippingZoneFeeType `json:"type"`
	CurrencyCode   string                        `json:"currency_code"`
	Rule           mShipping.ShippingZoneFeeRule `json:"rule"`
	Remark         string                        `json:"remark"`
	Cod            bool                          `json:"cod"`
	ShippingZoneId uint                          `json:"shipping_zone_id"`
	Conditions     []ConditionCreateItem         `json:"conditions"`
}

func (s *sShippingZoneFee) FeeCreate(in []FeeCreateItem) (err error) {
	// 创建费用
	var fees []mShipping.ShippingZoneFee
	for _, fee := range in {
		temp := mShipping.ShippingZoneFee{}
		if err = convertor.CopyProperties(&temp, fee); err != nil {
			return
		}
		temp.ShopId = s.shopId
		fees = append(fees, temp)
	}
	if err = s.orm.Create(&fees).Error; err != nil {
		return err
	}

	// 创建规则
	var conditions []ConditionCreateItem
	for _, fee := range fees {
		findIn, ok := slice.FindBy(in, func(index int, item FeeCreateItem) bool {
			return item.Name == fee.Name
		})
		if !ok {
			return code.IdMissing
		}
		for _, condition := range findIn.Conditions {
			condition.ShippingZoneFeeId = fee.ID
			conditions = append(conditions, condition)
		}
	}

	return s.ConditionCreate(conditions)
}
