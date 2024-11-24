package sCustomer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"strings"
)

func (s *sCustomer) Info(id uint) (res vo.CustomerInfoRes, err error) {
	var customer mCustomer.Customer
	if err = s.orm.Model(&customer).Where("id = ?", id).
		Omit("created", "updated", "deleted_at").First(&customer).Error; err != nil {
		return res, err
	}
	res.ID = customer.ID
	res.FirstName = customer.FirstName
	res.LastName = customer.LastName
	res.Email = customer.Email
	res.Note = customer.Note
	res.Gender = customer.Gender
	if customer.Birthday != nil {
		res.Birthday = customer.Birthday.UnixMilli()
	}
	res.Phone = customer.Phone
	res.Tags = customer.Tags

	if customer.Phone != "" {
		phoneArr := strings.Split(customer.Phone, "_____")
		res.Phone = "+" + phoneArr[0] + " " + phoneArr[1]
	}
	// 获取地址ids
	var customerAddressMap []mCustomer.CustomerAddress
	if err = s.orm.Model(&customerAddressMap).Where("customer_id = ?", customer.ID).
		Omit("created", "updated", "deleted_at").Find(&customerAddressMap).Error; err != nil {
		return res, err
	}
	addressIds := slice.Map(customerAddressMap, func(index int, item mCustomer.CustomerAddress) uint {
		return item.AddressID
	})
	// 获取地址
	address, err := sAddress.NewAddress(s.orm, s.shopId).ListByIds(addressIds)
	if err != nil {
		return vo.CustomerInfoRes{}, err
	}
	res.Address = address
	return res, err
}
