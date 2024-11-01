package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventoryChanger"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/sProduct/sProduct"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
	"shopkone-service/internal/module/setting/location/sLocation"
	ctx2 "shopkone-service/utility/ctx"
	"shopkone-service/utility/handle"
)

type aInventory struct {
}

func NewInventoryApi() *aInventory {
	return &aInventory{}
}

// List 库存列表
func (a *aInventory) List(ctx g.Ctx, req *vo.InventoryListReq) (res handle.PageRes[vo.InventoryListRes], err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb()

	// 获取库存列表
	res, err = sInventory.NewInventory(orm, shop.ID).List(*req)
	if err != nil {
		return res, err
	}

	// 根据库存获取变体列表
	variantIds := slice.Map(res.List, func(index int, item vo.InventoryListRes) uint {
		return item.VariantId
	})
	variants, err := sVariant.NewVariant(orm, shop.ID).ListByIds(variantIds, false)
	if err != nil {
		return res, err
	}

	// 获取商品列表
	productIds := slice.Map(variants, func(index int, item iProduct.VariantListByIdOut) uint {
		return item.ProductId
	})
	productIds = slice.Unique(productIds)
	products, err := sProduct.NewProduct(orm, shop.ID).ListByIdsWithoutVariants(productIds)
	if err != nil {
		return res, err
	}

	// 组装数据
	res.List = slice.Map(res.List, func(index int, item vo.InventoryListRes) vo.InventoryListRes {
		// 获取变体信息
		variant, ok := slice.FindBy(variants, func(index int, v iProduct.VariantListByIdOut) bool {
			return item.VariantId == v.Id
		})
		if !ok {
			return item
		}
		item.Name = variant.Name
		item.Image = variant.Image
		item.Sku = variant.Sku
		// 获取商品信息
		product, exist := slice.FindBy(products, func(index int, p iProduct.ListByIdsWithoutVariantsOut) bool {
			return variant.ProductId == p.Id
		})
		if !exist {
			return item
		}
		item.ProductName = product.Title
		if item.Name == "" { // 说明是独立变体
			item.Image = product.Image
		}
		return item
	})

	return res, err
}

// Move 库存转移
func (a *aInventory) Move(ctx g.Ctx, req *vo.InventoryMoveReq) (res vo.InventoryMoveRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	user := auth.User
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sInventory.NewInventory(tx, shop.ID).MoveInventory(req.From, req.To, user.Email)
	})
	return res, err
}

func (a *aInventory) HistoryList(ctx g.Ctx, req *vo.InventoryHistoryReq) (res []vo.InventoryHistoryRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sInventoryChanger.NewInventoryChange(sOrm.NewDb(), shop.ID).List(req.Id)
	return res, err
}

func (a *aInventory) InventoryListUnByVariantIds(ctx g.Ctx, req *vo.InventoryListUnByVariantIdsReq) (res []vo.InventoryListUnByVariantIdsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb()
	// 获取可用的地点
	is := true
	locations, err := sLocation.NewLocation(orm, shop.ID).List(&is)
	if err != nil {
		return res, err
	}
	locationIds := slice.Map(locations, func(index int, item vo.LocationListRes) uint {
		return item.Id
	})
	if len(locationIds) == 0 {
		return res, err
	}
	// 获取被删除的库存
	inventories, err := sInventory.NewInventory(orm, shop.ID).InventoryListUnByVariantIds(req.Ids, locationIds)
	if err != nil {
		return res, err
	}
	// 转换数据
	res = slice.Map(inventories, func(index int, item mInventory.Inventory) vo.InventoryListUnByVariantIdsRes {
		return vo.InventoryListUnByVariantIdsRes{
			Id:         item.ID,
			LocationId: item.LocationId,
			Quantity:   item.Quantity,
			VariantId:  item.VariantId,
		}
	})
	return res, err
}
