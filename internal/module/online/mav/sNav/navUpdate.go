package sNav

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/online/mav/mNav"
)

func (s *sNav) NavUpdate(req vo.OnlineNavUpdateReq) (res vo.OnlineNavUpdateRes, err error) {
	// 更新导航
	nav := mNav.Nav{
		Handle: req.Handle,
		Title:  req.Title,
		Links:  req.Links,
	}
	if err = s.orm.Model(&mNav.Nav{}).Where("id = ?", req.ID).Updates(&nav).Error; err != nil {
		return res, err
	}
	// 更新导航项
	return res, err
}
