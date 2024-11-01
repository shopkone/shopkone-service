package sProduct

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/seo/sSeo"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/product/sProduct/sTransfer"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sProduct) Update(in vo.ProductUpdateReq, handleEmail string) (err error) {
	// 判断variant是否超过500
	if len(in.Variants) > 500 {
		return code.VariantCountMax
	}
	// 获取商品信息
	var info mProduct.Product
	if err = s.orm.Where("id = ? AND shop_id = ?", in.Id, s.shopId).
		First(&info).Error; err != nil {
		return err
	}
	// 更新仓库
	if in.InventoryTracking {
		if err = s.UpdateEnabledLocationIds(in.Id, in.EnabledLocationIds, handleEmail); err != nil {
			return err
		}
	}
	// 更新seo
	in.Seo.ID = info.SeoId
	if err = sSeo.NewSeo(s.orm, s.shopId).Update(in.Seo); err != nil {
		return err
	}
	// 更新商品文件
	if err = s.UpdateProductFiles(in.Id, in.FileIds); err != nil {
		return err
	}
	// 更新商品标签图标
	if err = s.UpdateLabelImages(in.Id, in.LabelImages); err != nil {
		return err
	}
	// 更新商品信息
	data := sTransfer.NewProductTransfer(s.shopId).ProductToModel(in.BaseProduct, true)
	if data.Status == mProduct.VariantStatusPublished && info.Status != mProduct.VariantStatusPublished {
		data.PublishedAt = handle.GetNowTime()
		data.ScheduledAt = nil
	} else {
		data.PublishedAt = nil
	}
	if err = s.orm.Model(&info).
		Select("title",
			"description",
			"status",
			"spu",
			"vendor",
			"tags",
			"variant_type",
			"inventory_policy",
			"inventory_tracking",
			"category",
			"scheduled_at",
		).
		Updates(data).Error; err != nil {
		return err
	}
	// 更新变体
	variantUpdateIn := iProduct.VariantUpdateIn{
		List:              in.Variants,
		ProductId:         data.ID,
		TrackInventory:    data.InventoryTracking,
		HandleEmail:       handleEmail,
		EnableLocationIds: in.EnabledLocationIds,
	}
	variants, err := sVariant.NewVariant(s.orm, s.shopId).Update(variantUpdateIn)
	if err != nil {
		return err
	}
	// 更新变体名称（用于关联系列时使用）
	return s.UpdateNameHandler(variants, in.Id)
}
