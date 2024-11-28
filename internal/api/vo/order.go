package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/order/order/mOrder"
)

// 运费
type OrderPreBaseShippingFee struct {
	Free          bool    `json:"free"`
	Customer      bool    `json:"customer"`
	Name          string  `json:"name"`
	Price         float32 `json:"price"`
	ShippingFeeID uint    `json:"shipping_fee_id"`
}

// 优惠
type OrderPreBaseDiscount struct {
	ID    uint                     `json:"id"`
	Price float32                  `json:"value"`
	Type  mOrder.OrderDiscountType `json:"type"`
	Note  string                   `json:"note"`
}

// 商品项
type OrderPreBaseVariantItem struct {
	VariantID uint                 `json:"variant_id" v:"required"`
	Quantity  uint                 `json:"quantity" v:"required"`
	Discount  OrderPreBaseDiscount `json:"discount"`
}

// 返回可用运费方案
type BasePreShippingFeePlan struct {
	Name  string  `json:"name"`
	ID    uint    `json:"id"`
	Price float32 `json:"price"`
}

// 返回税务详情
type BasePreTaxDetail struct {
	Name  string  `json:"name"`
	Rate  float64 `json:"rate"`
	Price float32 `json:"price"`
}

// 预计算订单价格
type OrderCalPreReq struct {
	g.Meta       `path:"/order/calculate-pre" method:"post" summary:"计算订单价格" tags:"Order"`
	VariantItems []OrderPreBaseVariantItem `json:"variant_items" v:"required"`
	Discount     OrderPreBaseDiscount      `json:"discount"`
	Address      mAddress.Address          `json:"address"`
	CustomerID   uint                      `json:"customer_id"`
	ShippingFee  OrderPreBaseShippingFee   `json:"shipping_fee"`
}

type OrderCalPreRes struct {
	CostPrice        float32                  `json:"cost_price"`
	DiscountPrice    float32                  `json:"discount_price"`
	SumPrice         float32                  `json:"sum_price"`
	Total            float32                  `json:"total"`
	ShippingFeePlans []BasePreShippingFeePlan `json:"shipping_fee_plans"`
	Taxes            []BasePreTaxDetail       `json:"taxes"`
	ShippingFee      OrderPreBaseShippingFee  `json:"shipping_fee"`
}
