package sMarket

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
)

func (s *sMarket) MarketCreate(in vo.MarketCreateReq) (id uint, err error) {
	// 如果是主市场，则判断著市场是否已经存在
	if in.IsMain {
		var Main mMarket.Market
		if err = s.orm.Model(&Main).Where("shop_id = ? AND is_main = ?", s.shopId, true).
			Select("id").Find(&Main).Error; err != nil {
			return 0, err
		}
		if Main.ID != 0 {
			return 0, code.MarketMainExist
		}
	}

	// 创建市场
	var data mMarket.Market
	data.IsMain = in.IsMain
	data.Name = in.Name
	data.ShopId = s.shopId
	data.Status = mMarket.MarketStatusActive
	if err = s.orm.Create(&data).Error; err != nil {
		return 0, err
	}

	// 创建国家
	return data.ID, s.CountryCreate(in.CountryCodes, data.ID)
}
