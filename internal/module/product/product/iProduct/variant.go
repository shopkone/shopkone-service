package iProduct

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/mInventory"
)

type VariantUpdateIn struct {
	List              []vo.BaseVariantWithId
	ProductId         uint
	TrackInventory    bool
	HandleEmail       string
	EnableLocationIds []uint
}

type VariantCreateIn struct {
	List              []vo.BaseVariant
	ProductId         uint
	InventoryTracking bool
	Type              mInventory.InventoryType
	HandleEmail       string
	EnableLocationIds []uint
}

type VariantListByIdOut struct {
	Id        uint
	Name      string
	Image     string
	Sku       string
	ProductId uint
	IsDeleted bool
	Price     uint32
}
