package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sVariant) VariantNameToString(names []mProduct.VariantName) string {
	values := slice.Map(names, func(index int, item mProduct.VariantName) string {
		return item.Value
	})
	values = slice.Filter(values, func(index int, item string) bool {
		return item != ""
	})
	return slice.Join(values, " · ")
}
