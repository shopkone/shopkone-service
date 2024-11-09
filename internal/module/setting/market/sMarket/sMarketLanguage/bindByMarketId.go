package sMarketLanguage

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

// 根据market_id绑定语言
func (s *sMarketLanguage) BindLangByMarketId(in *vo.BindLangByMarketIdReq) (err error) {
	// 获取旧的
	var oldList []mMarket.MarketLanguage
	if err = s.orm.Model(&oldList).
		Where("market_id = ? AND shop_id = ?", in.MarketID, s.shopId).
		Omit("created_at", "updated_at", "deleted_at", "shop_id").Find(&oldList).Error; err != nil {
		return err
	}

	newList := slice.Map(in.LanguageIDs, func(index int, languageId uint) mMarket.MarketLanguage {
		i := mMarket.MarketLanguage{}
		i.ShopId = s.shopId
		i.MarketID = in.MarketID
		i.LanguageID = languageId
		return i
	})

	// 删除
	removes := slice.Filter(oldList, func(index int, oldItem mMarket.MarketLanguage) bool {
		_, ok := slice.FindBy(newList, func(index int, newItem mMarket.MarketLanguage) bool {
			return oldItem.MarketID == newItem.MarketID && oldItem.LanguageID == newItem.LanguageID
		})
		return !ok
	})
	if len(removes) > 0 {
		removeIds := slice.Map(removes, func(index int, item mMarket.MarketLanguage) uint {
			return item.ID
		})
		if err = s.orm.Model(&mMarket.MarketLanguage{}).Unscoped().
			Where("shop_id = ? AND id IN ?", s.shopId, removeIds).
			Delete(&mMarket.MarketLanguage{}).Error; err != nil {
			return err
		}
	}

	// 新增
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

	// 更新默认
	if err = s.orm.Model(&mMarket.MarketLanguage{}).
		Where("shop_id = ? AND is_default = ?", s.shopId, true).
		Update("is_default", false).Error; err != nil {
		return err
	}
	return s.orm.Model(&mMarket.MarketLanguage{}).
		Where("language_id = ? AND shop_id = ?", in.DefaultLanguageID, s.shopId).
		Update("is_default", true).Error
}
