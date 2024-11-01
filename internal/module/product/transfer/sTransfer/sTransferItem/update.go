package sTransferItem

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/utility/handle"
)

type ItemUpdateIn struct {
	Items         []vo.BaseTransferItem
	TransferId    uint
	OriginId      uint
	DestinationId uint
}

func (s *sTransferItem) Update(in ItemUpdateIn) error {
	// 获取旧的items
	oldItems, err := s.ListByTransferIds([]uint{in.TransferId})
	if err != nil {
		return err
	}

	// 组装新的items
	newItems := slice.Map(in.Items, func(_ int, item vo.BaseTransferItem) mTransfer.TransferItem {
		i := mTransfer.TransferItem{}
		i.ID = item.Id
		i.Quantity = item.Quantity
		i.VariantId = item.VariantId
		i.ShopId = s.shopId
		return i
	})

	insert, update, remove, err := handle.DiffUpdate(newItems, oldItems)

	if len(insert) > 0 {
		baseInsert := slice.Map(insert, func(_ int, item mTransfer.TransferItem) vo.BaseTransferItem {
			find, ok := slice.FindBy(in.Items, func(index int, i vo.BaseTransferItem) bool {
				return i.Id == item.ID
			})
			if !ok {
				return vo.BaseTransferItem{}
			}
			return find
		})
		baseInsert = slice.Filter(baseInsert, func(index int, item vo.BaseTransferItem) bool {
			return item.Id > 0
		})
		createIn := CreateTransferItemsIn{
			TransferId:    in.TransferId,
			OriginId:      in.OriginId,
			DestinationId: in.DestinationId,
			Items:         baseInsert,
		}
		if err = s.CreateTransferItems(createIn); err != nil {
			return err
		}
	}

	if len(update) > 0 {
		update = slice.Filter(update, func(index int, item mTransfer.TransferItem) bool {
			oldItem, ok := slice.FindBy(oldItems, func(index int, i mTransfer.TransferItem) bool {
				return i.ID == item.ID
			})
			if !ok {
				return false
			}
			return item.Quantity != oldItem.Quantity
		})
		update = slice.Map(update, func(_ int, item mTransfer.TransferItem) mTransfer.TransferItem {
			item.CanCreateId = true
			return item
		})
		batchUpdate := handle.BatchUpdateByIdIn{
			Orm:    s.orm,
			ShopID: s.shopId,
			Query:  []string{"quantity"},
		}
		if err = handle.BatchUpdateById(batchUpdate, &update); err != nil {
			return err
		}
	}

	if len(remove) > 0 {
		removeIds := slice.Map(remove, func(_ int, item mTransfer.TransferItem) uint {
			return item.ID
		})
		if err = s.RemoveByIds(removeIds); err != nil {
			return err
		}
	}

	return nil
}
