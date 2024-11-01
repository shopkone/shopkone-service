package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/shop/shop/mShop"
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
