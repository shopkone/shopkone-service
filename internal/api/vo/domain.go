package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/setting/domains/mDomains"
)

type DomainListReq struct {
	g.Meta `path:"/domain/list" method:"post" tags:"Domain" summary:"获取域名列表"`
	Status []mDomains.DomainStatus `json:"status"`
}
type DomainListRes struct {
	ID         uint                  `json:"id"`
	Status     mDomains.DomainStatus `json:"status"`
	IsMain     bool                  `json:"is_main"`
	IsShopKone bool                  `json:"is_shopkone"`
	Domain     string                `json:"domain"`
}

type DomainPreCheckReq struct {
	g.Meta `path:"/domain/pre" method:"post" tags:"Domain" summary:"检查域名是否可用"`
	Domain string `json:"domain" v:"required" dc:"域名"`
}

type DomainPreCheckRes struct {
	Type  string `json:"type"`
	Host  string `json:"host"`
	Value string `json:"value"`
}

type DomainConnectCheckReq struct {
	g.Meta `path:"/domain/connect/check" method:"post" tags:"Domain" summary:"检查域名是否可用"`
	Domain string `json:"domain" v:"required" dc:"域名"`
}
type DomainConnectCheckRes struct {
	IsConnect bool `json:"is_connect"`
}

type DomainBlockCountryUpdateReq struct {
	g.Meta `path:"/domain/block-country/update" method:"post" tags:"Domain" summary:"更新域名屏蔽国家"`
	Codes  []string `json:"codes"`
}
type DomainBlockCountryUpdateRes struct {
}

type DomainBlockCountryListReq struct {
	g.Meta `path:"/domain/block-country/list" method:"post" tags:"Domain" summary:"获取域名屏蔽国家列表"`
}
type DomainBlockCountryListRes struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type DomainBlackIpUpdateReq struct {
	g.Meta `path:"/domain/black-ip/update" method:"post" tags:"Domain" summary:"更新域名黑名单IP"`
	Ips    []string             `json:"ips"`
	Type   mDomains.BlackIpType `json:"type" v:"required"`
}
type DomainBlackIpUpdateRes struct {
}

type DomainBlackIpListReq struct {
	g.Meta `path:"/domain/black-ip/list" method:"post" tags:"Domain" summary:"获取域名黑名单IP列表"`
}
type DomainBlackIpListRes struct {
	Ip   string               `json:"ip"`
	Type mDomains.BlackIpType `json:"type"`
}

type DomainBlackIpRemoveReq struct {
	g.Meta `path:"/domain/black-ip/remove" method:"post" tags:"Domain" summary:"移除域名黑名单IP"`
	Ips    []string             `json:"ips" v:"required"`
	Type   mDomains.BlackIpType `json:"type" v:"required"`
}
type DomainBlackIpRemoveRes struct {
}
