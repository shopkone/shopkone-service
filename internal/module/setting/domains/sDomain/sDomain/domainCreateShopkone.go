package sDomain

import "github.com/gogf/gf/v2/frame/g"

func (s *sDomain) CreateShopkone() error {
	// 生成可用的域名
	domain, err := s.GeneralShopkoneDomain()
	if err != nil {
		return err
	}
	// 预检
	data, err := s.PreCheck(domain, true)
	if err != nil {
		return err
	}

	// 绑定域名
	g.Dump(data)

	// 校验域名是否已经绑定
	return s.ConnectCheck(domain)
}
