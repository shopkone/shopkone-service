package sCustomer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/internal/module/customer/customer/sCustomer/sCustomerAddress"
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
		data.Phone = GenPhoneStr(in.Phone)
	}
	// 设置默认姓名
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
	err = sCustomerAddress.NewCustomerAddress(s.shopId, s.orm).Add(data.ID, in.Address)
	return data.ID, err
}
