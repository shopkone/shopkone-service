package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/base/resource/iResource"
	"shopkone-service/internal/module/base/resource/mResource"
	"shopkone-service/internal/module/base/resource/sResource"
)

type aBaseApi struct {
}

func NewBaseApi() *aBaseApi {
	return &aBaseApi{}
}

// Countries 获取国家列表
func (a *aBaseApi) Countries(ctx g.Ctx, req *vo.CountriesReq) (res []vo.CountriesRes, err error) {
	list := sResource.NewCountry().List()
	res = slice.Map(list, func(index int, item iResource.CountryListOut) vo.CountriesRes {
		return vo.CountriesRes{
			Code:      item.Code,
			Continent: item.Continent,
			Name:      item.Name,
			Zones:     item.Zones,
			Flag:      item.Flag,
			Config: vo.AddressConfig{
				Address1:   item.Config.Address1,
				Address2:   item.Config.Address2,
				City:       item.Config.City,
				Company:    item.Config.Company,
				Country:    item.Config.Country,
				FirstName:  item.Config.FirstName,
				LastName:   item.Config.LastName,
				Phone:      item.Config.Phone,
				PostalCode: item.Config.PostalCode,
				Zone:       item.Config.Zone,
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
	res = slice.Map(resource.Timezones, func(index int, item mResource.Timezone) vo.TimezoneListRes {
		return vo.TimezoneListRes{
			Description: item.Description,
			OlsonName:   item.OlsonName,
		}
	})
	return res, err
}

// CurrencyList 获取货币列表
func (a *aBaseApi) CurrencyList(ctx g.Ctx, req *vo.CurrencyListReq) (res []vo.CurrencyListRes, err error) {
	res = slice.Map(resource.Currencies, func(index int, item mResource.Currency) vo.CurrencyListRes {
		return vo.CurrencyListRes{
			Code:   item.Code,
			Symbol: item.SymbolLeft,
			Title:  item.Title,
		}
	})
	return res, err
}

// CategoryList 获取分类列表
func (a *aBaseApi) CategoryList(ctx g.Ctx, req *vo.CategoryListReq) (res []vo.CategoryListRes, err error) {
	res = slice.Map(resource.Categories, func(index int, item mResource.Category) vo.CategoryListRes {
		return vo.CategoryListRes{
			Value: item.Value,
			Label: item.EnLabel,
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
func (a *aBaseApi) LanguageList(ctx g.Ctx, req *vo.LanguagesReq) (res vo.LanguagesRes, err error) {
	res.List = resource.Languages
	return res, err
}
