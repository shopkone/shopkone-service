package sCustomerTax

/*type customerTaxByCountryProductIn struct {
	TaxID         uint
	CountryCode   string
	ZoneCode      string
	CollectionIds []uint
}

func (s *sCustomerTax) customerTaxByCountryProduct(in customerTaxByCountryProductIn) {
	// 获取区域自定义税率
	var customerTaxArea mTax.CustomerTax
	if err := s.orm.Model(&customerTaxArea).
		Where("shop_id = ? AND tax_id = ?", s.shopId, in.TaxID).
		Where("collection_id IN ?", in.CollectionIds).
		Where("type = ?", mTax.CustomerTaxTypeCollection).
		First(&customerTaxArea).Error; err != nil {
		return
	}
	// 获取区域自定义税率详细内容

	// 获取运费自定义税率
	var shippingTaxArea mTax.CustomerTax
	if err := s.orm.Model(&shippingTaxArea).
		Where("shop_id = ? AND tax_id = ?", s.shopId, in.TaxID).
		Where("type = ?", mTax.CustomerTaxTypeDelivery).
		First(&shippingTaxArea).Error; err != nil {
		return
	}
}
*/
