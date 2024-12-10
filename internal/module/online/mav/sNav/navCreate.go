package sNav

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/online/mav/mNav"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sNav) NavCreate(nav mNav.Nav) (id uint, err error) {
	// 查询名称是否存在
	var count int64
	if err = s.orm.Model(&mNav.Nav{}).Where("handle = ?", nav.Handle).Count(&count).Error; err != nil {
		return id, err
	}
	if count > 0 {
		return id, code.ErrHandleNameRepeated
	}
	// 创建导航
	nav.Links = slice.Map(nav.Links, func(index int, item mNav.NavItem) mNav.NavItem {
		item.ID = handle.GenUid()
		return item
	})
	nav.ShopId = s.shopId
	return nav.ID, s.orm.Create(&nav).Error
}
