package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/order/order/mOrder"
)

type OrderCalculatePreDiscount struct {
	ID    uint                     `json:"id"`
	Type  mOrder.OrderDiscountType `json:"type"`
	Value float32                  `json:"value"`
}

type OrderCalculatePreVariantItem struct {
	VariantID uint                      `json:"variant_id" v:"required"`
	Quantity  uint                      `json:"quantity" v:"required"`
	Discount  OrderCalculatePreDiscount `json:"discount"`
}

type OrderCalculatePreReq struct {
	g.Meta       `path:"/order/calculate-pre" method:"post" summary:"计算订单价格" tags:"Order"`
	VariantItems []OrderCalculatePreVariantItem `json:"variant_items" v:"required"`
}
