package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/customer/customer/sCustomer"
	ctx2 "shopkone-service/utility/ctx"
	"shopkone-service/utility/handle"
)

type aCustomer struct {
}

func NewCustomerApi() *aCustomer {
	return &aCustomer{}
}

func (*aCustomer) Create(ctx g.Ctx, req *vo.CustomerCreateReq) (res vo.CustomerCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res.ID, err = sCustomer.NewCustomer(tx, shop.ID).Create(req)
		return err
	})
	return res, err
}

func (*aCustomer) Info(ctx g.Ctx, req *vo.CustomerInfoReq) (res vo.CustomerInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)
	res, err = sCustomer.NewCustomer(orm, shop.ID).Info(req.ID)
	return res, err
}

func (*aCustomer) List(ctx g.Ctx, req *vo.CustomerListReq) (res handle.PageRes[vo.CustomerListRes], err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)
	res, err = sCustomer.NewCustomer(orm, shop.ID).List(*req)
	return res, err
}

// 更新客户信息
func (*aCustomer) UpdateBase(ctx g.Ctx, req *vo.CustomerUpdateBaseReq) (res vo.CustomerUpdateBaseRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)
	err = orm.Transaction(func(tx *gorm.DB) error {
		return sCustomer.NewCustomer(tx, shop.ID).CustomerUpdateBaseInfo(req)
	})
	return res, err
}

// 更新客户标签
func (*aCustomer) UpdateTags(ctx g.Ctx, req *vo.CustomerUpdateTagsReq) (res vo.CustomerUpdateTagsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)
	err = orm.Transaction(func(tx *gorm.DB) error {
		return sCustomer.NewCustomer(tx, shop.ID).CustomerUpdateTags(req)
	})
	return res, err
}

// 更新客户备注
func (*aCustomer) UpdateNote(ctx g.Ctx, req *vo.CustomerUpdateNoteReq) (res vo.CustomerUpdateNoteRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)
	err = orm.Transaction(func(tx *gorm.DB) error {
		return sCustomer.NewCustomer(tx, shop.ID).CustomerUpdateNote(req)
	})
	return res, err
}

// 更新客户免税地区
func (*aCustomer) SetTax(ctx g.Ctx, req *vo.CustomerSetTaxReq) (res vo.CustomerSetTaxRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)
	err = orm.Transaction(func(tx *gorm.DB) error {
		return sCustomer.NewCustomer(tx, shop.ID).CustomerUpdateTax(req)
	})
	return res, err
}
