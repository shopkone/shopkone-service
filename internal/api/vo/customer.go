package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/utility/handle"
)

// 创建客户
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

// 获取客户详情
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
	Phone     mAddress.Phone       `json:"phone" dc:"电话"`
	Gender    mCustomer.GenderType `json:"gender"`
	Birthday  int64                `json:"birthday"`
	Language  string               `json:"language"`
	Address   []mAddress.Address   `json:"address"`
	Tags      []string             `json:"tags"`
}

// 获取客户列表
type CustomerListReq struct {
	g.Meta `path:"/customer/list" method:"post" summary:"客户列表" tags:"Customer"`
	handle.PageReq
}
type CustomerListRes struct {
	ID             uint             `json:"id"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	OrderCount     int64            `json:"order_count"`
	CostPrice      float32          `json:"cost_price"`
	EmailSubscribe bool             `json:"email_subscribe"`
	CountryInfo    string           `json:"country_info"`
	Email          string           `json:"email"`
	Phone          string           `json:"phone"`
	Address        mAddress.Address `json:"address"`
}

// 更新客户信息
type CustomerUpdateBaseReq struct {
	g.Meta    `path:"/customer/update/base" method:"post" summary:"更新客户" tags:"Customer"`
	ID        uint                 `json:"id" v:"required" dc:"客户ID"`
	FirstName string               `json:"first_name" dc:"名字"`
	LastName  string               `json:"last_name" dc:"姓氏"`
	Language  string               `json:"language"`
	Email     string               `json:"email" dc:"邮箱"`
	Phone     mAddress.Phone       `json:"phone"`
	Gender    mCustomer.GenderType `json:"gender"`
	Birthday  int64                `json:"birthday"`
}
type CustomerUpdateBaseRes struct {
}

// 添加收货地址
type CustomerAddAddress struct {
	g.Meta `path:"/customer/add/address" method:"post" summary:"添加地址" tags:"Customer"`
	mAddress.Address
}
type CustomerAddAddressRes struct {
}

// 更新收货地址
type CustomerUpdateAddress struct {
	g.Meta `path:"/customer/update/address" method:"post" summary:"更新地址" tags:"Customer"`
	mAddress.Address
	IsDefault bool `json:"is_default"`
}
type CustomerUpdateAddressRes struct {
}

type CustomerSetTaxItem struct {
	ID          uint     `json:"id" v:"required" dc:"单个地区id"`
	CountryCode string   `json:"country_code" v:"required" dc:"国家代码"`
	Zones       []string `json:"zones" v:"required" dc:"地区"`
}

// 设置免税地区
type CustomerSetTaxReq struct {
	g.Meta `path:"/customer/set/tax" method:"post" summary:"设置免税地区" tags:"Customer"`
	ID     uint                 `json:"id" v:"required" dc:"客户ID"`
	Areas  []CustomerSetTaxItem `json:"areas" v:"required" dc:"免税地区"`
	Free   bool                 `json:"free"`
	All    bool                 `json:"all"`
}
type CustomerSetTaxRes struct {
}

// 更新标签
type CustomerUpdateTagsReq struct {
	g.Meta `path:"/customer/update/tags" method:"post" summary:"更新标签" tags:"Customer"`
	ID     uint     `json:"id" v:"required" dc:"客户ID"`
	Tags   []string `json:"tags" dc:"标签ID"`
}
type CustomerUpdateTagsRes struct {
}

// 更新备注
type CustomerUpdateNoteReq struct {
	g.Meta `path:"/customer/update/note" method:"post" summary:"更新备注" tags:"Customer"`
	ID     uint   `json:"id" v:"required" dc:"客户ID"`
	Note   string `json:"note" dc:"备注"`
}
type CustomerUpdateNoteRes struct {
}
