package sCache

// 常量定义
const (
	EXCHANGE_RATE_PREFIX_KEY = "exchangerate_" // 汇率
	USER_PREFIX_KEY          = "user_"
	SHOP_PREFIX_KEY          = "shop_"
	STAFF_PREFIX_KEY         = "staff_"
	CAPTCHA_PREFIX_KEY       = "captcha_"
	TOKEN_PREFIX_KEY         = "shopkonetoken_"
	TEEMPLATE_PREFIX_KEY     = "template_"
)

// 允许的前缀集合
var allowedPrefixes = map[string]struct{}{
	EXCHANGE_RATE_PREFIX_KEY: {},
	USER_PREFIX_KEY:          {},
	SHOP_PREFIX_KEY:          {},
	STAFF_PREFIX_KEY:         {},
	CAPTCHA_PREFIX_KEY:       {},
	TOKEN_PREFIX_KEY:         {},
	TEEMPLATE_PREFIX_KEY:     {},
}
