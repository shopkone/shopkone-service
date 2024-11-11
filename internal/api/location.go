package api

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/setting/location/sLocation"
	ctx2 "shopkone-service/utility/ctx"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type aLocation struct {
}

func NewLocationApi() *aLocation {
	return &aLocation{}
}

func (a *aLocation) List(ctx g.Ctx, req *vo.LocationListReq) (res []vo.LocationListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sLocation.NewLocation(sOrm.NewDb(&auth.Shop.ID), shop.ID).List(req.Active)
}

func (a *aLocation) LocationAdd(ctx g.Ctx, req *vo.LocationAddReq) (res vo.LocationAddRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res.Id, err = sLocation.NewLocation(tx, shop.ID).Create(*req, shop.TimeZone)
		return err
	})
	return res, err
}

func (a *aLocation) LocationInfo(ctx g.Ctx, req *vo.LocationInfoReq) (res vo.LocationInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sLocation.NewLocation(sOrm.NewDb(&auth.Shop.ID), shop.ID).Info(req.Id)
}

func (a *aLocation) LocationUpdate(ctx g.Ctx, req *vo.LocationUpdateReq) (res vo.LocationUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sLocation.NewLocation(tx, shop.ID).Update(*req)
	})
	return res, err
}

func (a *aLocation) LocationDelete(ctx g.Ctx, req *vo.DeleteLocationReq) (res vo.DeleteLocationRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sLocation.NewLocation(tx, shop.ID).Delete(req.Id)
		return err
	})
	return res, err
}

func (a *aLocation) LocationExistInventory(ctx g.Ctx, req *vo.LocationExistInventoryReq) (res vo.LocationExistInventoryRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res.Exist, err = sInventory.NewInventory(sOrm.NewDb(&auth.Shop.ID), shop.ID).ExistQuantityByLocationId(req.Id)
	return res, err
}

func (a *aLocation) LocationSetDefault(ctx g.Ctx, req *vo.SetDefaultLocationReq) (res vo.SetDefaultLocationRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sLocation.NewLocation(tx, shop.ID).LocationSetDefault(req.Id)
	})
	return res, err
}

func (a *aLocation) SetLocationOrder(ctx g.Ctx, req *vo.SetLocationOrderReq) (res vo.SetLocationOrderRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sLocation.NewLocation(tx, shop.ID).LocationSetOrder(req.Items)
	})
	return res, err
}
