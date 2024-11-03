package sCustomerTax

import "shopkone-service/internal/module/setting/tax/mTax"

func (s *sCustomerTax) CustomerTaxIsChange(newCustomer, oldCustomer mTax.CustomerTax) bool {
	if newCustomer.CollectionID != oldCustomer.CollectionID {
		return true
	}
	if newCustomer.Type != oldCustomer.Type {
		return true
	}
	return false
}
