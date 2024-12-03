package sDesign

type sDesign struct {
	shopId uint
}

func NewDesign(shopId uint) *sDesign {
	return &sDesign{shopId: shopId}
}
