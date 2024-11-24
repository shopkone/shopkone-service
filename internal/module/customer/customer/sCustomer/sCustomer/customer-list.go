package sCustomer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/handle"
	"strings"
)

func (s *sCustomer) List(in vo.CustomerListReq) (res handle.PageRes[vo.CustomerListRes], err error) {
	var list []mCustomer.Customer
	query := s.orm.Model(&list).Where("shop_id = ?", s.shopId)
	// 获取count
	if err = query.Count(&res.Total).Error; err != nil {
		return res, err
	}
	// 分页
	query = query.Scopes(handle.Pagination(in.PageReq)).Order("id desc")
	query = query.Select("id", "first_name", "last_name", "email", "phone", "birthday", "gender", "note")
	// 查询数据
	if err = query.Find(&list).Error; err != nil {
		return res, err
	}
	// 组装数据
	res.List = slice.Map(list, func(_ int, item mCustomer.Customer) vo.CustomerListRes {
		result := vo.CustomerListRes{
			FirstName:      item.FirstName,
			LastName:       item.LastName,
			ID:             item.ID,
			OrderCount:     0,
			CostPrice:      0,
			EmailSubscribe: false,
			Email:          item.Email,
			Phone:          item.Phone,
		}
		if result.Phone != "" {
			phoneArr := strings.Split(result.Phone, "_____")
			result.Phone = "+" + phoneArr[0] + " " + phoneArr[1]
		}
		return result
	})
	return res, err
}
