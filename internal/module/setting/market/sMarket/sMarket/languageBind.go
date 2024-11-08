package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

// 语言的维度绑定
func (s *sMarket) LanguageBind(in vo.MarketBindLangReq) (err error) {
	// 先删除停绑的
	if len(in.UnBind) != 0 {
		deleteQuery := s.orm.Model(&mMarket.MarketLanguage{})
		slice.ForEach(in.UnBind, func(index int, item vo.LanguageBindItem) {
			deleteQuery = deleteQuery.Or("language_id = ? AND market_id = ? AND shop_id = ?", item.LanguageId, item.MarketId, s.shopId)
		})
		if err = deleteQuery.Unscoped().Delete(&mMarket.MarketLanguage{}).Error; err != nil {
			return err
		}
	}

	// 添加新邦的
	if len(in.Bind) != 0 {
		insertData := slice.Map(in.Bind, func(index int, item vo.LanguageBindItem) mMarket.MarketLanguage {
			i := mMarket.MarketLanguage{}
			i.MarketID = item.MarketId
			i.LanguageID = item.LanguageId
			i.ShopId = s.shopId
			i.IsDefault = item.IsDefault
			return i
		})
		if err = s.orm.Create(&insertData).Error; err != nil {
			return err
		}
	}

	// 后置校验
	return s.LanguageCheck()
}
