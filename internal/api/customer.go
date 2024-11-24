package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/customer/customer/sCustomer/sCustomer"
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
