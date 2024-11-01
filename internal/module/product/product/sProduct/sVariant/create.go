package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	sTransfer2 "shopkone-service/internal/module/product/product/sProduct/sTransfer"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/utility/code"
)

func (s *sVariant) Create(in iProduct.VariantCreateIn) (list []mProduct.Variant, err error) {
	if len(in.List) == 0 {
		return nil, code.NoEmptyVariants
	}
	// 校验变体名称是否重复
	var names []string
	slice.ForEach(in.List, func(index int, item vo.BaseVariant) {
		ns := slice.Map(item.Name, func(index int, item mProduct.VariantName) string {
			return item.Value + item.Label
		})
		names = append(names, slice.Join(ns, ""))
	})
	if len(slice.Unique(names)) != len(names) {
		return list, code.VariantNameRepeat
	}
	// 创建变体
	sTransfer := sTransfer2.NewProductTransfer(s.shopId)
	list = slice.Map(in.List, func(index int, item vo.BaseVariant) mProduct.Variant {
		data := sTransfer.VariantToModel(item, in.ProductId, false)
		data.ShopId = s.shopId
		return data
	})
	if err = s.orm.Create(&list).Error; err != nil {
		return nil, err
	}
	// 如果不追踪库存，则不创建库存
	if !in.InventoryTracking {
		return list, nil
	}
	// 创建库存
	genIn := GenIn{
		Variants:           list,
		Base:               in.List,
		EnabledLocationIds: in.EnableLocationIds,
	}
	inventories, err := s.GenInventoriesByVariants(genIn)
	if err != nil {
		return nil, err
	}
	// 校验locationId是否都存在且启用
	locationIds := slice.Map(inventories, func(_ int, item iInventory.CreateInventoryIn) uint {
		return item.LocationId
	})
	if err = sLocation.NewLocation(s.orm, s.shopId).IsAllActive(locationIds); err != nil {
		return list, err
	}
	// 创建库存
	return list, sInventory.NewInventory(s.orm, s.shopId).Create(inventories, in.Type, in.HandleEmail).Error
}
