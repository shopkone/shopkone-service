package sCustomer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

func (s *sCustomer) CustomerUpdateTags(in *vo.CustomerUpdateTagsReq) (err error) {
	data := mCustomer.Customer{
		Tags: in.Tags,
	}
	return s.orm.Model(&data).
		Where("id = ?", in.ID).
		Select("tags").
		Updates(&data).Error
}
