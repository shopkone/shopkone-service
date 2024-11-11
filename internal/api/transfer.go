package api

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransfer"
	ctx2 "shopkone-service/utility/ctx"
	"shopkone-service/utility/handle"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type aTransfer struct {
}

func NewTransferApi() *aTransfer {
	return &aTransfer{}
}

func (a *aTransfer) CreateTransfer(ctx g.Ctx, req *vo.TransferCreateReq) (res vo.TransferCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res.Id, err = sTransfer.NewTransfer(tx, shop.ID).CreateTransfer(*req)
		return err
	})
	return res, err
}

func (a *aTransfer) List(ctx g.Ctx, req *vo.TransferListReq) (res handle.PageRes[vo.TransferListRes], err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sTransfer.NewTransfer(sOrm.NewDb(&auth.Shop.ID), shop.ID).List(*req)
	return res, err
}

func (a *aTransfer) Info(ctx g.Ctx, req *vo.TransferInfoReq) (res vo.TransferInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sTransfer.NewTransfer(sOrm.NewDb(&auth.Shop.ID), shop.ID).Info(req.Id)
	return res, err
}

func (a *aTransfer) Mark(ctx g.Ctx, req *vo.TransferMarkReq) (res vo.TransferInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sTransfer.NewTransfer(tx, shop.ID).Mark(req.Id)
	})
	return res, err
}

func (a *aTransfer) Adjust(ctx g.Ctx, req *vo.TransferAdjustReq) (res vo.TransferAdjustRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sTransfer.NewTransfer(tx, shop.ID).Adjust(*req)
	})
	return res, err
}

func (a *aTransfer) Remove(ctx g.Ctx, req *vo.TransferRemoveReq) (res vo.TransferRemoveRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sTransfer.NewTransfer(tx, shop.ID).RemoveByIds([]uint{req.Id})
	})
	return res, err
}

func (a *aTransfer) Update(ctx g.Ctx, req *vo.TransferUpdateReq) (res vo.TransferUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sTransfer.NewTransfer(tx, shop.ID).Update(*req)
	})
	return res, err
}
