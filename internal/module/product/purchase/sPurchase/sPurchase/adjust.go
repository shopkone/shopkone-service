package sPurchase

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/sProduct/sProduct"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sPurchase) Adjust(in vo.PurchaseAdjustReceiveReq, email string) (err error) {
	// 筛出不为0的items
	items := slice.Filter(in.Items, func(index int, item vo.PurchaseAdjustReceiveItem) bool {
		return item.RejectedCount != 0 || item.ReceivedCount != 0
	})
	if len(items) == 0 {
		return err
	}
	itemIds := slice.Map(items, func(index int, item vo.PurchaseAdjustReceiveItem) uint {
		return item.Id
	})

	// 查找采购单
	var purchase mPurchase.Purchase
	if err = s.orm.Model(&purchase).Where("shop_id = ? AND id = ?", s.shopId, in.Id).
		Select("id", "destination_id", "supplier_id", "status").First(&purchase).Error; err != nil {
		return err
	}

	// 只有三种状态可以修改
	if purchase.Status != mPurchase.PurchaseStatusOrdered && // 已下单
		purchase.Status != mPurchase.PurchaseStatusPartialReceived && // 部分收货
		purchase.Status != mPurchase.PurchaseStatusReceived { // 全部收货
		return code.SupplierAdjustError
	}

	// 查找旧的采购单项
	var oldItems []mPurchase.PurchaseItem
	if err = s.orm.Model(&oldItems).Where("purchase_id = ? AND shop_id = ? AND id IN ?", in.Id, s.shopId, itemIds).
		Select("id", "variant_id", "received", "rejected").Find(&oldItems).Error; err != nil {
		return err
	}
	if len(oldItems) == 0 {
		return err
	}
	variantIds := slice.Map(oldItems, func(index int, item mPurchase.PurchaseItem) uint {
		return item.VariantID
	})

	// 根据变体id获取变体列表
	variants, err := sVariant.NewVariant(s.orm, s.shopId).ListByIds(variantIds, false)
	// 筛出没有的变体id
	variantIds = slice.Filter(variantIds, func(index int, item uint) bool {
		_, ok := slice.FindBy(variants, func(index int, v iProduct.VariantListByIdOut) bool {
			return v.Id == item
		})
		return ok
	})
	oldItems = slice.Filter(oldItems, func(index int, item mPurchase.PurchaseItem) bool {
		_, ok := slice.FindBy(variants, func(index int, v iProduct.VariantListByIdOut) bool {
			return v.Id == item.VariantID
		})
		return ok
	})

	// 是否全部跟踪库存
	is, err := sProduct.NewProduct(s.orm, s.shopId).IsAllTrack(variantIds)
	if !is {
		return code.ProductNotTrackInventory
	}

	// 调整库存
	inventoryService := sInventory.NewInventory(s.orm, s.shopId)
	// 根据变体id获取库存列表
	inventories, err := inventoryService.ListByVariantsIds(variantIds, purchase.DestinationId)
	if err != nil {
		return err
	}
	// 调整库存
	newInventories := slice.Map(inventories, func(index int, item mInventory.Inventory) mInventory.Inventory {
		oldItem, ok := slice.FindBy(oldItems, func(index int, i mPurchase.PurchaseItem) bool {
			return i.VariantID == item.VariantId
		})
		if !ok {
			return item
		}
		find, ok := slice.FindBy(items, func(index int, i vo.PurchaseAdjustReceiveItem) bool {
			return i.Id == oldItem.ID
		})
		if !ok {
			return item
		}
		item.Quantity = item.Quantity + uint(find.ReceivedCount)
		item.ShopId = s.shopId
		return item
	})
	// 更新库存
	updateIn := iInventory.UpdateByDiffIn{
		News:        newInventories,
		Olds:        inventories,
		HandleEmail: email,
		UpdateType:  mInventory.InventoryChangePurchase,
	}
	if err = inventoryService.UpdateByDiff(updateIn); err != nil {
		return err
	}

	// 更新采购单
	newItems := slice.Map(oldItems, func(index int, item mPurchase.PurchaseItem) mPurchase.PurchaseItem {
		find, ok := slice.FindBy(items, func(index int, i vo.PurchaseAdjustReceiveItem) bool {
			return i.Id == item.ID
		})
		if !ok {
			return item
		}
		item.Received = find.ReceivedCount + item.Received
		item.Rejected = find.RejectedCount + item.Rejected
		item.ShopId = s.shopId
		item.CanCreateId = true
		return item
	})
	batchUpdateIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"received", "rejected"},
	}
	if err = handle.BatchUpdateById(batchUpdateIn, &newItems); err != nil {
		return err
	}

	// 更新采购单状态
	return s.SetStatus(in.Id, purchase.Status)
}
