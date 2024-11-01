package mResource

/*Category 分类*/
type Category struct {
	Value   uint   `json:"value"`
	Label   string `json:"label"`
	EnLabel string `json:"en_label"`
	Pid     uint   `json:"pid"`
	Deep    uint   `json:"deep"`
}

/*Currency 货币*/
type Currency struct {
	Code        string `json:"code" gorm:"primary_key"` // 货币代码
	SymbolLeft  string `json:"symbol_left"`             // 货币符号左
	SymbolRight string `json:"symbol_right"`            // 货币符号右
	Title       string `json:"title"`                   // 货币名称
}

/*Timezone 时区*/
type Timezone struct {
	OlsonName   string `json:"olsonName"`   // 时区名称
	Description string `json:"description"` // 时区描述
}

/*承运商*/
type Carrier struct {
	Id                       uint   `json:"id"`
	Name                     string `json:"name"`
	DisplayName              string `json:"displayName"`
	SupportsShipmentTracking bool   `json:"supportsShipmentTracking"`
}

/*税*/
type Tax struct {
	CountryCode string  `json:"country_code"`
	Tax         float64 `json:"tax"`
}
