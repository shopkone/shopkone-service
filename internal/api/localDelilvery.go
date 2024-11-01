package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/delivery/local-delivery/sLocalDelivery/sLocalDelivery"
	"shopkone-service/internal/module/setting/location/sLocation"
	ctx2 "shopkone-service/utility/ctx"
)

type aLocalDelivery struct {
}

func NewLocalDeliveryApi() *aLocalDelivery {
	return &aLocalDelivery{}
}

func (a *aLocalDelivery) List(ctx g.Ctx, req *vo.LocalDeliveryListReq) (res []vo.LocalDeliveryListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return nil, err
	}
	shop := auth.Shop
	orm := sOrm.NewDb()
	locationIds, err := sLocation.NewLocation(orm, shop.ID).GetActiveIds()
	if err != nil {
		return nil, err
	}
	return sLocalDelivery.NewLocalDelivery(orm, shop.ID).LocalDeliveryList(locationIds)
}

func (a *aLocalDelivery) Info(ctx g.Ctx, req *vo.LocalDeliveryInfoReq) (res vo.LocalDeliveryInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	shop := auth.Shop
	return sLocalDelivery.NewLocalDelivery(sOrm.NewDb(), shop.ID).LocalDeliveryInfo(req.Id)
}

func (a *aLocalDelivery) Update(ctx g.Ctx, req *vo.UpdateLocalDeliveryReq) (res vo.UpdateLocalDeliveryRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sLocalDelivery.NewLocalDelivery(tx, shop.ID).LocalDeliveryUpdate(*req)
	})
	return res, err
}
