package sDomain

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/domains/mDomains"
)

func (s *sDomain) DomainList() (res []vo.DomainListRes, err error) {
	list := []mDomains.Domain{}
	if err = s.orm.Model(&list).Where("shop_id = ?", s.shopId).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").
		Find(&list).Error; err != nil {
		return res, err
	}
	res = slice.Map(list, func(index int, item mDomains.Domain) vo.DomainListRes {
		i := vo.DomainListRes{}
		i.ID = item.ID
		i.IsMain = item.IsMain
		i.IsShopKone = item.IsShopKone
		i.Status = item.Status
		i.Domain = item.Domain
		return i
	})

	// TODO:挨个校验是否可以连接
	return res, err
}
