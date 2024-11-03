package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/handle"
)

func (s *sCustomerTax) CustomerZoneUpdate(zones []mTax.CustomerTaxZone, customerIds []uint) (err error) {
	if len(customerIds) == 0 {
		return err
	}
	// 获取旧的区域税率
	var oldZones []mTax.CustomerTaxZone
	if err = s.orm.Model(&oldZones).Where("shop_id = ? AND customer_tax_id IN ?", s.shopId, customerIds).
		Omit("created_at", "updated_at", "deleted_at").Find(&oldZones).Error; err != nil {
		return err
	}
	// 对比差异
	insert, update, remove, err := handle.DiffUpdate(zones, oldZones)
	if err != nil {
		return err
	}
	if err = s.CustomerZoneCreate(insert); err != nil {
		return err
	}
	removeIds := slice.Map(remove, func(index int, item mTax.CustomerTaxZone) uint {
		return item.ID
	})
	if err = s.CustomerZoneRemoveIds(removeIds); err != nil {
		return err
	}
	return s.TaxZoneUpdateBatch(update, oldZones)
}
