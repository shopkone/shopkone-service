package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/customer/customer/sCustomer/sCustomer"
	ctx2 "shopkone-service/utility/ctx"
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
