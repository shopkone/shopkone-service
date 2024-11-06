package sMarket

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/language/sLanguage"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
)

func (s *sMarket) MarketCreate(in vo.MarketCreateReq) (res vo.MarketCreateRes, err error) {
	// 如果是主市场，则判断著市场是否已经存在
	if in.IsMain {
		var Main mMarket.Market
		if err = s.orm.Model(&Main).Where("shop_id = ? AND is_main = ?", s.shopId, true).
			Select("id").Find(&Main).Error; err != nil {
			return res, err
		}
		if Main.ID != 0 {
			return res, code.MarketMainExist
		}
	}

	// 创建市场
	var data mMarket.Market
	data.IsMain = in.IsMain
	data.Name = in.Name
	data.ShopId = s.shopId
	data.Status = mMarket.MarketStatusActive
	if err = s.orm.Create(&data).Error; err != nil {
		return res, err
	}

	// 创建国家
	if err = s.CountryCreate(in.CountryCodes, data.ID); err != nil {
		return res, err
	}

	// 绑定默认语言
	defaultLanguage, err := sLanguage.NewLanguage(s.orm, s.shopId).LanguageDefault()
	if err != nil {
		return res, err
	}
	bindIn := vo.MarketBindLangReq{}
	bindIn.Bind = []vo.LanguageBindItem{{
		LanguageId: defaultLanguage.ID,
		MarketId:   data.ID,
	}}
	if err = s.LanguageBind(bindIn); err != nil {
		return res, err
	}

	// 后置校验
	if res.RemoveNames, err = s.MarketCheck(in.Force, data.ID, data.Name); err != nil {
		return res, err
	}

	res.ID = data.ID
	return res, err
}
