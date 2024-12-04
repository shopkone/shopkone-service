package sDesign

import "shopkone-service/hack"

type sDesign struct {
	shopId uint
	host   string
}

func NewDesign(shopId uint) *sDesign {
	config, _ := hack.GetConfig()
	return &sDesign{shopId, config.Render.Host}
}
