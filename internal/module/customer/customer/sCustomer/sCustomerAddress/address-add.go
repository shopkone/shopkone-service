package sCustomerAddress

import (
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

func (s *sCustomerAddress) Add(customerId uint, address mAddress.Address) (err error) {
	// 查询用户中是否存在默认地址
	var defaultCount int64
	if err = s.orm.Model(&mCustomer.CustomerAddress{}).
		Where("shop_id = ? AND is_default = ?", s.shopID, true).
		Count(&defaultCount).Error; err != nil {
		return err
	}
	// 开始创建地址
	addressId, err := sAddress.NewAddress(s.orm, s.shopID).CreateAddress(address)
	if err != nil {
		return err
	}
	// 建立关联
	customerAddress := mCustomer.CustomerAddress{
		CustomerID: customerId,
		AddressID:  addressId,
		IsDefault:  defaultCount == 0,
	}
	return s.orm.Select("is_default", "customer_id", "address_id", "shop_id").Create(&customerAddress).Error
}
