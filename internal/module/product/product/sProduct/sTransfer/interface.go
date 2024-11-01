package sTransfer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
)

type IProductTransfer interface {
	// ProductToModel 转模型
	ProductToModel(product vo.BaseProduct) mProduct.Product
	// ModelToProduct 转vo
	ModelToProduct(product mProduct.Product) vo.BaseProduct
	// VariantToModel 转模型
	VariantToModel(variant vo.BaseVariant) mProduct.Variant
	// ModelToVariant 转vo
	ModelToVariant(variant mProduct.Variant) vo.BaseVariant
}

type sProductTransfer struct {
	shopId uint
}

func NewProductTransfer(shopId uint) *sProductTransfer {
	return &sProductTransfer{shopId: shopId}
}
