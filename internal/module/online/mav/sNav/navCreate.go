package sNav

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/online/mav/mNav"
	"shopkone-service/utility/code"
)

func (s *sNav) NavCreate(nav mNav.Nav, items []mNav.NavItem) (err error) {
	// 查询名称是否存在
	var count int64
	if err = s.orm.Model(&mNav.Nav{}).Where("handle = ?", nav.Handle).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return code.ErrHandleNameRepeated
	}
	// 创建导航
	nav.ShopId = s.shopId
	if err = s.orm.Create(&nav).Error; err != nil {
		return err
	}
	// 创建导航项
	items = slice.Map(items, func(index int, item mNav.NavItem) mNav.NavItem {
		item.ShopId = s.shopId
		item.NavId = nav.ID
		return item
	})
	if err = s.orm.Create(&items).Error; err != nil {
		return err
	}
	return err
}
