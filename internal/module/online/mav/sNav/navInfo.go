package sNav

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/online/mav/mNav"
)

func (s *sNav) NavInfo(req vo.OnlineNavInfoReq) (res vo.OnlineNavInfoRes, err error) {
	// 获取导航信息
	var nav mNav.Nav
	if err = s.orm.Model(&mNav.Nav{}).Where("id = ?", req.ID).
		Select("id", "title", "handle", "links").First(&nav).Error; err != nil {
		return res, err
	}
	return vo.OnlineNavInfoRes{
		ID:     nav.ID,
		Title:  nav.Title,
		Handle: nav.Handle,
		Links:  FillLink(nav.Links),
	}, err
}

func FillLink(links []mNav.NavItem) []mNav.NavItem {
	return slice.Map(links, func(index int, item mNav.NavItem) mNav.NavItem {
		if item.Links == nil {
			item.Links = []mNav.NavItem{}
		}
		if len(item.Links) > 0 {
			FillLink(item.Links)
		}
		return item
	})
}
