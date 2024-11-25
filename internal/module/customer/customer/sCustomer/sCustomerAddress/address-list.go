package sCustomerAddress

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

type CustomerAddressListOut struct {
	CustomerID       uint
	DefaultAddressID uint
	Address          []mAddress.Address
}

func (s *sCustomerAddress) List(customerIds []uint) (list []CustomerAddressListOut, err error) {
	// 获取客户地址关系
	var customerAddressMap []mCustomer.CustomerAddress
	if err = s.orm.Model(&customerAddressMap).Where("customer_id in ?", customerIds).
		Select("customer_id", "address_id", "is_default").
		Find(&customerAddressMap).Error; err != nil {
		return list, err
	}
	// 晒出地址id
	addressIds := slice.Map(customerAddressMap, func(index int, item mCustomer.CustomerAddress) uint {
		return item.AddressID
	})
	// 获取地址
	address, err := sAddress.NewAddress(s.orm, s.shopID).ListByIds(addressIds)
	if err != nil {
		return nil, err
	}
	list = slice.Map(customerIds, func(index int, id uint) CustomerAddressListOut {
		i := CustomerAddressListOut{}
		i.CustomerID = id
		currentCustomerAddressMap := slice.Filter(customerAddressMap, func(index int, item mCustomer.CustomerAddress) bool {
			return item.CustomerID == id
		})
		currentAddress := slice.Filter(address, func(index int, addr mAddress.Address) bool {
			_, ok := slice.FindBy(currentCustomerAddressMap, func(index int, item mCustomer.CustomerAddress) bool {
				if item.IsDefault {
					i.DefaultAddressID = addr.ID
				}
				return item.AddressID == addr.ID
			})
			return ok
		})
		i.Address = currentAddress
		return i
	})
	return list, err
}
