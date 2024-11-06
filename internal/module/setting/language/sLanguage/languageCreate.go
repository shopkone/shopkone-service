package sLanguage

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/setting/language/mLanguage"
	"shopkone-service/utility/code"
)

func (s *sLanguage) LanguageCreate(languages []string, isDefault bool) (ids []uint, err error) {
	// 校验语言是否有效
	isAllHas := slice.Every(languages, func(index int, item string) bool {
		_, has := slice.FindBy(resource.Languages, func(index int, l string) bool {
			return l == item
		})
		return has
	})
	if !isAllHas {
		return ids, code.LanguageValid
	}
	// 判断语言是否已经存在于该店铺
	var count int64
	if err = s.orm.Model(&mLanguage.Language{}).Where("shop_id = ?", s.shopId).
		Where("code IN ?", languages).
		Count(&count).Error; err != nil {
		return ids, err
	}
	if count != 0 {
		return ids, code.LanguageRepeat
	}
	// 开始添加语言
	data := slice.Map(languages, func(index int, item string) mLanguage.Language {
		i := mLanguage.Language{}
		i.Code = item
		i.ShopId = s.shopId
		i.IsDefault = isDefault
		i.IsActive = isDefault
		return i
	})
	if err = s.orm.Create(&data).Error; err != nil {
		return ids, err
	}
	ids = slice.Map(data, func(index int, item mLanguage.Language) uint {
		return item.ID
	})
	return ids, err
}
