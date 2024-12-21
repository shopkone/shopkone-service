package api

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/file/sFile"
	"shopkone-service/internal/module/shop/shop/mShop"
	"shopkone-service/internal/module/shop/shop/sShop"
	"shopkone-service/internal/module/shop/transaction/sTransaction"
	ctx2 "shopkone-service/utility/ctx"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type aShopApi struct{}

func NewShopApi() *aShopApi {
	return &aShopApi{}
}

func (a *aShopApi) Info(ctx g.Ctx, req *vo.ShopInfoReq) (res vo.ShopInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	res.Uuid = auth.Shop.Uuid
	res.StoreName = auth.Shop.StoreName
	res.Status = auth.Shop.Status
	res.StoreCurrency = auth.Shop.StoreCurrency
	res.TimeZone = auth.Shop.TimeZone
	res.Country = auth.Shop.Country
	files, err := sFile.NewFile(sOrm.NewDb(&auth.Shop.ID), auth.Shop.ID).FileListByIds([]uint{auth.Shop.WebsiteFaviconId})
	if err != nil {
		return res, err
	}
	if files != nil && len(files) > 0 {
		res.WebsiteFavicon = files[0].Path
	}
	return res, err
}

func (a *aShopApi) List(ctx g.Ctx, req *vo.ShopListReq) (res []vo.ShopListRes, err error) {
	user, err := ctx2.NewCtx(ctx).GetUser()
	if err != nil {
		return res, err
	}
	shops, err := sShop.NewShop(sOrm.NewDb(nil)).ShopListByUserId(user.ID)
	if err != nil {
		return res, err
	}
	res = slice.Map(shops, func(_ int, shop mShop.Shop) vo.ShopListRes {
		r := vo.ShopListRes{}
		r.Status = shop.Status
		r.StoreName = shop.StoreName
		r.Uuid = shop.Uuid
		//r.WebsiteFavicon = shop.WebsiteFaviconId
		return r
	})
	return res, err
}

func (a *aShopApi) General(ctx g.Ctx, req *vo.ShopGeneralReq) (res vo.ShopGeneralRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	res.StoreName = shop.StoreName
	res.StoreOwnerEmail = shop.StoreOwnerEmail
	res.CustomerServiceEmail = shop.CustomerServiceEmail
	res.StoreCurrency = shop.StoreCurrency
	res.CurrencyFormatting = shop.CurrencyFormatting
	res.Timezone = shop.TimeZone
	res.PasswordProtection = shop.PasswordProtection
	res.Password = shop.Password
	res.PasswordMessage = shop.PasswordMessage
	res.OrderIdPrefix = shop.OrderIdPrefix
	res.OrderIdSuffix = shop.OrderIdSuffix
	res.WebsiteFaviconId = shop.WebsiteFaviconId
	// 获取地址
	res.Address, err = sAddress.NewAddress(sOrm.NewDb(&auth.Shop.ID), shop.ID).GetAddress(shop.AddressId)
	return res, err
}

func (a *aShopApi) UpdateGeneral(ctx g.Ctx, req *vo.ShopUpdateGeneralReq) (res vo.ShopUpdateGeneralRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		s := sShop.NewShop(tx)
		return s.UpdateShopGeneral(shop.ID, *req, shop.AddressId)
	})
	return res, err
}

func (a *aShopApi) TaxSwitchShipping(ctx g.Ctx, req *vo.ShopTaxSwitchShippingReq) (res vo.ShopTaxSwitchShippingRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	res.TaxShipping = shop.TaxShipping
	return res, err
}

func (s *aShopApi) ShopTaxSwitchShippingUpdate(ctx g.Ctx, req *vo.ShopTaxSwitchShippingUpdateReq) (res vo.ShopTaxSwitchShippingUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sShop.NewShop(tx).ShopUpdateTaxShipping(shop.ID, req.TaxShipping)
	})
	return res, err
}

// 根据shopUuid获取shopId
func (s *aShopApi) ShopId(ctx g.Ctx, req *vo.ShopIdReq) (res vo.ShopIdRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	res.ShopId = shop.ID
	return res, err
}

func (s *aShopApi) TransactionInfo(ctx g.Ctx, req *vo.ShopTransactionInfoReq) (res vo.ShopTransactionInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res, err = sTransaction.NewTransaction(sOrm.NewDb(&shop.ID), shop.ID).Info()
		return err
	})
	return res, err
}

func (s *aShopApi) UpdateTransaction(ctx g.Ctx, req *vo.ShopUpdateTransactionReq) (res vo.ShopUpdateTransactionRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sTransaction.NewTransaction(tx, shop.ID).UpdateTransaction(*req)
	})
	return res, err
}
