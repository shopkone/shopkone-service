package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/utility/handle"
)

type BasePurchaseItem struct {
	Id         uint    `json:"id" dc:"变体ID"`
	Cost       float64 `json:"cost" v:"required" dc:"成本"`
	Purchasing int     `json:"purchasing" v:"required" dc:"采购数量"`
	SKU        string  `json:"sku" dc:"SKU"`
	TaxRate    float64 `json:"tax_rate" v:"required" dc:"税率"`
	VariantID  uint    `json:"variant_id" v:"required" dc:"变体ID"`
	Rejected   int     `json:"rejected"`
	Received   int     `json:"received"`
	Total      float64 `json:"total"`
	IsActive   bool    `json:"is_active"`
}

type PurchaseCreateReq struct {
	g.Meta           `path:"/purchase/create" method:"post" tags:"Purchase" summary:"创建采购单"`
	SupplierId       uint                           `json:"supplier_id" v:"required" dc:"供应商ID"`
	DestinationId    uint                           `json:"destination_id" v:"required" dc:"目的地ID"`
	CarrierId        *uint                          `json:"carrier_id" dc:"物流商ID"`
	DeliveryNumber   string                         `json:"delivery_number" dc:"物流商单号"`
	CurrencyCode     string                         `json:"currency_code" v:"required" dc:"货币代码"`
	Remarks          string                         `json:"remarks" dc:"备注"`
	EstimatedArrival int64                          `json:"estimated_arrival" dc:"预计送达时间"`
	PaymentTerms     mPurchase.PaymentTermsType     `json:"payment_terms" dc:"付款条款"`
	PurchaseItems    []BasePurchaseItem             `json:"purchase_items" v:"required" dc:"采购项"`
	Adjust           []mPurchase.PurchaseAdjustItem `json:"adjust" dc:"采购调整项"`
}
type PurchaseCreateRes struct {
	Id uint `json:"id" dc:"采购单ID"`
}

type PurchaseListReq struct {
	g.Meta `path:"/purchase/list" method:"post" tags:"Purchase" summary:"采购单列表"`
	handle.PageReq
}
type PurchaseListRes struct {
	OrderNumber      string                   `json:"order_number"`      // 订单编号
	SupplierId       uint                     `json:"supplier_id"`       // 供应商ID
	DestinationId    uint                     `json:"destination_id"`    // 目的地
	Status           mPurchase.PurchaseStatus `json:"status"`            // 采购状态
	Received         uint                     `json:"received"`          // 收货数量
	Rejected         uint                     `json:"rejected"`          // 拒绝数量
	Purchasing       int                      `json:"purchasing"`        // 采购数量
	Total            float64                  `json:"total"`             // 总价格
	EstimatedArrival int64                    `json:"estimated_arrival"` // 预计送达时间
	Id               uint                     `json:"id"`                //
}

type PurchaseInfoReq struct {
	g.Meta `path:"/purchase/info" method:"post" tags:"Purchase" summary:"采购单详情"`
	Id     uint `json:"id" v:"required" dc:"采购单ID"`
}
type PurchaseInfoRes struct {
	PurchaseCreateReq
	Id          uint                     `json:"id"`
	Status      mPurchase.PurchaseStatus `json:"status"`
	OrderNumber string                   `json:"order_number"`
	Rejected    int                      `json:"rejected"`
	Received    int                      `json:"received"`
	Purchasing  int                      `json:"purchasing"`
}

type PurchaseUpdateReq struct {
	g.Meta `path:"/purchase/update" method:"post" tags:"Purchase" summary:"更新采购单"`
	PurchaseCreateReq
	Id uint `json:"id" v:"required" dc:"采购单ID"`
}
type PurchaseUpdateRes struct {
}

type PurchaseMarkToOrderedReq struct {
	g.Meta `path:"/purchase/mark-to-ordered" method:"post" tags:"Purchase" summary:"标记为已下单"`
	Id     uint `json:"id" v:"required" dc:"采购单ID"`
}
type PurchaseMarkToOrderedRes struct {
}

type PurchaseRemoveReq struct {
	g.Meta `path:"/purchase/remove" method:"post" tags:"Purchase" summary:"删除采购单"`
	Id     uint `json:"id" v:"required" dc:"采购单ID"`
}
type PurchaseRemoveRes struct {
}

type PurchaseCloseReq struct {
	g.Meta `path:"/purchase/close" method:"post" tags:"Purchase" summary:"关闭采购单"`
	Id     uint `json:"id" v:"required" dc:"采购单ID"`
	Close  bool `json:"close" v:"required" dc:""`
}
type PurchaseCloseRes struct {
}

type PurchaseAdjustReceiveItem struct {
	Id            uint `json:"id" v:"required"`
	RejectedCount int  `json:"rejected_count"`
	ReceivedCount int  `json:"received_count"`
}

type PurchaseAdjustReceiveReq struct {
	g.Meta `path:"/purchase/adjust-receive" method:"post" tags:"Purchase" summary:"采购单收货调整"`
	Id     uint                        `json:"id" v:"required" dc:"采购单ID"`
	Items  []PurchaseAdjustReceiveItem `json:"items" v:"required"`
}
type PurchaseAdjustReceiveRes struct {
}
