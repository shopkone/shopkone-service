package mShop

import (
	"gorm.io/gorm"
)

type DefaultDomain struct {
	Domain string `json:"domain"`
	Open   bool   `json:"open"`
}

type ShopStatus int

const (
	SHOP_SERVEING    ShopStatus = iota + 1 // 服务中
	SHOP_EXPIRED                           // 过期
	SHOP_TRIAL                             // 试用
	SHOP_FROZEN                            // 冻结
	SHOP_WILL_EXPIRE                       // 即将过期
)

// CurrencyFormatType 货币格式
type CurrencyFormatType string

const (
	// 带有千位分隔符（逗号）和小数点的金额
	AmountWithCommaAndDecimal CurrencyFormatType = "123,456.78"

	// 仅带有千位分隔符但没有小数的金额
	AmountWithCommaNoDecimal CurrencyFormatType = "123,456"

	// 小数点为逗号的金额
	AmountWithCommaAsDecimal CurrencyFormatType = "123.456,78"

	// 无千位分隔符但有小数的金额
	AmountNoCommaWithDecimal CurrencyFormatType = "123.456"

	// 带有特殊字符（如单引号）的金额
	AmountWithSpecialCharacter CurrencyFormatType = "123’456.65"
)

// Shop 店铺
type Shop struct {
	gorm.Model
	Status               ShopStatus         `gorm:"index"`                      // 店铺状态
	StoreName            string             `gorm:"size:50;default:My Store"`   // 店铺名称
	WebsiteFaviconId     uint               `gorm:"index"`                      // 店铺图标id
	AddressId            uint               `gorm:"index"`                      // 地址id
	CustomerServiceEmail string             `gorm:"size:256"`                   // 客服邮箱
	StoreOwnerEmail      string             `gorm:"size:256"`                   // 店主邮箱
	StoreCurrency        string             `gorm:"size:3"`                     // 店铺货币
	CurrencyFormatting   CurrencyFormatType `gorm:"size:50;default:123,456.78"` // 货币格式
	TimeZone             string             `gorm:"size:50"`                    // 时区
	OrderIdPrefix        string             `gorm:"size:50"`                    // 订单id前缀
	OrderIdSuffix        string             `gorm:"size:50"`                    // 订单id后缀
	PasswordProtection   bool               `gorm:"default:true"`               // 密码保护
	Password             string             `gorm:"size:256"`                   // 密码保护密码
	PasswordMessage      string             `gorm:"size:500"`                   // 密码保护提示
	Uuid                 string             `gorm:"index"`                      // UUID
	Country              string             `gorm:"size:50"`                    // 国家
	TaxShipping          bool               `gorm:"default:false"`              // 对运费收的税包含在运费中
}
