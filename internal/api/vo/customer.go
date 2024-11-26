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

type CustomerFreeTax struct {
	Areas []CustomerSetTaxItem `json:"areas"`
	Free  bool                 `json:"free"`
	All   bool                 `json:"all"`
}

// 获取客户详情
type CustomerInfoReq struct {
	g.Meta `path:"/customer/info" method:"post" summary:"客户详情" tags:"Customer"`
	ID     uint `json:"id" v:"required" dc:"客户ID"`
}
type CustomerInfoRes struct {
	ID               uint                 `json:"id"`
	FirstName        string               `json:"first_name" dc:"名字"`
	LastName         string               `json:"last_name" dc:"姓氏"`
	Email            string               `json:"email" dc:"邮箱"`
	Note             string               `json:"note" dc:"备注"`
	Phone            mAddress.Phone       `json:"phone" dc:"电话"`
	Gender           mCustomer.GenderType `json:"gender"`
	Birthday         int64                `json:"birthday"`
	Language         string               `json:"language"`
	Address          []mAddress.Address   `json:"address"`
	Tags             []string             `json:"tags"`
	Tax              CustomerFreeTax      `json:"tax"`
	DefaultAddressID uint                 `json:"default_address_id"`
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
type CustomerAddAddressReq struct {
	g.Meta     `path:"/customer/add/address" method:"post" summary:"添加地址" tags:"Customer"`
	CustomerID uint             `json:"customer_id" v:"required"`
	Address    mAddress.Address `json:"address" v:"required"`
}
type CustomerAddAddressRes struct {
}

// 更新收货地址
type CustomerUpdateAddressReq struct {
	g.Meta     `path:"/customer/update/address" method:"post" summary:"更新地址" tags:"Customer"`
	Address    mAddress.Address `json:"address" V:"required"`
	IsDefault  bool             `json:"is_default"`
	CustomerID uint             `json:"customer_id" v:"required"`
}
type CustomerUpdateAddressRes struct {
}

type CustomerSetTaxItem struct {
	ID          uint     `json:"id" dc:"单个地区id"`
	CountryCode string   `json:"country_code" dc:"国家代码"`
	Zones       []string `json:"zones" dc:"地区"`
}

// 设置免税地区
type CustomerSetTaxReq struct {
	g.Meta `path:"/customer/set/tax" method:"post" summary:"设置免税地区" tags:"Customer"`
	ID     uint                 `json:"id" v:"required" dc:"客户ID"`
	Areas  []CustomerSetTaxItem `json:"areas" dc:"免税地区"`
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

// 删除地址
type CustomerDeleteAddressReq struct {
	g.Meta    `path:"/customer/delete/address" method:"post" summary:"删除地址" tags:"Customer"`
	AddressID uint `json:"address_id" v:"required"`
}
type CustomerDeleteAddressRes struct {
}
