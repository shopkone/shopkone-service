package sTax

type TaxByCountryProductIn struct {
	CountryCode string
	ZoneCode    string
	ProductIDs  []uint
}

type TaxByCountryProductOut struct {
	TaxName         string
	TaxRate         float64
	TaxShippingName string
	TaxShippingRate float64
}

func (s *sTax) TaxByCountryProduct(in TaxByCountryProductIn) (out TaxByCountryProductOut, err error) {
	// 根据商品获取专辑
	/*	cp, err := sCollection.NewCollection(s.orm, s.shopId).CollectionsByProductIds(in.ProductIDs)

		// 获取国家税率
		var tax mTax.Tax
		if err = s.orm.Where("country_code = ?", in.CountryCode).
			Omit("created_at", "updated_at", "deleted_at").
			Find(&tax).Error; err != nil {
			return out, err
		}

		// 如果没有启用，直接返回
		if tax.Status == mTax.TaxStatusInactive {
			return out, nil
		}*/

	// 查找自定义运费税率

	// 查找同运费
	return out, err
}
