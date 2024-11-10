package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) MarketUpdate(in vo.MarketUpdateReq) (res vo.MarketUpdateRes, err error) {
	// 获取市场信息
	var info mMarket.Market
	if err = s.orm.Model(&mMarket.Market{}).Where("id = ? AND shop_id = ?", in.ID, s.shopId).
		Omit("created_at", "updated_at", "deleted_at", "shop_id").First(&info).Error; err != nil {
		return res, err
	}

	// 如果是主市场
	if info.IsMain {
		// 不能修改名称
		in.Name = info.Name
		// 国家只能有一个
		in.CountryCodes = slice.Filter(in.CountryCodes, func(index int, item string) bool {
			return index == 0
		})
	}

	// 更新市场信息
	if err = s.orm.Model(&mMarket.Market{}).Where("id = ? AND shop_id = ?", in.ID, s.shopId).
		Update("name", in.Name).Error; err != nil {
		return res, err
	}

	// 更新国家绑定
	if err = s.CountryUpdate(in.CountryCodes, in.ID); err != nil {
		return res, err
	}

	// 后置校验
	if res.RemoveNames, err = s.MarketCheck(in.Force, info.ID); err != nil {
		return res, err
	}

	return res, err
}
