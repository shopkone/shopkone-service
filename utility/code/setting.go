package code

var (
	MarketMainExist        = SettingError(1, "主市场已经存在")
	MarketNoCountry        = SettingError(2, "会导致出现没有绑定国家的市场")
	MainMarketCanNotMove   = SettingError(3, "不允许存在主要市场中的国家")
	MarketNameExist        = SettingError(4, "市场名称已经存在")
	LanguageRepeat         = SettingError(5, "存在重复的语言")
	LanguageValid          = SettingError(6, "无效的语言")
	MarketMustLanguage     = SettingError(7, "市场必须绑定语言")
	DomainAlreadyConnected = SettingError(8, "该域名已被使用，请绑定其他域名")
	DomainNotRegistered    = SettingError(9, "该域名还未注册")
	DomainUnknown          = SettingError(10, "该域名异常")
	DomainValid            = SettingError(11, "无效的域名")
	MarketMustSubDomain    = SettingError(12, "请选择一个可用的子域名")
	MarketMustPrefixDomain = SettingError(13, "请填写域名后缀")
)
