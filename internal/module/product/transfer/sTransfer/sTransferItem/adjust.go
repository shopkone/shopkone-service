package sTransferItem

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/utility/handle"
)

func (s *sTransferItem) Adjust(in vo.TransferAdjustReq) error {
	// 过滤不为0的
	items := slice.Filter(in.Items, func(index int, item vo.TransferAdjustItem) bool {
		return item.Rejected > 0 || item.Received > 0
	})
	if len(items) == 0 {
		return nil
	}

	// 获取旧的数据
	variantIds := slice.Map(items, func(index int, item vo.TransferAdjustItem) uint {
		return item.VariantID
	})
	variantIds = slice.Unique(variantIds)
	variantIds = slice.Filter(variantIds, func(index int, id uint) bool {
		return id > 0
	})
	var oldItems []mTransfer.TransferItem
	if err := s.orm.Model(&oldItems).Where("shop_id = ? AND transfer_id = ? AND variant_id IN ?", s.shopId, in.Id, variantIds).
		Omit("created_at", "updated_at", "deleted_at").Find(&oldItems).Error; err != nil {
		return err
	}

	// 更新数据
	oldItems = slice.Map(oldItems, func(index int, item mTransfer.TransferItem) mTransfer.TransferItem {
		item.Received += items[index].Received
		item.Rejected += items[index].Rejected
		item.CanCreateId = true
		return item
	})
	diffUpdateIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"received", "rejected"},
	}
	if err := handle.BatchUpdateById(diffUpdateIn, &oldItems); err != nil {
		return err
	}
	return nil
}
