package code

var (
	TaxZooeCodeRepeat   = taxError(1, "区域不能重复")
	TaxCollectionMust   = taxError(2, "必须选择一个商品集合")
	TaxZoneNotFound     = taxError(3, "该国家无该区域")
	TaxRateLessThanZero = taxError(4, "税率不能小于0")
	TaxZonesMust        = taxError(5, "必须添加一个区域")
	TaxCollectionRepeat = taxError(6, "商品系列重复")
)
