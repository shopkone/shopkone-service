package resource

import (
	_ "embed"
	"encoding/json"
	"shopkone-service/internal/module/base/resource/mResource"
)

//go:embed constant/category.json
var categories []byte

//go:embed constant/country.json
var country []byte

//go:embed constant/currency.json
var currency []byte

//go:embed constant/languages.json
var languages []byte

//go:embed constant/timezone.json
var timezone []byte

//go:embed constant/carriers.json
var carriers []byte

//go:embed constant/tax.json
var tax []byte

var Countries []mResource.Country
var Timezones []mResource.Timezone
var Currencies []mResource.Currency
var Languages []string
var Categories []mResource.Category
var Carriers []mResource.Carrier
var Tax []mResource.Tax

func InitResource() error {
	// 分类
	if err := json.Unmarshal(categories, &Categories); err != nil {
		return err
	}
	// 国家
	if err := json.Unmarshal(country, &Countries); err != nil {
		return err
	}
	// 货币
	if err := json.Unmarshal(currency, &Currencies); err != nil {
		return err
	}
	// 时区
	if err := json.Unmarshal(timezone, &Timezones); err != nil {
		return err
	}
	// 语言
	if err := json.Unmarshal(languages, &Languages); err != nil {
		return err
	}
	// 快递
	if err := json.Unmarshal(carriers, &Carriers); err != nil {
		return err
	}
	// 税率
	if err := json.Unmarshal(tax, &Tax); err != nil {
		return err
	}
	return nil
}
