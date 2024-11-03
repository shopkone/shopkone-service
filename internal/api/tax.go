package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/product/collection/sCollection"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/internal/module/setting/tax/sTax/sTax"
	ctx2 "shopkone-service/utility/ctx"
)

type aTax struct {
}

func NewTaxApi() *aTax {
	return &aTax{}
}

func (a *aTax) List(ctx g.Ctx, req *vo.TaxListReq) (res []vo.TaxListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return nil, err
	}
	shop := auth.Shop
	return sTax.NewTax(sOrm.NewDb(), shop.ID).TaxList()
}

func (a *aTax) Info(ctx g.Ctx, req *vo.TaxInfoReq) (res vo.TaxInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	shop := auth.Shop
	return sTax.NewTax(sOrm.NewDb(), shop.ID).TaxInfo(req.Id)
}

func (a *aTax) TaxUpdate(ctx g.Ctx, req *vo.TaxUpdateReq) (res vo.TaxUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	shop := auth.Shop
	orm := sOrm.NewDb()
	// 校验collectionIds是否都存在
	collectionIds := slice.Map(req.Customers, func(_ int, item vo.BaseCustomerTax) uint {
		if item.Type == mTax.CustomerTaxTypeCollection {
			return item.CollectionID
		}
		return 0
	})
	if err = sCollection.NewCollection(orm, shop.ID).CheckAllExist(collectionIds); err != nil {
		return res, err
	}
	// 更新
	return res, sTax.NewTax(orm, shop.ID).TaxUpdate(*req)
}
