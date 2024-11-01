package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/delivery/in-store-pick-up/sInStorePickup"
	"shopkone-service/internal/module/setting/location/sLocation"
	ctx2 "shopkone-service/utility/ctx"
)

type aInStorePickup struct {
}

func NewInStorePickup() *aInStorePickup {
	return &aInStorePickup{}
}

func (a *aInStorePickup) List(ctx g.Ctx, req *vo.InStorePickUpListReq) (res []vo.InStorePickUpListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return nil, err
	}
	shopId := auth.Shop.ID
	orm := sOrm.NewDb()
	locationIds, err := sLocation.NewLocation(orm, shopId).GetActiveIds()
	if err != nil {
		return nil, err
	}
	return sInStorePickup.NewInStorePickup(orm, shopId).List(locationIds)
}

func (a *aInStorePickup) Info(ctx g.Ctx, req *vo.InStorePickUpInfoReq) (res vo.InStorePickUpInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	return sInStorePickup.NewInStorePickup(sOrm.NewDb(), auth.Shop.ID).Info(req.Id)
}

func (a *aInStorePickup) Update(ctx g.Ctx, req *vo.InStorePickUpUpdateReq) (res vo.InStorePickUpUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	if err != nil {
		return
	}
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sInStorePickup.NewInStorePickup(tx, shop.ID).Update(*req)
	})
	return res, err
}
