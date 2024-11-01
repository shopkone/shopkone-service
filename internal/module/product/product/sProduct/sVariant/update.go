package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	sTransfer2 "shopkone-service/internal/module/product/product/sProduct/sTransfer"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sVariant) Update(in iProduct.VariantUpdateIn) ([]mProduct.Variant, error) {
	if len(in.List) <= 0 {
		return []mProduct.Variant{}, code.NoEmptyVariants
	}
	productId := in.ProductId
	trackInventory := in.TrackInventory
	list := in.List
	// 获取旧的变体
	var oldVariants []mProduct.Variant
	err := s.orm.Model(&oldVariants).Where("shop_id = ? AND product_id = ?", s.shopId, productId).
		Find(&oldVariants).Error
	if err != nil {
		return oldVariants, err
	}
	// 如果不跟踪库存，则删除所有的库存
	if !trackInventory {
		err = sInventory.NewInventory(s.orm, s.shopId).RemoveByVariantIds(slice.Map(oldVariants, func(index int, item mProduct.Variant) uint {
			return item.ID
		}), 0)
		if err != nil {
			return oldVariants, err
		}
	}
	// 获取新的变体
	sTransfer := sTransfer2.NewProductTransfer(s.shopId)
	newVariants := slice.Map(list, func(index int, item vo.BaseVariantWithId) mProduct.Variant {
		i := sTransfer.VariantToModel(item.BaseVariant, productId, true)
		i.ShopId = s.shopId
		i.ID = item.Id
		return i
	})
	// 差异
	insert, update, remove, err := handle.DiffUpdate(newVariants, oldVariants)
	if err != nil {
		return newVariants, err
	}
	// 更新
	if len(update) > 0 {
		// 如果跟踪库存，则更新库存
		if err = s.UpdateInventoryWithUpdateVariants(update, list, in.TrackInventory, in.HandleEmail); err != nil {
			return newVariants, err
		}
		// 更新变体
		if err = s.UpdateVariants(update, oldVariants); err != nil {
			return newVariants, err
		}
	}
	// 删除
	if len(remove) > 0 {
		removeIds := slice.Map(remove, func(index int, item mProduct.Variant) uint {
			return item.ID
		})
		if err = s.RemoveByIds(removeIds); err != nil {
			return newVariants, err
		}
	}
	// 插入
	if len(insert) > 0 {
		i := slice.Map(insert, func(index int, item mProduct.Variant) vo.BaseVariant {
			temp := sTransfer.ModelToVariant(item)
			temp.Id = 0
			return temp
		})
		createIn := iProduct.VariantCreateIn{
			List:              i,
			ProductId:         productId,
			InventoryTracking: trackInventory,
			Type:              mInventory.InventoryChangeProduct,
			EnableLocationIds: in.EnableLocationIds,
		}
		if _, err = s.Create(createIn); err != nil {
			return newVariants, err
		}
	}
	return newVariants, err
}
