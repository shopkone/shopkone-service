package iInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

type CreateInventoryIn struct {
	LocationId uint
	VariantId  uint
	Quantity   uint
	Id         uint
}

type UpdateByDiffIn struct {
	News        []mInventory.Inventory
	Olds        []mInventory.Inventory
	HandleEmail string
	UpdateType  mInventory.InventoryType
}

type CountByVariantIdsOut struct {
	VariantId uint
	Quantity  uint
}
