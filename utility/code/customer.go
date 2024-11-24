package code

var (
	ErrCusterCreateErr    = CustomerError(1, "名字、姓氏、电子邮件、电话号码，不能同时为空")
	ErrCustomerEmailExist = CustomerError(2, "此邮箱已绑定至其他客户")
	ErrCustomerPhoneExist = CustomerError(3, "此手机号码已绑定至其他客户")
)
