package sCustomer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
	"strings"
)

func (s *sCustomer) CustomerUpdateBaseInfo(in *vo.CustomerUpdateBaseReq) (err error) {
	// 校验必须至少存在一项
	if in.FirstName == "" && in.LastName == "" && in.Phone.Num == "" && in.Email == "" {
		return code.ErrCusterCreateErr
	}
	// 校验手机号码或电话是否重复
	checkIn := CheckPhoneOrEmailIn{
		Phone:      in.Phone,
		Email:      in.Email,
		CustomerID: in.ID,
	}
	if err = s.CheckPhoneOrEmail(checkIn); err != nil {
		return err
	}
	// 开始更新
	data := mCustomer.Customer{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Gender:    in.Gender,
		Language:  in.Language,
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
	return s.orm.Model(&mCustomer.Customer{}).
		Where("id = ?", in.ID).
		Select("first_name", "last_name", "phone", "email", "gender", "birthday", "language").
		Updates(&data).Error
}
