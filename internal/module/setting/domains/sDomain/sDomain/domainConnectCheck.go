package sDomain

import "shopkone-service/internal/module/setting/domains/mDomains"

type ConnectCheckIn struct {
	Domain     string
	IsShopkone bool
}

func (s *sDomain) ConnectCheck(in ConnectCheckIn) (err error) {
	// 再进行一次校验
	v, err := s.PreCheck(in.Domain)
	if err != nil {
		return err
	}

	// TODO: 校验是否绑定
	// 如果没有绑定，直接返回错误信息，如果绑定了，则创建记录
	data := mDomains.Domain{}
	data.ShopId = s.shopId
	data.Domain = in.Domain
	data.Status = mDomains.DomainStatusConnectSuccess
	data.IsShopKone = in.IsShopkone
	data.BindIp = v.BindIp
	data.BindDomain = v.BindDomain
	data.IsMain = in.IsShopkone
	return s.orm.Create(&data).Error
}
