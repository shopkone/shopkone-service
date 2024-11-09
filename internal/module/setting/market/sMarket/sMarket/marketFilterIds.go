package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
)

func (s *sMarket) MarketFilterIds(marketIds []uint) (ids []uint, err error) {
	// 获取markets
	markets, err := s.MarketList()
	if err != nil {
		return nil, err
	}

	slice.ForEach(marketIds, func(index int, item uint) {
		find, ok := slice.FindBy(markets, func(index int, i vo.MarketListRes) bool {
			return i.ID == item
		})
		if ok {
			ids = append(ids, find.ID)
		}
	})
	return ids, err
}
