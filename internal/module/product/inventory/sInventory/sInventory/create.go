package sInventory

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventoryChanger"
	"shopkone-service/utility/code"
)

func (s *sInventory) Create(in []iInventory.CreateInventoryIn, inType mInventory.InventoryType, email string) *sInventory {
	if len(in) == 0 {
		return s
	}
	// 校验变体id是否都存在
	noZeroVariantList := slice.Filter(in, func(_ int, item iInventory.CreateInventoryIn) bool {
		return item.VariantId != 0
	})
	if len(noZeroVariantList) != len(in) {
		s.Error = code.VariantSomeIdRequired
		return s
	}
	// 创建库存
	s.data = slice.Map(in, func(_ int, item iInventory.CreateInventoryIn) mInventory.Inventory {
		temp := mInventory.Inventory{}
		temp.ShopId = s.shopId
		temp.Quantity = item.Quantity
		temp.VariantId = item.VariantId
		temp.LocationId = item.LocationId
		return temp
	})
	if s.Error = s.orm.Create(&s.data).Error; s.Error != nil {
		return s
	}
	// 创建库存记录
	inventories := slice.Map(s.data, func(index int, item mInventory.Inventory) mInventory.InventoryChange {
		temp := mInventory.InventoryChange{}
		temp.HandleEmail = email
		temp.ShopId = s.shopId
		temp.InventoryId = item.ID
		temp.Type = inType
		temp.DiffQuantity = int(item.Quantity)
		temp.Style = mInventory.InventoryStyleOut
		temp.InventoryId = item.ID
		return temp
	})
	if s.Error = sInventoryChanger.NewInventoryChange(s.orm, s.shopId).AddHistory(inventories); s.Error != nil {
		return s
	}
	return s
}
