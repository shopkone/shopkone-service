package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
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
	return res, sTax.NewTax(sOrm.NewDb(), shop.ID).TaxUpdate(*req)
}
