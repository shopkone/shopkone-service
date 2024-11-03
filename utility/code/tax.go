package code

var (
	TaxZooeCodeRepeat    = taxError(1, "区域不能重复")
	TaxCollectionMust    = taxError(2, "必须选择一个商品集合")
	TaxCustomerZonesMust = taxError(3, "必须选择一个区域")
)
