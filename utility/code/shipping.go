package code

var (
	ProductMust            = shippingError(1, "商品不能为空")
	LocationMust           = shippingError(2, "发货地点不能为空")
	ZoneMust               = shippingError(3, "收货区域不能为空")
	ZoneFeeMust            = shippingError(4, "运费方案不能为空")
	ZoneCodeMust           = shippingError(5, "国家/地区不能为空")
	ZoneFeeConditionMust   = shippingError(6, "运费费用不能为空")
	GeneralShippingExist   = shippingError(7, "通用运费方案已存在")
	ShippingNameExist      = shippingError(8, "方案名称已存在")
	ShippingNameMust       = shippingError(9, "方案名称不能为空")
	ShippingZoneMust       = shippingError(10, "区域名称不能为空")
	ShippingZoneNameRepeat = shippingError(11, "区域名称重复")
	ShippingFeeNameRepeat  = shippingError(12, "运费方案名称重复")
	ShippingFeeNameMust    = shippingError(13, "运费方案名称不能为空")
	ShippingTypeMust       = shippingError(14, "请选择运费类型")
)
