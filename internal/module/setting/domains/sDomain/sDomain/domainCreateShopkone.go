package sDomain

import "github.com/gogf/gf/v2/frame/g"

func (s *sDomain) CreateShopkone() error {
	// 生成可用的域名
	domain, err := s.GeneralShopkoneDomain()
	if err != nil {
		return err
	}
	// 预检
	data, err := s.PreCheck(domain)
	if err != nil {
		return err
	}

	// TODO:绑定域名
	g.Dump(data)

	// 校验域名是否已经绑定
	checkIn := ConnectCheckIn{
		Domain:     domain,
		IsShopkone: true,
	}
	return s.ConnectCheck(checkIn)
}
