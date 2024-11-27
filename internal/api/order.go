package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/order/order/sOrder"
	ctx2 "shopkone-service/utility/ctx"
)

type aOrder struct {
}

func NewOrderApi() *aOrder {
	return &aOrder{}
}

func (*aOrder) OrderPreCalPrice(ctx g.Ctx, req *vo.OrderCalPreReq) (res vo.OrderCalPreRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sOrder.NewOrder(sOrm.NewDb(&auth.Shop.ID), shop.ID).OrderPreCalPrice(req)
}
