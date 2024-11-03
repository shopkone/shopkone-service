package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sCustomerTax) CustomerTaxUpdate(in []vo.BaseCustomerTax, taxId uint) (err error) {
	// 获取旧的自定义区域税率
	var oldCustomerTax []mTax.CustomerTax
	if err = s.orm.Model(&oldCustomerTax).
		Where("shop_id = ? AND tax_id = ?", s.shopId, taxId).
		Omit("created_at", "updated_at", "deleted_at").Find(&oldCustomerTax).Error; err != nil {
		return err
	}

	// 新的自定义税率
	newCustomerTax := slice.Map(in, func(index int, item vo.BaseCustomerTax) mTax.CustomerTax {
		i := mTax.CustomerTax{}
		i.ID = item.ID
		i.ShopId = s.shopId
		i.TaxID = taxId
		i.Type = item.Type
		i.CollectionID = item.CollectionID
		return i
	})

	// 校验如果是collection，必须不能为0
	someZero := slice.Some(newCustomerTax, func(index int, item mTax.CustomerTax) bool {
		if item.Type == mTax.CustomerTaxTypeCollection && item.CollectionID == 0 {
			return true
		}
		return false
	})
	if someZero {
		return code.TaxCollectionMust
	}

	// 差异
	insert, update, remove, err := handle.DiffUpdate(newCustomerTax, oldCustomerTax)
	if err != nil {
		return err
	}
	if err = s.CustomerTaxCreate(insert, in); err != nil {
		return err
	}
	if err = s.CustomerTaxRemove(remove); err != nil {
		return err
	}
	if err = s.CustomerTaxUpdateBatch(update, oldCustomerTax, in); err != nil {
		return err
	}

	return err
}
