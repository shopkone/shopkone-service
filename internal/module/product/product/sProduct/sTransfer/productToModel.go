package sTransfer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/handle"
)

func (s *sProductTransfer) ProductToModel(in vo.BaseProduct, needId bool) (out mProduct.Product) {
	out = mProduct.Product{
		Title:             in.Title,
		Description:       in.Description,
		Status:            in.Status,
		Spu:               in.Spu,
		Vendor:            in.Vendor,
		Tags:              in.Tags,
		VariantType:       in.VariantType,
		InventoryPolicy:   in.InventoryPolicy,
		InventoryTracking: in.InventoryTracking,
		Category:          in.Category,
		ScheduledAt:       handle.ParseTime(in.ScheduledAt),
	}
	if needId && in.Id == 0 {
		return out
	}
	if needId {
		out.ID = in.Id
	}
	out.ShopId = s.shopId
	return out
}
