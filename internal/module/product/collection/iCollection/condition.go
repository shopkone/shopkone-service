package iCollection

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/collection/mCollection"
)

type CreateListIn struct {
	Conditions   []vo.BaseCondition
	CollectionId uint
	MatchMode    mCollection.CollectionMatchMode
}

type MatchProductsIn struct {
	Conditions   []vo.BaseCondition
	CollectionId uint
	MatchMode    mCollection.CollectionMatchMode
}

type UpdateListIn struct {
	Conditions   []vo.BaseCondition
	CollectionId uint
	MatchMode    mCollection.CollectionMatchMode
}

type MatchOneOutType string

const (
	MatchOneOutTypeProduct     MatchOneOutType = "product"
	MatchOneOutTypeVariant     MatchOneOutType = "variant"
	MatchOneOutTypeInventory   MatchOneOutType = "inventory"
	MatchOneOutTypeVariantName MatchOneOutType = "variantName"
)

type MatchOneOut struct {
	Type      MatchOneOutType
	Query     string
	Value     string
	Valid     bool
	NeedValue bool
}

type MatchActionOut struct {
	Query     string
	Value     string
	Valid     bool
	NeedValue bool
}
