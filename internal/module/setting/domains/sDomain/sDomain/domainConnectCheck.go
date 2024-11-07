package sDomain

import "shopkone-service/internal/module/setting/domains/mDomains"

func (s *sDomain) ConnectCheck(domain string, ip string, isShopkone bool) (err error) {
	// TODO: 校验是否绑定
	// 如果没有绑定，直接返回错误信息，如果绑定了，则创建记录
	data := mDomains.Domain{}
	data.ShopId = s.shopId
	data.Domain = domain
	data.Status = mDomains.DomainStatusConnectSuccess
	data.IsShopKone = isShopkone
	data.Ip = ip
	data.IsMain = isShopkone
	return s.orm.Create(&data).Error
}
