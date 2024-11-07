package code

var (
	CaptchaMaxSendTimes = baseError(1, "验证码发送次数已达到上限")
	CaptchaMaxTryTimes  = baseError(2, "验证码尝试次数已达到上限")
	CaptchaCodeExpired  = baseError(3, "验证码已失效，请重新发送")
	CaptchaCodeError    = baseError(4, "验证码错误")
	SenseError          = baseError(5, "场景值错误")
	AuthError           = baseError(6, "认证失败")
	LoginSystemLock     = baseError(7, "系统繁忙，请稍后再试")
	CountryCodeUnknown  = baseError(8, "不支持的国家代码")
	ZoneCodeUnknown     = baseError(9, "不支持的省份代码")
	ErrAddressId        = baseError(10, "地址ID错误")
	IdMissing           = baseError(11, "缺少参数")
	SystemError         = baseError(12, "系统错误")
)
