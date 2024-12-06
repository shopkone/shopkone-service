package sNav

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/online/mav/mNav"
)

func (s *sNav) NavList() (out []vo.OnlineNavListRes, err error) {
	var navs []mNav.Nav
	if err = s.orm.Model(&mNav.Nav{}).
		Select("id", "title", "handle").Find(&navs).Error; err != nil {
		return out, err
	}
	// 获取第一级的导航菜单
	var items []mNav.NavItem
	if err = s.orm.Model(&mNav.NavItem{}).
		Where("levels = 1").Select("title").
		Find(&items).Error; err != nil {
		return out, err
	}
	for _, nav := range navs {
		out = append(out, vo.OnlineNavListRes{
			ID:     nav.ID,
			Title:  nav.Title,
			Handle: nav.Handle,
			FirstNames: slice.Map(items, func(index int, item mNav.NavItem) string {
				return item.Title
			}),
		})
	}
	return out, err
}
