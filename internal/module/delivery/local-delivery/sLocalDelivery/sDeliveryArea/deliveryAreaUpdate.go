package sDeliveryArea

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/utility/handle"
)

func (s *sDeliveryArea) DeliveryAreaUpdate(areas []vo.BaseLocalDeliverArea, localDeliveryId uint) (err error) {
	// 获取旧的区域
	var oldAreas []mLocalDelivery.LocalDeliveryArea
	if err = s.orm.Model(&oldAreas).Where("local_delivery_id = ?", localDeliveryId).
		Where("shop_id = ?", s.shopId).
		Omit("updated_at", "created_at", "deleted_at").Find(&oldAreas).Error; err != nil {
		return err
	}

	// 组装新的区域
	var newAreas []mLocalDelivery.LocalDeliveryArea
	for _, area := range areas {
		temp := mLocalDelivery.LocalDeliveryArea{}
		temp.ID = area.Id
		temp.Note = area.Note
		temp.Name = area.Name
		temp.PostalCode = area.PostalCode
		temp.LocalDeliveryID = localDeliveryId
		temp.ShopId = s.shopId
		newAreas = append(newAreas, temp)
	}

	// 对比新老区域，删除不存在的区域，新增不存在的区域，更新存在的区域
	insert, update, remove, err := handle.DiffUpdate(newAreas, oldAreas)
	if err != nil {
		return err
	}

	// 删除
	removeIds := slice.Map(remove, func(_ int, item mLocalDelivery.LocalDeliveryArea) uint {
		return item.ID
	})
	if err = s.DeliveryAreaRemove(removeIds); err != nil {
		return err
	}

	// 添加
	if len(insert) > 0 {
		var insertIn []DeliveryAreaCreateItem
		for _, area := range insert {
			temp := DeliveryAreaCreateItem{}
			base, ok := slice.FindBy(areas, func(index int, item vo.BaseLocalDeliverArea) bool {
				return item.Id == area.ID
			})
			if !ok {
				return err
			}
			if err = convertor.CopyProperties(&temp, base); err != nil {
				return err
			}
			temp.LocalDeliveryID = localDeliveryId
			insertIn = append(insertIn, temp)
		}
		if err = s.DeliveryAreaCreate(insertIn); err != nil {
			return err
		}
	}

	// 更新
	if len(update) > 0 {
		//更新区域费用
		var fees []mLocalDelivery.LocalDeliveryFee
		slice.ForEach(update, func(index int, area mLocalDelivery.LocalDeliveryArea) {
			baseItem, ok := slice.FindBy(areas, func(index int, item vo.BaseLocalDeliverArea) bool {
				return item.Id == area.ID
			})
			if !ok {
				return
			}
			slice.ForEach(baseItem.Fees, func(_ int, f vo.BaseLocalDeliveryFee) {
				temp := mLocalDelivery.LocalDeliveryFee{}
				temp.Fee = f.Fee
				temp.LocalDeliveryAreaID = baseItem.Id
				temp.Condition = f.Condition
				temp.ShopId = s.shopId
				temp.ID = f.Id
				fees = append(fees, temp)
			})
		})
		if err = s.DeliveryFeeUpdate(fees); err != nil {
			return err
		}
		// 更新区域
		update = slice.Filter(update, func(index int, newItem mLocalDelivery.LocalDeliveryArea) bool {
			oldItem, ok := slice.FindBy(oldAreas, func(index int, old mLocalDelivery.LocalDeliveryArea) bool {
				return old.ID == newItem.ID
			})
			if !ok {
				return false
			}
			return s.DeliveryAreaIsChange(oldItem, newItem)
		})
		return s.DeliveryAreaUpdateBatch(update)
	}

	return err
}
