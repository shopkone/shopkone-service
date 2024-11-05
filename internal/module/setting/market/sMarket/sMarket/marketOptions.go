package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) MarketOptions() (res []vo.MarketOptionsRes, err error) {
	var list []mMarket.Market
	if err = s.orm.Model(&mMarket.Market{}).Where("shop_id = ?", s.shopId).
		Select("id", "name", "is_main").Find(&list).Error; err != nil {
		return nil, err
	}
	res = slice.Map(list, func(index int, item mMarket.Market) vo.MarketOptionsRes {
		return vo.MarketOptionsRes{
			Label:  item.Name,
			Value:  item.ID,
			IsMain: item.IsMain,
		}
	})
	return res, err
}
