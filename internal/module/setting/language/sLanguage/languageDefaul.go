package sLanguage

import "shopkone-service/internal/module/setting/language/mLanguage"

// 获取默认语言
func (s *sLanguage) LanguageDefault() (lang mLanguage.Language, err error) {
	if err = s.orm.Model(&lang).Where("shop_id = ? AND is_default = ?", s.shopId, true).
		Select("id").First(&lang).Error; err != nil {
		return lang, err
	}
	return lang, nil
}
