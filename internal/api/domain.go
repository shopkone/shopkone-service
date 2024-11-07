package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/domains/sDomain/sDomain"
	ctx2 "shopkone-service/utility/ctx"
)

type aDomain struct {
}

func NewDomainApi() *aDomain {
	return &aDomain{}
}

func (a *aDomain) List(ctx g.Ctx, req *vo.DomainListReq) (res []vo.DomainListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return nil, err
	}
	shop := auth.Shop
	return sDomain.NewDomain(sOrm.NewDb(), shop.ID).DomainList()
}
