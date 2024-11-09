package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) UnbindByMarketIds(ids []uint) (err error) {

	if err = s.orm.Model(&mMarket.MarketLanguage{}).
		Where("shop_id = ? AND market_id IN ?", s.shopId, ids).
		Delete(&mMarket.MarketLanguage{}).Error; err != nil {
		return err
	}

	// 同步主市场语言配置至其他市场主域名
	return s.BindAutoUpdateMain()
}
