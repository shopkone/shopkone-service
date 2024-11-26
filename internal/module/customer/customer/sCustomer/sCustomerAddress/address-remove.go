package sCustomerAddress

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

func (s *sCustomerAddress) Remove(addressId uint) (err error) {
	// 获取客户地址关联
	var data mCustomer.CustomerAddress
	if err = s.orm.Model(&data).Where("address_id = ?", addressId).First(&data).Error; err != nil {
		return err
	}
	// 删除客户地址关联
	if err = s.orm.Delete(&mCustomer.CustomerAddress{}, "address_id = ?", addressId).Error; err != nil {
		return err
	}
	// 删除地址
	if err = sAddress.NewAddress(s.orm, s.shopID).
		AddressRemoveIds([]uint{addressId}); err != nil {
		return err
	}
	// 如果没有默认地址，则设置第一个为默认地址
	var list []mCustomer.CustomerAddress
	if err = s.orm.Model(&list).Where("customer_id = ?", data.CustomerID).
		Select("id", "customer_id", "address_id", "is_default").Find(&list).Error; err != nil {
		return err
	}
	if len(list) == 0 {
		return err
	}
	hasDefault := slice.Some(list, func(index int, item mCustomer.CustomerAddress) bool {
		return item.IsDefault
	})
	if !hasDefault {
		return s.orm.Model(&list[0]).Update("is_default", true).Error
	}
	return err
}
