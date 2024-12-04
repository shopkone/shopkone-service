package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/design/sDesign/sDesign"
	ctx2 "shopkone-service/utility/ctx"
)

type aDesign struct {
}

func NewDesignApi() *aDesign {
	return &aDesign{}
}

func (a *aDesign) DesignDataList(ctx g.Ctx, req *vo.DesignDataListReq) (res *vo.DesignDataListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	return sDesign.NewDesign(shop.ID).ListData(ctx)
}

func (a *aDesign) SchemaList(ctx g.Ctx, req *vo.DesignSchemaListReq) (res []vo.DesignSchemaListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	return sDesign.NewDesign(shop.ID).ListSchema(ctx, req)
}

func (a *aDesign) BlockUpdate(ctx g.Ctx, req *vo.DesignUpdateBlockReq) (res vo.DesignUpdateBlockRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	return sDesign.NewDesign(shop.ID).BlockUpdate(ctx, req)
}
