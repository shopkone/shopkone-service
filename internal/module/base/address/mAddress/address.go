package mAddress

import "shopkone-service/internal/module/base/orm/mOrm"

type Phone struct {
	Country string `json:"country" gorm:"size:20"` // 国家
	Prefix  int    `json:"prefix" gorm:"size:20"`  // 国家区号
	Num     string `json:"num" gorm:"size:20"`     // 手机号
}

type Address struct {
	mOrm.Model
	LegalBusinessName string `json:"legal_business_name" gorm:"size:200"` // 商家名称
	Address1          string `json:"address1" gorm:"size:200"`            // 地址1
	Address2          string `json:"address2" gorm:"size:200"`            // 地址2
	City              string `json:"city" gorm:"size:200"`                // 城市
	Company           string `json:"company" gorm:"size:200"`             // 公司
	Country           string `json:"country" gorm:"size:200"`             // 国家
	FirstName         string `json:"first_name" gorm:"size:200"`          // 名
	LastName          string `json:"last_name" gorm:"size:200"`           // 姓
	Phone             Phone  `json:"phone" gorm:"serializer:json"`        // 电话
	PostalCode        string `json:"postal_code" gorm:"size:200"`         // 邮编
	Zone              string `json:"zone" gorm:"size:200"`                // 省份
}
