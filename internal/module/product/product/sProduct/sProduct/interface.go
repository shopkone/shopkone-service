package sProduct

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/handle"

	"gorm.io/gorm"
)

type IProduct interface {
	// 创建商品
	Create(in vo.ProductCreateReq, email string) (res vo.ProductCreateRes, err error)
	// 更新商品
	Update(in vo.ProductUpdateReq, handleEmail string) (err error)
	// 获取商品详情
	Info(id uint) (res vo.ProductInfoRes, err error)
	// 获取商品列表
	List(in vo.ProductListReq) (res handle.PageRes[vo.ProductListRes], err error)
	// 更新商品文件
	UpdateProductFiles(productId uint, fileIds []uint) (err error)
	// 更新标签图片
	UpdateProductOptions(productId uint, images []vo.ProductOption) (err error)
	// ListByIds 批量获取商品信息
	ListByIds(in vo.ListByIdsReq) (res []vo.ListByIdsRes, err error)
	// ListByIdsWithoutVariants 批量获取商品信息(不包含变体)
	ListByIdsWithoutVariants(ids []uint) (res []iProduct.ListByIdsWithoutVariantsOut, err error)
	//  UpdateNameHandler 更新变体名称（用于关联）
	UpdateNameHandler(variants []mProduct.Variant, productId uint) error
}

type sProduct struct {
	orm    *gorm.DB
	shopId uint
}

func NewProduct(orm *gorm.DB, shopId uint) *sProduct {
	return &sProduct{orm, shopId}
}
