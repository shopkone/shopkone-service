package sInventory

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/utility/handle"
)

func (s *sInventory) List(in vo.InventoryListReq) (out handle.PageRes[vo.InventoryListRes], err error) {
	var list []mInventory.Inventory
	query := s.orm.Model(&list).Where("shop_id = ?", s.shopId)
	query = query.Where("location_id = ?", in.LocationId)
	// 获取count
	if err = query.Count(&out.Total).Error; err != nil {
		return out, err
	}
	// 分页
	query = query.Scopes(handle.Pagination(in.PageReq)).Order("id desc")
	// 获取库存列表
	query = query.Select("id", "variant_id", "quantity")
	if err = query.Find(&list).Error; err != nil {
		return out, err
	}
	out.List = slice.Map(list, func(index int, item mInventory.Inventory) vo.InventoryListRes {
		temp := vo.InventoryListRes{}
		temp.Id = item.ID
		temp.Quantity = item.Quantity
		temp.VariantId = item.VariantId
		return temp
	})
	return out, err
}
