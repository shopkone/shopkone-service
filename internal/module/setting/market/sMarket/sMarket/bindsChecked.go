package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/language/mLanguage"
	"shopkone-service/internal/module/setting/language/sLanguage"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
)

func (s *sMarket) BindsChecked() (err error) {
	markets, err := s.MarketOptions()
	if err != nil {
		return err
	}
	marketIds := slice.Map(markets, func(index int, item vo.MarketOptionsRes) uint {
		return item.Value
	})
	marketLanguages, err := s.BindListALl()
	if err != nil {
		return err
	}

	// 每个marketId必须有一个语言
	isAllHasLanguages := slice.Every(marketIds, func(index int, marketId uint) bool {
		_, ok := slice.FindBy(marketLanguages, func(index int, item mMarket.MarketLanguage) bool {
			return item.MarketID == marketId
		})
		return ok
	})
	if !isAllHasLanguages {
		return code.MarketMustLanguage
	}

	// 每个marketId必须有一个默认语言
	isAllHasDefaultLanguages := slice.Every(marketIds, func(index int, marketId uint) bool {
		_, ok := slice.FindBy(marketLanguages, func(index int, item mMarket.MarketLanguage) bool {
			return item.MarketID == marketId && item.IsDefault
		})
		return ok
	})
	if !isAllHasDefaultLanguages {
		return code.MarketMustDefaultLanguage
	}

	// 每个market_id必须可用
	isAllCanMarket := slice.Every(marketLanguages, func(index int, bind mMarket.MarketLanguage) bool {
		_, ok := slice.FindBy(markets, func(index int, market vo.MarketOptionsRes) bool {
			return market.Value == bind.MarketID
		})
		return ok
	})
	if !isAllCanMarket {
		return code.MarketValid
	}

	// 每个languages_id必须可用
	languages, err := sLanguage.NewLanguage(s.orm, s.shopId).LanguageList()
	if err != nil {
		return err
	}
	isAllCanLangs := slice.Every(marketLanguages, func(index int, bind mMarket.MarketLanguage) bool {
		_, ok := slice.FindBy(languages, func(index int, language mLanguage.Language) bool {
			return language.ID == bind.LanguageID
		})
		return ok
	})
	if !isAllCanLangs {
		return code.LanguageValid
	}

	return err
}
