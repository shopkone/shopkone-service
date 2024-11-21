package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

type CustomerCreateReq struct {
	g.Meta    `path:"/customer/create" method:"post" summary:"创建客户" tags:"Customer"`
	FirstName string               `json:"first_name" dc:"名字"`
	LastName  string               `json:"last_name" dc:"姓氏"`
	Email     string               `json:"email" dc:"邮箱"`
	Note      string               `json:"note" dc:"备注"`
	Phone     mAddress.Phone       `json:"phone" dc:"电话"`
	Gender    mCustomer.GenderType `json:"gender"`
	Birthday  int64                `json:"birthday"`
	Address   mAddress.Address     `json:"address"`
}
type CustomerCreateRes struct {
	ID uint `json:"id"`
}
