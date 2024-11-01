package code

var (
	UserIsRegistered         = userError(1, "用户已注册")
	UserUnRegistered         = userError(2, "用户未注册")
	UserIsBlocked            = userError(3, "该账户已被冻结")
	UserPwdError             = userError(4, "密码错误")
	UserColumnTypeErr        = userError(5, "列类型错误")
	UserColumnNameRepeatErr  = userError(6, "列名重复")
	UserColumnListEmptyErr   = userError(7, "列列表为空")
	UserColumnListTooLongErr = userError(8, "列列表过长")
)
