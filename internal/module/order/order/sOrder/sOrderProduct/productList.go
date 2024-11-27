package sOrderProduct

import (
	"shopkone-service/internal/module/product/product/sProduct/sProduct"
)

func (s *sOrderProduct) GetProductList(variantIds []uint) ([]sProduct.ProductsToOrderOut, error) {
	return sProduct.NewProduct(s.orm, s.shopId).ListToOrder(variantIds)
}
