package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShipping"
	ctx2 "shopkone-service/utility/ctx"
)

type aShipping struct {
}

func NewShippingApi() *aShipping {
	return &aShipping{}
}

func (a *aShipping) Create(ctx g.Ctx, req *vo.ShippingCreateReq) (res vo.ShippingCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return vo.ShippingCreateRes{}, err
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		res.ID, err = sShipping.NewShipping(tx, shop.ID).ShippingCreate(*req)
		return err
	})
	return res, err
}

func (a *aShipping) Update(ctx g.Ctx, req *vo.ShippingUpdateReq) (res vo.ShippingUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return vo.ShippingUpdateRes{}, err
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sShipping.NewShipping(tx, shop.ID).ShippingUpdate(*req)
	})
	return res, err
}

func (a *aShipping) Info(ctx g.Ctx, req *vo.ShippingInfoReq) (res vo.ShippingInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return vo.ShippingInfoRes{}, err
	}
	shop := auth.Shop
	res.BaseShipping, err = sShipping.NewShipping(sOrm.NewDb(), shop.ID).ShippingInfo(req.ID)
	return res, err
}

func (a *aShipping) List(ctx g.Ctx, req *vo.ShippingListReq) (res []vo.ShippingListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sShipping.NewShipping(sOrm.NewDb(), shop.ID).ShippingList()
}
