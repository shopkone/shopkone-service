package sProduct

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/seo/sSeo"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/product/sProduct/sTransfer"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sProduct) Create(in vo.ProductCreateReq, email string) (res vo.ProductCreateRes, err error) {
	// 判断variant是否超过500
	if len(in.Variants) > 500 {
		return res, code.VariantCountMax
	}
	// 创建商品seo
	seoId, err := sSeo.NewSeo(s.orm, s.shopId).Create(in.Seo)
	if err != nil {
		return res, err
	}
	// 创建商品
	data := sTransfer.NewProductTransfer(s.shopId).ProductToModel(in.BaseProduct, false)
	data.SeoId = seoId
	if data.Status == mProduct.VariantStatusPublished {
		data.PublishedAt = handle.GetNowTime()
		data.ScheduledAt = nil
	}
	err = s.orm.Create(&data).Error
	if err != nil {
		return vo.ProductCreateRes{}, err
	}
	res.Id = data.ID
	// 创建变体
	createIn := iProduct.VariantCreateIn{
		List:              in.Variants,
		ProductId:         data.ID,
		InventoryTracking: in.InventoryTracking,
		Type:              mInventory.InventoryAddProduct,
		HandleEmail:       email,
		EnableLocationIds: in.EnabledLocationIds,
	}
	variants, err := sVariant.NewVariant(s.orm, s.shopId).Create(createIn)
	if err != nil {
		return res, err
	}
	// 创建商品文件
	if len(in.FileIds) > 0 {
		if err = s.UpdateProductFiles(data.ID, in.FileIds); err != nil {
			return vo.ProductCreateRes{}, err
		}
	}
	// 创建标签图片
	if len(in.ProductOptions) > 0 {
		if err = s.UpdateProductOptions(data.ID, in.ProductOptions); err != nil {
			return vo.ProductCreateRes{}, err
		}
	}
	// 关联系列
	if err = s.RelativeCollection(in.Collections, data.ID); err != nil {
		return vo.ProductCreateRes{}, err
	}
	// 更新变体名称（用于关联系列时使用）
	return res, s.UpdateNameHandler(variants, data.ID)
}
