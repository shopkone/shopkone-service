package sDomain

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/domains/mDomains"
)

func (s *sDomain) DomainList(req *vo.DomainListReq) (res []vo.DomainListRes, err error) {
	list := []mDomains.Domain{}
	query := s.orm.Model(&list).Where("shop_id = ?", s.shopId)
	query = query.Omit("shop_id", "created_at", "updated_at", "deleted_at")
	if len(req.Status) > 0 {
		query = query.Where("status IN ?", req.Status)
	}
	if err = query.Find(&list).Error; err != nil {
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
	// TODO:如果主域名不是shopkone域名，且主域名不可用，立即设置为shopkone域名（需要做定时任务）
	return res, err
}
