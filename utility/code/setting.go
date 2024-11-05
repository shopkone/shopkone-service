package code

var (
	MarketMainExist      = SettingError(1, "主市场已经存在")
	MarketNoCountry      = SettingError(2, "会导致出现没有绑定国家的市场")
	MainMarketCanNotMove = SettingError(3, "不允许存在主要市场中的国家")
	MarketNameExist      = SettingError(4, "市场名称已经存在")
	LanguageRepeat       = SettingError(5, "存在重复的语言")
	LanguageValid        = SettingError(6, "无效的语言")
)
