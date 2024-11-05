package code

var (
	MarketMainExist      = SettingError(1, "主市场已经存在")
	MarketNoCountry      = SettingError(2, "市场没有国家")
	MainMarketCanNotMove = SettingError(3, "不允许存在主要市场中的国家")
)
