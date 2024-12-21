package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/shop/shop/mShop"
	"shopkone-service/internal/module/shop/transaction/mTransaction"
)

// 获取店铺详情
type ShopInfoReq struct {
	g.Meta `path:"/shop/info" method:"post" summary:"Shop Info" tags:"Shop"`
}
type ShopInfoRes struct {
	Uuid           string           `json:"uuid"`
	Status         mShop.ShopStatus `json:"status"`
	StoreName      string           `json:"store_name"`
	StoreCurrency  string           `json:"store_currency"`
	TimeZone       string           `json:"time_zone"`
	WebsiteFavicon string           `json:"website_favicon"`
	Country        string           `json:"country"`
}

// ShopListReq 获取店铺列表
type ShopListReq struct {
	g.Meta `path:"/shop/list" method:"post" summary:"Shop List" tags:"Shop"`
}
type ShopListRes struct {
	Uuid           string           `json:"uuid"`
	Status         mShop.ShopStatus `json:"status"`
	StoreName      string           `json:"store_name"`
	WebsiteFavicon string           `json:"website_favicon"`
}

// ShopGeneralReq 获取店铺基础设置
type ShopGeneralReq struct {
	g.Meta `path:"/shop/general" method:"post" summary:"Shop General" tags:"Shop"`
}
type ShopGeneralRes struct {
	StoreName            string                   `json:"store_name"`             // 店铺名称
	StoreOwnerEmail      string                   `json:"store_owner_email"`      // 店铺管理员邮箱
	CustomerServiceEmail string                   `json:"customer_service_email"` // 客服邮箱
	WebsiteFaviconId     uint                     `json:"website_favicon_id"`     // 网站图标
	Address              mAddress.Address         `json:"address"`                // 店铺地址
	StoreCurrency        string                   `json:"store_currency"`         // 货币
	CurrencyFormatting   mShop.CurrencyFormatType `json:"currency_formatting"`    // 货币格式
	Timezone             string                   `json:"timezone"`               // 时区
	Password             string                   `json:"password"`               // 密码
	PasswordProtection   bool                     `json:"password_protection"`    // 密码保护
	PasswordMessage      string                   `json:"password_message"`       // 密码提示
	OrderIdPrefix        string                   `json:"order_id_prefix"`        // 订单编号前缀
	OrderIdSuffix        string                   `json:"order_id_suffix"`        // 订单编号后缀
}

// ShopUpdateGeneralReq 更新店铺基础设置
type ShopUpdateGeneralReq struct {
	g.Meta               `path:"/shop/update/general" method:"post" summary:"Update Shop General" tags:"Shop"`
	StoreName            string                   `json:"store_name"`             // 店铺名称
	StoreOwnerEmail      string                   `json:"store_owner_email"`      // 店铺管理员邮箱
	CustomerServiceEmail string                   `json:"customer_service_email"` // 客服邮箱
	WebsiteFaviconId     uint                     `json:"website_favicon_id"`     // 网站图标
	Address              mAddress.Address         `json:"address"`                // 店铺地址
	StoreCurrency        string                   `json:"store_currency"`         // 货币
	CurrencyFormatting   mShop.CurrencyFormatType `json:"currency_formatting"`    // 货币格式
	Timezone             string                   `json:"timezone"`               // 时区
	Password             string                   `json:"password"`               // 密码
	PasswordProtection   bool                     `json:"password_protection"`    // 密码保护
	PasswordMessage      string                   `json:"password_message"`       // 密码提示
	OrderIdPrefix        string                   `json:"order_id_prefix"`        // 订单编号前缀
	OrderIdSuffix        string                   `json:"order_id_suffix"`        // 订单编号后缀
}

type ShopUpdateGeneralRes struct {
}

type ShopTaxSwitchShippingReq struct {
	g.Meta `path:"/shop/tax/shipping/switch" method:"post" summary:"Shop Tax Switch" tags:"Shop"`
}
type ShopTaxSwitchShippingRes struct {
	TaxShipping bool `json:"tax_shipping"` // 对运费收的税包含在运费中
}

// 更新运费税
type ShopTaxSwitchShippingUpdateReq struct {
	g.Meta      `path:"/shop/tax/shipping/update" method:"post" summary:"Update Shop Tax Switch" tags:"Shop"`
	TaxShipping bool `json:"tax_shipping"` // 对运费收的税包含在运费中
}
type ShopTaxSwitchShippingUpdateRes struct {
}

// 获取店铺的shopId
type ShopIdReq struct {
	g.Meta `path:"/shop/id" method:"post" summary:"Shop Id By Uuid" tags:"Shop"`
}
type ShopIdRes struct {
	ShopId uint `json:"shop_id"`
}

// 交易设置
type ShopUpdateTransactionReq struct {
	g.Meta                      `path:"/shop/update/transaction" method:"post" summary:"Shop Transaction" tags:"Shop"`
	TargetType                  mTransaction.TransactionTargetType `json:"target_type" v:"required"`            // 目标客户
	ReduceTime                  mTransaction.TransactionReduceTime `json:"reduce_time" v:"required"`            // 库存扣减时机
	IsForceCheckProduct         bool                               `json:"is_force_check_product" v:"required"` // 是否强制检测商品库存
	IsAutoFinish                bool                               `json:"is_auto_finish" v:"required"`         // 是否开启自动收货
	AutoFinishDay               uint                               `json:"auto_finish_day"`                     // 自动收货时间
	OrderAutoCancel             mTransaction.OrderAutoCancelType   `json:"order_auto_cancel" v:"required"`      // 订单自动取消时间
	OrderAutoCancelCustomerHour uint                               `json:"order_auto_cancel_customer_hour"`     // 自定义时间
}
type ShopUpdateTransactionRes struct {
}

// 交易详情z
type ShopTransactionInfoReq struct {
	g.Meta `path:"/shop/transaction/info" method:"post" summary:"Shop Transaction Info" tags:"Shop"`
}
type ShopTransactionInfoRes struct {
	TargetType                  mTransaction.TransactionTargetType `json:"target_type"`                     // 目标客户
	ReduceTime                  mTransaction.TransactionReduceTime `json:"reduce_time"`                     // 库存扣减时机
	IsForceCheckProduct         bool                               `json:"is_force_check_product"`          // 是否强制检测商品库存
	IsAutoFinish                bool                               `json:"is_auto_finish"`                  // 是否开启自动收货
	AutoFinishDay               uint                               `json:"auto_finish_day"`                 // 自动收货时间
	OrderAutoCancel             mTransaction.OrderAutoCancelType   `json:"order_auto_cancel"`               // 订单自动取消时间
	OrderAutoCancelCustomerHour uint                               `json:"order_auto_cancel_customer_hour"` // 自定义时间
}
