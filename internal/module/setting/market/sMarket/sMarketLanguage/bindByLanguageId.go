package sMarketLanguage

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

// BindUpdateByLanguageId 根据语言id更新语言市场绑定
func (s *sMarketLanguage) BindByLanguageId(req *vo.BindLangByLangIdReq) (err error) {
	// 获取旧的
	var oldList []mMarket.MarketLanguage
	if err = s.orm.Model(&oldList).
		Where("language_id = ? AND shop_id = ?", req.LanguageID, s.shopId).
		Omit("created_at", "updated_at", "deleted_at", "shop_id").Find(&oldList).Error; err != nil {
		return err
	}

	// 新的
	newList := slice.Map(req.MarketIDs, func(index int, marketId uint) mMarket.MarketLanguage {
		i := mMarket.MarketLanguage{}
		i.ShopId = s.shopId
		i.MarketID = marketId
		i.LanguageID = req.LanguageID
		return i
	})

	// 获取删除的
	removes := slice.Filter(oldList, func(index int, oldItem mMarket.MarketLanguage) bool {
		_, ok := slice.FindBy(newList, func(index int, newItem mMarket.MarketLanguage) bool {
			return oldItem.MarketID == newItem.MarketID && oldItem.LanguageID == newItem.LanguageID
		})
		return !ok
	})
	if len(removes) > 0 {
		if err = s.orm.Model(&mMarket.MarketLanguage{}).Unscoped().
			Delete(removes).Error; err != nil {
			return err
		}
	}

	// 获取新增
	adds := slice.Filter(newList, func(index int, newItem mMarket.MarketLanguage) bool {
		_, ok := slice.FindBy(oldList, func(index int, oldItem mMarket.MarketLanguage) bool {
			return oldItem.MarketID == newItem.MarketID && oldItem.LanguageID == newItem.LanguageID
		})
		return !ok
	})
	if len(adds) > 0 {
		if err = s.orm.Model(&mMarket.MarketLanguage{}).Create(adds).Error; err != nil {
			return err
		}
	}
	return err
}
