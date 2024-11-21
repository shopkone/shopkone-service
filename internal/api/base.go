package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	exchang_rate "shopkone-service/internal/module/base/exchang-rate"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/base/resource/iResource"
	"shopkone-service/internal/module/base/resource/mResource"
	"shopkone-service/internal/module/base/resource/sResource"
	ctx2 "shopkone-service/utility/ctx"
)

type aBaseApi struct {
}

func NewBaseApi() *aBaseApi {
	return &aBaseApi{}
}

// Countries 获取国家列表
func (a *aBaseApi) Countries(ctx g.Ctx, req *vo.CountriesReq) (res []vo.CountriesRes, err error) {
	t := ctx2.NewCtx(ctx).GetT
	list := sResource.NewCountry().List()
	res = slice.Map(list, func(index int, item iResource.CountryListOut) vo.CountriesRes {
		zones := slice.Map(item.Zones, func(index int, zone iResource.ZoneListOut) iResource.ZoneListOut {
			zone.Name = t(item.Name + "___" + zone.Name)
			return zone
		})
		return vo.CountriesRes{
			Code:      item.Code,
			Continent: item.Continent,
			Name:      t(item.Name),
			Zones:     zones,
			Flag:      item.Flag,
			Config: vo.AddressConfig{
				Address1:   t(item.Config.Address1),
				Address2:   t(item.Config.Address2),
				City:       t(item.Config.City),
				Company:    t(item.Config.Company),
				Country:    t(item.Config.Country),
				FirstName:  t(item.Config.FirstName),
				LastName:   t(item.Config.LastName),
				Phone:      t(item.Config.Phone),
				PostalCode: t(item.Config.PostalCode),
				Zone:       t(item.Config.Zone),
			},
			Formatting:       item.Formatting.Show,
			PostalCodeConfig: item.PostalCodeConfig,
		}
	})
	return res, err
}

// PhoneAreaCode 获取手机区号列表
func (a *aBaseApi) PhonePrefix(ctx g.Ctx, req *vo.PhonePrefixReq) (res []vo.PhonePrefixRes, err error) {
	list := sResource.NewCountry().PhonePrefixList()
	res = slice.Map(list, func(index int, item iResource.PhonePrefixListOut) vo.PhonePrefixRes {
		return vo.PhonePrefixRes{
			Code:   item.Code,
			Prefix: item.Prefix,
		}
	})
	return res, err
}

// TimezoneList 获取时区列表
func (a *aBaseApi) TimezoneList(ctx g.Ctx, req *vo.TimezoneListReq) (res []vo.TimezoneListRes, err error) {
	t := ctx2.NewCtx(ctx).GetT
	res = slice.Map(resource.Timezones, func(index int, item mResource.Timezone) vo.TimezoneListRes {
		return vo.TimezoneListRes{
			Description: t(item.Description),
			OlsonName:   item.OlsonName,
		}
	})
	return res, err
}

// CurrencyList 获取货币列表
func (a *aBaseApi) CurrencyList(ctx g.Ctx, req *vo.CurrencyListReq) (res []vo.CurrencyListRes, err error) {
	t := ctx2.NewCtx(ctx).GetT
	res = slice.Map(resource.Currencies, func(index int, item mResource.Currency) vo.CurrencyListRes {
		return vo.CurrencyListRes{
			Code:   item.Code,
			Symbol: item.SymbolLeft,
			Title:  t(item.Title),
		}
	})
	return res, err
}

// CategoryList 获取分类列表
func (a *aBaseApi) CategoryList(ctx g.Ctx, req *vo.CategoryListReq) (res []vo.CategoryListRes, err error) {
	t := ctx2.NewCtx(ctx).GetT
	res = slice.Map(resource.Categories, func(index int, item mResource.Category) vo.CategoryListRes {
		return vo.CategoryListRes{
			Value: item.Value,
			Label: t(item.EnLabel),
			Deep:  item.Deep,
			Pid:   item.Pid,
		}
	})
	return res, err
}

// CarrierList 获取物流商列表
func (a *aBaseApi) CarrierList(ctx g.Ctx, req *vo.CarrierListReq) (res []vo.CarrierListRes, err error) {
	res = slice.Map(resource.Carriers, func(index int, item mResource.Carrier) vo.CarrierListRes {
		return vo.CarrierListRes{
			Id:                       item.Id,
			Name:                     item.Name,
			DisplayName:              item.DisplayName,
			SupportsShipmentTracking: item.SupportsShipmentTracking,
		}
	})
	return res, err
}

// 获取语言列表
func (a *aBaseApi) LanguageList(ctx g.Ctx, req *vo.LanguagesReq) (res []vo.LanguagesRes, err error) {
	t := ctx2.NewCtx(ctx).GetT
	res = slice.Map(resource.Languages, func(index int, item string) vo.LanguagesRes {
		i := vo.LanguagesRes{}
		i.Label = t(item)
		i.Value = item
		return i
	})
	return res, err
}

// 汇率
func (a *aBaseApi) ExchangeRate(ctx g.Ctx, req *vo.GetExchangeRateReq) (res vo.GetExchangeRateRes, err error) {
	getRateIn := exchang_rate.GetRateIn{
		FromCode: req.From,
		ToCode:   req.To,
	}
	exchange, err := exchang_rate.NewExchangeRate().GetRate(getRateIn)
	if err != nil {
		return res, err
	}
	res.Rate = exchange.Rate
	res.TimeStamp = exchange.Timestamp
	return res, err
}
