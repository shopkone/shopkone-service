package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sMarket) MarketUpdateLangByLangID(req *vo.MarketUpdateLangByLangIDReq) (err error) {
	// 获取市场列表
	list, err := s.MarketOptions(nil)
	if err != nil {
		return err
	}
	mainMarket, ok := slice.FindBy(list, func(index int, item vo.MarketOptionsRes) bool {
		return item.IsMain
	})
	if !ok {
		return code.SystemError
	}
	markets := slice.Map(list, func(index int, market vo.MarketOptionsRes) mMarket.Market {
		enterHas := slice.Contain(req.MarketIds, market.Value)
		i := mMarket.Market{}
		i.ID = market.Value
		langIds := market.LanguageIds
		if enterHas {
			langIds = slice.Concat(langIds, []uint{req.LangId})
		} else {
			langIds = slice.Filter(langIds, func(index int, langId uint) bool {
				return langId != req.LangId
			})
		}
		i.LanguageIds = slice.Unique(langIds)
		i.CanCreateId = true
		i.ShopId = s.shopId
		if slice.Equal(langIds, market.LanguageIds) {
			i.ID = 0
		}
		if !slice.Contain(i.LanguageIds, market.DefaultLanguageId) {
			i.DefaultLanguageID = i.LanguageIds[0]
		}
		i.DomainType = market.DomainType
		return i
	})
	markets = slice.Map(markets, func(index int, market mMarket.Market) mMarket.Market {
		// 其他市场，使用主域名，则使用主要市场的配置
		if market.DomainType == mMarket.DomainTypeMain {
			market.LanguageIds, market.DefaultLanguageID = syncMarketConfig(markets, mainMarket.Value)
		}
		return market
	})
	markets = slice.Filter(markets, func(index int, item mMarket.Market) bool {
		return item.ID != 0
	})
	if len(markets) == 0 {
		return err
	}
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"language_ids", "default_language_id"},
	}
	if err = handle.BatchUpdateById(batchIn, &markets); err != nil {
		return err
	}
	return s.MarketCheckValid()
}

func syncMarketConfig(markets []mMarket.Market, mainMarketID uint) (langIds []uint, defaultLangId uint) {
	market, _ := slice.FindBy(markets, func(index int, item mMarket.Market) bool {
		return item.ID == mainMarketID
	})
	return market.LanguageIds, market.DefaultLanguageID
}
