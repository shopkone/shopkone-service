package api

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/product/product/sProduct/sProduct"
	"shopkone-service/internal/module/product/product/sProduct/sSupplier"
	ctx2 "shopkone-service/utility/ctx"
	"shopkone-service/utility/handle"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type aProduct struct {
}

func NewProductApi() *aProduct {
	return &aProduct{}
}

func (a *aProduct) Create(ctx g.Ctx, req *vo.ProductCreateReq) (res vo.ProductCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	user := auth.User
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res, err = sProduct.NewProduct(tx, shop.ID).Create(*req, user.Email)
		return err
	})
	return res, err
}

func (a *aProduct) Info(ctx g.Ctx, req *vo.ProductInfoReq) (res vo.ProductInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sProduct.NewProduct(sOrm.NewDb(&auth.Shop.ID), shop.ID).Info(req.Id)
	return res, err
}

func (a *aProduct) List(ctx g.Ctx, req *vo.ProductListReq) (res handle.PageRes[vo.ProductListRes], err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sProduct.NewProduct(sOrm.NewDb(&auth.Shop.ID), shop.ID).List(*req)
	return res, err
}

func (a *aProduct) Update(ctx g.Ctx, req *vo.ProductUpdateReq) (res vo.ProductUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	user := auth.User
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sProduct.NewProduct(tx, shop.ID).Update(*req, user.Email)
	})
	return res, err
}

func (a *aProduct) ListByIds(ctx g.Ctx, req *vo.ListByIdsReq) (res []vo.ListByIdsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sProduct.NewProduct(sOrm.NewDb(&auth.Shop.ID), shop.ID).ListByIds(req.Ids)
}

func (a *aProduct) VariantsByIDs(ctx g.Ctx, req *vo.VariantListByIdsReq) (res []vo.VariantListByIdsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)
	return sProduct.NewProduct(orm, shop.ID).VariantsWithProduct(req)
}

// 创建供应商
func (a *aProduct) CreateSupplier(ctx g.Ctx, req *vo.CreateSupplierReq) (res vo.CreateSupplierRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res.Id, err = sSupplier.NewSupplier(tx, shop.ID).Create(req.Address)
		return err
	})
	return res, err
}

// 获取供应商列表
func (a *aProduct) SupplierList(ctx g.Ctx, req *vo.SupplierListReq) (res []vo.SupplierListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sSupplier.NewSupplier(sOrm.NewDb(&auth.Shop.ID), shop.ID).List()
}

// 更新供应商信息
func (a *aProduct) UpdateSupplier(ctx g.Ctx, req *vo.SupplierUpdateReq) (res vo.SupplierUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		return sSupplier.NewSupplier(tx, shop.ID).Update(*req)
	})
	return res, err
}
