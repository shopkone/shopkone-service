package sMarketLanguage

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarketLanguage) BindListByLangIds(langIds []uint) (out []mMarket.MarketLanguage, err error) {
	if err = s.orm.Model(&out).Where("language_id in (?)", langIds).
		Where("shop_id = ?", s.shopId).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").
		Find(&out).Error; err != nil {
		return out, err
	}
	return out, err
}
