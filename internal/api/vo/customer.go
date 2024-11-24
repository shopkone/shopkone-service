package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/handle"
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
	Tags      []string             `json:"tags"`
}
type CustomerCreateRes struct {
	ID uint `json:"id"`
}

type CustomerInfoReq struct {
	g.Meta `path:"/customer/info" method:"post" summary:"客户详情" tags:"Customer"`
	ID     uint `json:"id" v:"required" dc:"客户ID"`
}
type CustomerInfoRes struct {
	ID        uint                 `json:"id"`
	FirstName string               `json:"first_name" dc:"名字"`
	LastName  string               `json:"last_name" dc:"姓氏"`
	Email     string               `json:"email" dc:"邮箱"`
	Note      string               `json:"note" dc:"备注"`
	Phone     string               `json:"phone" dc:"电话"`
	Gender    mCustomer.GenderType `json:"gender"`
	Birthday  int64                `json:"birthday"`
	Address   []mAddress.Address   `json:"address"`
	Tags      []string             `json:"tags"`
}

type CustomerListReq struct {
	g.Meta `path:"/customer/list" method:"post" summary:"客户列表" tags:"Customer"`
	handle.PageReq
}
type CustomerListRes struct {
	ID             uint    `json:"id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	OrderCount     int64   `json:"order_count"`
	CostPrice      float32 `json:"cost_price"`
	EmailSubscribe bool    `json:"email_subscribe"`
	CountryInfo    string  `json:"country_info"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
}
