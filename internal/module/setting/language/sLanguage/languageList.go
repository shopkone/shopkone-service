package sLanguage

import "shopkone-service/internal/module/setting/language/mLanguage"

func (s *sLanguage) LanguageList() (res []mLanguage.Language, err error) {
	if err = s.orm.Model(&res).Where("shop_id = ?", s.shopId).
		Omit("shop_id", "created_at", "deleted_at", "updated_at").Find(&res).Error; err != nil {
		return res, err
	}
	return res, err
}
