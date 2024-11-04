package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
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
	s := sTax.NewTax(orm, shop.ID)
	if err = s.TaxUpdate(*req); err != nil {
		return vo.TaxUpdateRes{}, err
	}
	// 校验错误
	res.TaxInfoRes, err = s.CheckError(req.ID)
	if err != nil {
		return vo.TaxUpdateRes{}, err
	}
	return res, err
}

func (a *aTax) TaxCreate(ctx g.Ctx, req *vo.TaxCreateReq) (res vo.TaxCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sTax.NewTax(tx, shop.ID).TaxCreate(req.CountryCodes)
	})
	return res, err
}

func (a *aTax) TaxRemove(ctx g.Ctx, req *vo.TaxRemoveReq) (res vo.TaxRemoveRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		s := sTax.NewTax(tx, shop.ID)
		return s.TaxRemoveByIds(req.Ids)
	})
	return res, err
}

func (a *aTax) TaxActive(ctx g.Ctx, req *vo.TaxActiveReq) (res vo.TaxActiveRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		s := sTax.NewTax(tx, shop.ID)
		if err = s.TaxActive(*req); err != nil {
			return err
		}
		res.List, err = s.TaxList()
		return err
	})
	return res, err
}
