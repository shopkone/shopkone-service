package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/product/collection/sCollection"
	ctx2 "shopkone-service/utility/ctx"
	"shopkone-service/utility/handle"
)

type aProductCollection struct {
}

func NewProductCollectionApi() *aProductCollection {
	return &aProductCollection{}
}

func (a *aProductCollection) Create(ctx g.Ctx, req *vo.CreateProductCollectionReq) (res vo.CreateProductCollectionRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		res.Id, err = sCollection.NewCollection(tx, shop.ID).Create(*req)
		return err
	})
	return res, err
}

func (a *aProductCollection) Info(ctx g.Ctx, req *vo.ProductCollectionInfoReq) (res vo.ProductCollectionInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sCollection.NewCollection(sOrm.NewDb(), shop.ID).Info(req.Id)
	return res, err
}

func (a *aProductCollection) List(ctx g.Ctx, req *vo.ProductCollectionListReq) (res handle.PageRes[vo.ProductCollectionListRes], err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sCollection.NewCollection(sOrm.NewDb(), shop.ID).List(*req)
	return res, err
}

func (a *aProductCollection) Update(ctx g.Ctx, req *vo.UpdateProductCollectionReq) (res vo.UpdateProductCollectionRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sCollection.NewCollection(tx, shop.ID).Update(*req)
	})
	return res, err
}

func (a *aProductCollection) CollectionOptions(ctx g.Ctx, req *vo.CollectionOptionsReq) (res []vo.CollectionOptionsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	res, err = sCollection.NewCollection(sOrm.NewDb(), shop.ID).CollectionOptions()
	return res, err
}
