package api

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchase"
	ctx2 "shopkone-service/utility/ctx"
	"shopkone-service/utility/handle"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type aPurchase struct {
}

func NewPurchaseApi() *aPurchase {
	return &aPurchase{}
}

func (a *aPurchase) Create(ctx g.Ctx, req *vo.PurchaseCreateReq) (res vo.PurchaseCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res.Id, err = sPurchase.NewPurchase(tx, shop.ID).Create(*req)
		return err
	})
	return res, err
}

func (a *aPurchase) Update(ctx g.Ctx, req *vo.PurchaseUpdateReq) (res vo.PurchaseUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sPurchase.NewPurchase(tx, shop.ID).Update(*req)
	})
	return res, err
}

func (a *aPurchase) List(ctx g.Ctx, req *vo.PurchaseListReq) (res handle.PageRes[vo.PurchaseListRes], err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sPurchase.NewPurchase(sOrm.NewDb(&auth.Shop.ID), shop.ID).List(*req)
	return res, err
}

func (a *aPurchase) Info(ctx g.Ctx, req *vo.PurchaseInfoReq) (res vo.PurchaseInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sPurchase.NewPurchase(sOrm.NewDb(&auth.Shop.ID), shop.ID).Info(req.Id)
}

func (a *aPurchase) MarkToOrdered(ctx g.Ctx, req *vo.PurchaseMarkToOrderedReq) (res vo.PurchaseMarkToOrderedRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sPurchase.NewPurchase(tx, shop.ID).MarkToOrdered(req.Id)
	})
	return res, err
}

func (a *aPurchase) Remove(ctx g.Ctx, req *vo.PurchaseRemoveReq) (res vo.PurchaseRemoveRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sPurchase.NewPurchase(tx, shop.ID).Remove(req.Id)
	})
	return res, err
}

// 接收采购单
func (a *aPurchase) Adjust(ctx g.Ctx, req *vo.PurchaseAdjustReceiveReq) (res vo.PurchaseAdjustReceiveRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	user := auth.User
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sPurchase.NewPurchase(tx, shop.ID).Adjust(*req, user.Email)
	})
	return res, err
}

// 关闭/开启采购单
func (a *aPurchase) Close(ctx g.Ctx, req *vo.PurchaseCloseReq) (res vo.PurchaseCloseRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sPurchase.NewPurchase(tx, shop.ID).Close(req.Id, req.Close)
	})
	return res, err
}
