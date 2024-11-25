package sCustomer

import (
	"github.com/duke-git/lancet/v2/convertor"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/code"
	"strings"
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
	phone := GenPhoneStr(in.Phone)
	if in.Phone.Num != "" && in.Email != "" {
		query = query.Where("(phone = ? or email = ?)", phone, in.Email)
	} else if in.Phone.Num != "" && in.Email == "" {
		query = query.Where("phone = ?", phone)
	} else if in.Email != "" && in.Phone.Num == "" {
		query = query.Where("email = ?", in.Email)
	}
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

func GenPhoneStr(phone mAddress.Phone) string {
	return "(+" + convertor.ToString(phone.Prefix) + ") " + phone.Num
}

func ParsePhoneStr(phoneStr string) (phone mAddress.Phone, err error) {
	phoneArr := strings.Split(phoneStr, ") ")
	if len(phoneArr) > 1 {
		// 去掉(+
		prefixStr := strings.Replace(phoneArr[0], "(+", "", 1)
		prefixInt, err := convertor.ToInt(prefixStr)
		if err == nil {
			phone.Prefix = int(prefixInt)
			phone.Num = phoneArr[1]
		} else {
			return phone, err
		}
	}
	return phone, err
}
