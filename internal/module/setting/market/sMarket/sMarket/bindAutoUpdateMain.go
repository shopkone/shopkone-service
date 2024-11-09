package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

// 自动同步主域名且不是主市场
func (s *sMarket) BindAutoUpdateMain() (err error) {
	// 获取语言绑定列表
	languages, err := s.BindListALl()
	if err != nil {
		return err
	}

	// 获取市场列表
	markets, err := s.MarketList()
	if err != nil {
		return err
	}

	// 获取不是主市场，但是用了主域名的市场
	mainDomainMarkets := slice.Filter(markets, func(index int, item vo.MarketListRes) bool {
		return !item.IsMain && item.DomainType == mMarket.DomainTypeMain
	})
	if len(mainDomainMarkets) == 0 {
		return nil
	}

	// 对比与主市场的差异并进行增删改查

	return nil
}
