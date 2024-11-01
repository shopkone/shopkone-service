package sTransfer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/handle"
)

func (s *sProductTransfer) ModelToProduct(in mProduct.Product) (out vo.BaseProduct) {
	out = vo.BaseProduct{
		Description:       in.Description,
		Spu:               in.Spu,
		Status:            in.Status,
		Tags:              in.Tags,
		Title:             in.Title,
		VariantType:       in.VariantType,
		Vendor:            in.Vendor,
		InventoryPolicy:   in.InventoryPolicy,
		InventoryTracking: in.InventoryTracking,
		ScheduledAt:       handle.ToUnix(in.ScheduledAt),
		Category:          in.Category,
	}
	out.Id = in.ID
	return out
}
