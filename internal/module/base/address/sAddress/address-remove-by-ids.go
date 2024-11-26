package sAddress

import "shopkone-service/internal/module/base/address/mAddress"

func (s *sAddress) AddressRemoveIds(addressIds []uint) (err error) {
	return s.orm.Model(&mAddress.Address{}).
		Where("shop_id = ? AND id in (?)", s.shopId, addressIds).
		Delete(&mAddress.Address{}).Error
}
