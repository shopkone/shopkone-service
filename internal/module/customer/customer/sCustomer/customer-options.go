package sCustomer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

func (s *sCustomer) CustomerOptions() (res []vo.CustomerOptionsRes, err error) {
	var list []mCustomer.Customer
	if err = s.orm.Model(&list).
		Select("id", "first_name", "last_name", "email", "phone").
		Find(&list).Error; err != nil {
		return res, err
	}
	res = slice.Map(list, func(_ int, item mCustomer.Customer) vo.CustomerOptionsRes {
		return vo.CustomerOptionsRes{
			ID:        item.ID,
			Email:     item.Email,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Phone:     item.Phone,
		}
	})
	return res, err
}
