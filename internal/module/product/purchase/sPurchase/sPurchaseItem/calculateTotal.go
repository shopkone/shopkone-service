package sPurchaseItem

import (
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/utility/handle"
)

func (s *sPurchaseItem) CalculateTotal(item mPurchase.PurchaseItem, status mPurchase.PurchaseStatus) (total float64) {
	// 如果状态是草稿单，则计算全部
	if status == mPurchase.PurchaseStatusDraft {
		total = float64(item.Purchasing) * (item.Cost + item.Cost*item.TaxRate*0.01)
		return handle.RoundMoney(total)
	}
	// 否则计算已收货
	return handle.RoundMoney(float64(item.Received) * (item.Cost + item.Cost*item.TaxRate*0.01))
}
