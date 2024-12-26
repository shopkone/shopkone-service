package sOrderTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/collection/mCollection"
	"shopkone-service/internal/module/product/collection/sCollection"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
	"shopkone-service/internal/module/setting/tax/sTax/sCustomerTax"
	"shopkone-service/internal/module/setting/tax/sTax/sTax"
)

type TaxCalIn struct {
	CountryCode string
	ZoneCode    string
	Variants    []sVariant.VariantToOrderOut
	InVariants  []vo.OrderPreBaseVariantItem
}

type TaxCalOut struct {
	VariantID uint
	TaxRate   float64
	TaxName   string
	Tax       uint32
}

func (s *sOrderTax) TaxCal(in TaxCalIn) (out []TaxCalOut, err error) {
	if in.CountryCode == "" {
		return out, err
	}
	productIds := slice.Map(in.Variants, func(index int, item sVariant.VariantToOrderOut) uint {
		return item.ProductID
	})
	productIds = slice.Unique(productIds)

	// 获取集合
	cp, err := sCollection.NewCollection(s.orm, s.shopId).CollectionsByProductIds(productIds)
	if err != nil {
		return out, err
	}
	collectionIDs := slice.Map(cp, func(index int, item mCollection.CollectionProduct) uint {
		return item.CollectionId
	})

	// 获取税率
	TaxByCountryProductIn := sTax.TaxByCountryProductIn{
		CountryCode:   in.CountryCode,
		ZoneCode:      in.ZoneCode,
		CollectionIDs: collectionIDs,
	}
	taxes, err := sTax.NewTax(s.orm, s.shopId).TaxByCountryProduct(TaxByCountryProductIn)
	if err != nil {
		return out, err
	}

	// 计算税费
	out = slice.Map(in.Variants, func(index int, variant sVariant.VariantToOrderOut) TaxCalOut {
		i := TaxCalOut{}
		i.VariantID = variant.ID
		// 找出这个变体所属的集合
		collection, ok := slice.FindBy(cp, func(index int, item mCollection.CollectionProduct) bool {
			return item.ProductId == variant.ProductID
		})
		// 找出数量
		inVariant, ok := slice.FindBy(in.InVariants, func(index int, inVariant vo.OrderPreBaseVariantItem) bool {
			return inVariant.VariantID == variant.ID
		})
		if !ok {
			return TaxCalOut{}
		}
		quantity := inVariant.Quantity
		// 找出价格
		price := variant.Price * quantity
		// 先找到符合的自定义税费
		customerTax, ok := slice.FindBy(taxes.Customer, func(index int, item sCustomerTax.CustomerTaxByCountryProductOut) bool {
			return item.CollectionID == collection.CollectionId
		})
		// 如果找到了，直接用自定义税费
		if ok {
			i.TaxRate = customerTax.TaxRate
			i.TaxName = customerTax.TaxName
			i.Tax = uint32(float64(price) * i.TaxRate / 100)
			return i
		}
		// 找到区域，也直接使用
		if taxes.Zone.Name != "" {
			i.TaxRate = taxes.Zone.Rate
			i.TaxName = taxes.Zone.Name
			i.Tax = uint32(float64(price) * i.TaxRate / 100)
			return i
		}
		// 使用国家税率
		i.TaxRate = taxes.Base.Rate
		i.TaxName = taxes.Base.Name
		i.Tax = uint32(float64(price) * i.TaxRate / 100)
		return i
	})
	out = slice.Filter(out, func(index int, item TaxCalOut) bool {
		return item.Tax != 0
	})
	return out, err
}
