package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) BindListALl() (list []mMarket.MarketLanguage, err error) {
	return list, s.orm.Model(&list).Where("shop_id = ?", s.shopId).
		Omit("created_at", "updated_at", "deleted_at", "shop_id").Find(&list).Error
}
