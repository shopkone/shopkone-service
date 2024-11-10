package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sMarket) MarketUpdateLang(req vo.MarketUpdateLangReq) (err error) {
	data := mMarket.Market{}
	if err = s.orm.Model(&data).Where("shop_id = ? AND id = ?", s.shopId, req.ID).
		Select("domain_type", "is_main").First(&data).Error; err != nil {
		return err
	}

	if !(data.DomainType == mMarket.DomainTypeMain && !data.IsMain) {
		data.DefaultLanguageID = req.DefaultLanguageID
		data.LanguageIds = req.LanguageIds

		if !slice.Contain(data.LanguageIds, req.DefaultLanguageID) {
			data.DefaultLanguageID = 0
		}
		if data.DefaultLanguageID == 0 {
			data.DefaultLanguageID = req.LanguageIds[0]
		}

		if err = s.orm.Model(&data).
			Where("shop_id = ? AND id = ?", s.shopId, req.ID).
			Select("default_language_id", "language_ids").
			Updates(&data).Error; err != nil {
			return err
		}
	}

	// 如果是主域名，则同步市场的主域名
	if data.DomainType != mMarket.DomainTypeMain {
		return err
	}
	options, err := s.MarketOptions()
	if err != nil {
		return err
	}
	otherMarketMainDomain := slice.Filter(options, func(index int, opt vo.MarketOptionsRes) bool {
		return !opt.IsMain && opt.DomainType == mMarket.DomainTypeMain
	})
	if len(otherMarketMainDomain) == 0 {
		return nil
	}
	mainMarket, ok := slice.FindBy(options, func(index int, opt vo.MarketOptionsRes) bool {
		return opt.IsMain && opt.DomainType == mMarket.DomainTypeMain
	})
	if !ok {
		return code.SystemError
	}
	otherMarkets := slice.Map(otherMarketMainDomain, func(index int, item vo.MarketOptionsRes) mMarket.Market {
		i := mMarket.Market{}
		i.ID = item.Value
		i.ShopId = s.shopId
		i.CanCreateId = true
		i.DefaultLanguageID = mainMarket.DefaultLanguageId
		i.LanguageIds = mainMarket.LanguageIds
		return i
	})
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"default_language_id", "language_ids"},
	}
	return handle.BatchUpdateById(batchIn, &otherMarkets)
}
