package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/base/resource/mResource"
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

	// 获取店铺默认语言
	defaultLanguage, err := sLanguage.NewLanguage(s.orm, s.shopId).LanguageDefault()
	if err != nil {
		return res, err
	}

	// 设置默认货币
	var currencyCodes []string
	slice.ForEach(resource.Countries, func(index int, item mResource.Country) {
		_, ok := slice.FindBy(in.CountryCodes, func(index int, code string) bool {
			return item.Code == code
		})
		if !ok {
			return
		}
		currencyCodes = append(currencyCodes, item.Currencies...)
	})
	// 过滤不在当前列表的货币
	currencyCodes = slice.Filter(currencyCodes, func(index int, code string) bool {
		_, ok := slice.FindBy(resource.Currencies, func(index int, item mResource.Currency) bool {
			return item.Code == code
		})
		return ok
	})
	// 找出出现次数最多的货币
	maxCurrencyCode := mostFrequentString(currencyCodes)
	if maxCurrencyCode == "" {
		maxCurrencyCode = "USD"
	}

	// 创建市场
	var data mMarket.Market
	data.IsMain = in.IsMain
	data.Name = in.Name
	data.ShopId = s.shopId
	data.Status = mMarket.MarketStatusActive
	data.DomainType = mMarket.DomainTypeMain
	data.CurrencyCode = maxCurrencyCode
	data.AdjustPercent = 0
	data.AdjustType = mMarket.PriceAdjustmentTypeAdd
	if err = s.orm.Create(&data).Error; err != nil {
		return res, err
	}

	// 更新语言
	updateLangIn := vo.MarketUpdateLangReq{
		ID:                data.ID,
		LanguageIds:       []uint{defaultLanguage.ID},
		DefaultLanguageID: defaultLanguage.ID,
	}
	if err = s.MarketUpdateLang(updateLangIn); err != nil {
		return vo.MarketCreateRes{}, err
	}

	// 创建国家
	if err = s.CountryCreate(in.CountryCodes, data.ID); err != nil {
		return res, err
	}

	// 后置校验
	if res.RemoveNames, err = s.MarketCheck(in.Force, data.ID); err != nil {
		return res, err
	}

	res.ID = data.ID
	return res, err
}

func mostFrequentString(strings []string) string {
	if len(strings) == 0 {
		return ""
	}

	// 创建一个 map 来存储每个字符串的出现次数
	frequencyMap := make(map[string]int)
	for _, s := range strings {
		frequencyMap[s]++
	}

	// 找到出现次数最多的字符串
	maxFreq := 0
	var mostFrequent string
	for str, freq := range frequencyMap {
		if freq > maxFreq {
			maxFreq = freq
			mostFrequent = str
		}
	}

	return mostFrequent
}
