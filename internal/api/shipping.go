package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZone"
	ctx2 "shopkone-service/utility/ctx"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
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
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
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
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
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
	res.BaseShipping, err = sShipping.NewShipping(sOrm.NewDb(&auth.Shop.ID), shop.ID).ShippingInfo(req.ID)
	return res, err
}

func (a *aShipping) List(ctx g.Ctx, req *vo.ShippingListReq) (res []vo.ShippingListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sShipping.NewShipping(sOrm.NewDb(&auth.Shop.ID), shop.ID).ShippingList()
}

func (a *aShipping) ZoneListByCountries(ctx g.Ctx, req *vo.ShippingZoneListByCountriesReq) (res []vo.ShippingZoneListByCountriesRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	db := sOrm.NewDb(&shop.ID)
	// 获取区域
	s := sShippingZone.NewShippingZone(shop.ID, db)
	zones, err := s.ZonesByCountries(req.CountryCodes)
	if zones == nil {
		zones = []vo.BaseShippingZone{}
	}
	// 获取所属方案信息
	shippingIds := slice.Map(zones, func(index int, item vo.BaseShippingZone) uint {
		return item.ShippingID
	})
	shippingIds = slice.Unique(shippingIds)
	shippings, err := sShipping.NewShipping(db, shop.ID).SimpleList(shippingIds)
	if err != nil {
		return res, err
	}
	res = slice.Map(zones, func(index int, zone vo.BaseShippingZone) vo.ShippingZoneListByCountriesRes {
		i := vo.ShippingZoneListByCountriesRes{}
		i.BaseShippingZone = zone
		shipping, ok := slice.FindBy(shippings, func(index int, shipping mShipping.Shipping) bool {
			return shipping.ID == zone.ShippingID
		})
		if !ok {
			return i
		}
		i.ShippingID = shipping.ID
		i.ShippingName = shipping.Name
		return i
	})
	return res, err
}
