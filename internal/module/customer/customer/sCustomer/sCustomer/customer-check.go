package sCustomer

import (
	"github.com/duke-git/lancet/v2/convertor"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/code"
)

type CheckPhoneOrEmailIn struct {
	CustomerID uint
	Phone      mAddress.Phone
	Email      string
}

func (s *sCustomer) CheckPhoneOrEmail(in CheckPhoneOrEmailIn) (err error) {
	if in.Email == "" && in.Phone.Num == "" {
		return err
	}
	var customer mCustomer.Customer
	query := s.orm.Model(&mCustomer.Customer{})
	if in.CustomerID > 0 {
		query = query.Where("id != ?", in.CustomerID)
	}
	if in.Phone.Num != "" {
		phone := convertor.ToString(in.Phone.Prefix) + "_____" + in.Phone.Num
		query = query.Where("phone = ?", phone)
	}
	if in.Email != "" && in.Phone.Num == "" {
		query = query.Where("email = ?", in.Email)
	}
	if in.Email != "" && in.Phone.Num != "" {
		query = query.Or("email = ?", in.Email)
	}
	query = query.Select("id", "email", "phone")
	if err = query.Find(&customer).Error; err != nil {
		return err
	}
	if customer.Email != "" {
		return code.ErrCustomerEmailExist
	}
	if customer.Phone != "" {
		return code.ErrCustomerPhoneExist
	}
	return nil
}
