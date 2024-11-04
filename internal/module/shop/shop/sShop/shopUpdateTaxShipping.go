package sShop

import "shopkone-service/internal/module/shop/shop/mShop"

func (s *sShop) ShopUpdateTaxShipping(id uint, active bool) (err error) {
	if err = s.orm.Model(&mShop.Shop{}).Where("id = ?", id).
		Update("tax_shipping", active).Error; err != nil {
		return err
	}
	return NewShopCache().UpdateShopCache(id, s.orm)
}
