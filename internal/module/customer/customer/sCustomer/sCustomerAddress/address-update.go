package sCustomerAddress

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/code"
)

func (s *sCustomerAddress) Update(in *vo.CustomerUpdateAddressReq) (err error) {
	if in.Address.ID == 0 {
		return code.IdMissing
	}
	var data mCustomer.CustomerAddress
	// 如果要设置 default 为 true，则先将之前的设为false
	if in.IsDefault {
		if err = s.orm.Model(&data).Where("is_default = ?", true).
			Where("customer_id = ?", in.CustomerID).
			Update("is_default", false).Error; err != nil {
			return err
		}
		if err = s.orm.Model(&data).Where("address_id = ?", in.Address.ID).
			Update("is_default", true).Error; err != nil {
			return err
		}
	}
	// 更新地址信息
	return sAddress.NewAddress(s.orm, s.shopID).UpdateById(in.Address)
}
