package sPurchaseItem

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

type PurchaseItemUpdateIn struct {
	Items      []vo.BasePurchaseItem
	PurchaseId uint
}

func (s *sPurchaseItem) Update(in PurchaseItemUpdateIn) (err error) {
	// 获取之前的list
	oldList, err := s.ListByPurchaseId(in.PurchaseId)
	if err != nil {
		return err
	}

	// 获取新的list
	newList := slice.Map(in.Items, func(index int, item vo.BasePurchaseItem) mPurchase.PurchaseItem {
		i := mPurchase.PurchaseItem{
			Cost:       item.Cost,
			Purchasing: item.Purchasing,
			SKU:        item.SKU,
			TaxRate:    item.TaxRate,
			VariantID:  item.VariantID,
			PurchaseId: in.PurchaseId,
		}
		i.ShopId = s.shopId
		i.ID = item.Id
		return i
	})

	insert, update, remove, err := handle.DiffUpdate(newList, oldList)
	if err != nil {
		return err
	}

	// 判断被删除的是否有拒收或者收货的，有则不被允许
	_, hasNo := slice.FindBy(remove, func(index int, item mPurchase.PurchaseItem) bool {
		return item.Received > 0 || item.Rejected > 0
	})
	if hasNo {
		return code.PurchaseItemCanNotRemoveByReceivedOrRejected
	}

	// 插入
	if len(insert) > 0 {
		insertList := slice.Map(insert, func(index int, item mPurchase.PurchaseItem) vo.BasePurchaseItem {
			find, ok := slice.FindBy(in.Items, func(index int, i vo.BasePurchaseItem) bool {
				return item.VariantID == i.VariantID
			})
			if !ok {
				return vo.BasePurchaseItem{}
			}
			find.Id = 0
			return find
		})
		if err = s.Create(insertList, in.PurchaseId); err != nil {
			return err
		}
	}

	// 更新
	if len(update) > 0 {
		update = slice.Map(update, func(index int, item mPurchase.PurchaseItem) mPurchase.PurchaseItem {
			item.CanCreateId = true
			return item
		})
		update = slice.Filter(update, func(index int, i mPurchase.PurchaseItem) bool {
			find, ok := slice.FindBy(oldList, func(index int, old mPurchase.PurchaseItem) bool {
				return old.VariantID == i.VariantID
			})
			if !ok {
				return false
			}
			return s.IsChange(find, i)
		})
		if len(update) > 0 {
			batchIn := handle.BatchUpdateByIdIn{
				Orm:    s.orm,
				ShopID: s.shopId,
				Query:  []string{"sku", "purchasing", "cost", "tax_rate"},
			}
			if err = handle.BatchUpdateById(batchIn, &update); err != nil {
				return err
			}
		}
	}

	// 移除
	removeIds := slice.Map(remove, func(index int, item mPurchase.PurchaseItem) uint {
		return item.ID
	})
	if len(remove) > 0 {
		err = s.orm.Where("id in ?", removeIds).Delete(&mPurchase.PurchaseItem{}).Error
	}

	return err
}
