package sLanguage

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/language/mLanguage"
)

func (s *sLanguage) LanguageFilter(languageIds []uint) (ids []uint, err error) {
	languages, err := s.LanguageList()
	if err != nil {
		return nil, err
	}
	return slice.Filter(languageIds, func(index int, item uint) bool {
		find, ok := slice.FindBy(languages, func(index int, i mLanguage.Language) bool {
			return i.ID == item
		})
		return ok && find.ID != 0
	}), nil
}
