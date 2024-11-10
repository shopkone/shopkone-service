package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
)

func (s *sMarket) MarketCheck(force bool, id uint) (res []string, err error) {
	market, err := s.MarketInfo(id)
	if err != nil {
		return nil, err
	}

	// 判断是否有默认市场
	if market.DefaultLanguageId == 0 {
		return res, code.MarketMustDefaultLanguage
	}

	// 查看没有国家的市场
	noCountryMarkets, err := s.MarketListByUnCountry()
	if err != nil {
		return res, err
	}

	// 如果是主要国家，则报错
	_, isMain := slice.FindBy(noCountryMarkets, func(index int, item mMarket.Market) bool {
		return item.IsMain
	})
	if isMain {
		return res, code.MainMarketCanNotMove
	}

	// 如果没有强制删除，则返错误
	if !force && len(noCountryMarkets) > 0 {
		res = slice.Map(noCountryMarkets, func(index int, item mMarket.Market) string {
			return item.Name
		})
		return res, code.MarketNoCountry
	}

	// 删除这些不存在国家的市场
	noCountryMarketIds := slice.Map(noCountryMarkets, func(index int, item mMarket.Market) uint {
		return item.ID
	})
	if err = s.MarketRemoveByIds(noCountryMarketIds); err != nil {
		return res, err
	}

	// 校验名称是否重复
	var count int64
	if err = s.orm.Model(mMarket.Market{}).Where("name = ? AND id != ? AND shop_id = ?", market.Name, id, s.shopId).
		Select("id").Count(&count).Error; err != nil {
		return res, err
	}
	if count > 0 {
		return res, code.MarketNameExist
	}

	// 如果是子文件夹，校验子文件名称是否重复
	if market.DomainType == mMarket.DomainTypeSuffix {
		if err = s.orm.Model(&mMarket.Market{}).
			Where("shop_id = ? AND id != ? AND domain_suffix = ? AND domain_type = ?",
				s.shopId, id, market.DomainSuffix, mMarket.DomainTypeSuffix,
			).Count(&count).Error; err != nil {
			return nil, err
		}
		if count > 0 {
			return res, code.MarketSuffixExist
		}
	}

	return res, err
}
