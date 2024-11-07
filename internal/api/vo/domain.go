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
