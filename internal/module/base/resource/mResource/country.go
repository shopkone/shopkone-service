package mResource

// 地址标签
type AddressLabel struct {
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	Company    string `json:"company"`
	Country    string `json:"country"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postalCode"`
	Zone       string `json:"zone"`
}

// 地址格式化
type AddressFormatting struct {
	Edit string `json:"edit"`
	Show string `json:"show"`
}

// 邮编配置
type PostalCodeConfig struct {
	Format string `json:"format"`
	Regex  string `json:"regex"`
}

// 省份
type CountryZone struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// 国家国旗
type CountryFlag struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

// 国家
type Country struct {
	Name              string            `json:"name"`              // 名称
	Code              string            `json:"code"`              // ISO 3166-1 alpha-2
	Continent         string            `json:"continent"`         // 大洲
	PhoneNumberPrefix int               `json:"phoneNumberPrefix"` // 电话号码前缀
	ProvinceKey       string            `json:"provinceKey"`       // 省的key
	AddressLabel      AddressLabel      `json:"labels"`            // 标签
	AddressFormatting AddressFormatting `json:"formatting"`        // 格式化
	Zones             []CountryZone     `json:"zones"`             // 省份
	Flag              CountryFlag       `json:"flag"`              // 国旗
	Currencies        []string          `json:"currencies"`        // 货币
	Timezones         []string          `json:"timezones"`         // 时区
	PostalCodeConfig  PostalCodeConfig  `json:"postalCode"`        // 邮编配置
}
