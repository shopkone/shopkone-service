package sCustomer

import (
	"github.com/duke-git/lancet/v2/convertor"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
	"strings"
)

func (s *sCustomer) Create(in *vo.CustomerCreateReq) (id uint, err error) {
	// 校验必须至少存在一项
	if in.FirstName == "" && in.LastName == "" && in.Phone.Num == "" && in.Email == "" {
		return 0, code.ErrCusterCreateErr
	}
	// 校验手机号码或电话是否重复
	checkIn := CheckPhoneOrEmailIn{
		Phone: in.Phone,
		Email: in.Email,
	}
	if err = s.CheckPhoneOrEmail(checkIn); err != nil {
		return 0, err
	}
	data := mCustomer.Customer{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Phone:     in.Phone.Num,
		Email:     in.Email,
		Note:      in.Note,
		Gender:    in.Gender,
		Tags:      in.Tags,
	}
	if in.Birthday != 0 {
		data.Birthday = handle.ParseTime(in.Birthday)
	}
	if in.Phone.Num != "" {
		data.Phone = convertor.ToString(in.Phone.Prefix) + "_____" + in.Phone.Num
	}
	if in.FirstName == "" && in.LastName == "" {
		if in.Email != "" {
			emailArr := strings.Split(in.Email, "@")
			data.LastName = emailArr[0]
		} else if in.Phone.Num != "" {
			data.LastName = in.Phone.Num
		}
	}
	// 创建客户
	if err = s.orm.Create(&data).Error; err != nil {
		return 0, err
	}
	// 如果不存在国家，说明没没有设置收货地址
	if in.Address.Country == "" {
		return data.ID, err
	}
	// 创建收货地址
	addressId, err := sAddress.NewAddress(s.orm, s.shopId).CreateAddress(in.Address)
	if err != nil {
		return 0, err
	}
	// 创建收货地址关联
	customerAddres := mCustomer.CustomerAddress{
		CustomerID: data.ID,
		AddressID:  addressId,
	}
	return data.ID, s.orm.Create(&customerAddres).Error
}
