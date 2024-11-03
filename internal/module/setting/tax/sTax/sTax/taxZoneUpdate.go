package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sTax) TaxZoneUpdate(in []vo.BaseTaxZone, taxId uint) (err error) {
	// 获取旧的区域税率
	var oldZones []mTax.TaxZone
	if err = s.orm.Model(&oldZones).Where("tax_id = ?", taxId).Where("shop_id = ?", s.shopId).
		Omit("created_at", "updated_at", "updated_at").Find(&oldZones).Error; err != nil {
		return err
	}

	// 获取新的区域税率
	newZones := slice.Map(in, func(index int, item vo.BaseTaxZone) mTax.TaxZone {
		i := mTax.TaxZone{}
		i.ShopId = s.shopId
		i.TaxId = taxId
		i.ID = item.ID
		i.Name = item.Name
		i.ZoneCode = item.ZoneCode
		i.TaxRate = item.TaxRate
		return i
	})

	// 对比差异
	insert, update, remove, err := handle.DiffUpdate(newZones, oldZones)
	if err != nil {
		return err
	}

	// 判断不能出现相同的code
	list := slice.Concat(insert, update)
	listCodes := slice.Map(list, func(index int, item mTax.TaxZone) string {
		return item.ZoneCode
	})
	listCodes = slice.Unique(listCodes)
	if len(listCodes) != len(list) {
		return code.TaxZooeCodeRepeat
	}

	if err = s.TaxZoneCreate(insert); err != nil {
		return err
	}
	if err = s.TaxZoneRemove(remove); err != nil {
		return err
	}
	if err = s.TaxZoneUpdateBatch(update, oldZones); err != nil {
		return err
	}
	return err
}
