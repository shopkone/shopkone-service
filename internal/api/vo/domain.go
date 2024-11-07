package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/setting/domains/mDomains"
)

type DomainListReq struct {
	g.Meta `path:"/domain/list" method:"post" tags:"Domain" summary:"获取域名列表"`
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
