package sDesign

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"shopkone-service/hack"
)

type sDesign struct {
	shopId uint
	host   string
}

func NewDesign(shopId uint) *sDesign {
	config, _ := hack.GetConfig()
	return &sDesign{shopId, config.Render.Host}
}

func (s *sDesign) GetClient() *gclient.Client {
	c := g.Client().SetHeader("shop_id", convertor.ToString(s.shopId)).ContentJson()
	return c.Prefix(s.host)
}
