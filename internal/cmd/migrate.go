package cmd

import (
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/seo/mSeo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/online/mav/mNav"
	"shopkone-service/internal/module/order/order/mOrder"
	"shopkone-service/internal/module/product/collection/mCollection"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/setting/domains/mDomains"
	"shopkone-service/internal/module/setting/file/mFile"
	"shopkone-service/internal/module/setting/language/mLanguage"
	"shopkone-service/internal/module/setting/location/mLocation"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/internal/module/shop/policy/mPolicy"
	"shopkone-service/internal/module/shop/shop/mShop"
	"shopkone-service/internal/module/shop/staff/mStaff"
	"shopkone-service/internal/module/shop/transaction/mTransaction"
	"shopkone-service/internal/module/shop/user/mUser"
)

type Migrate []interface{}

var migrate = Migrate{
	// 店铺
	mUser.User{},
	mUser.UserLoginRecord{},
	mUser.UserColumn{},
	mStaff.Staff{},
	mShop.Shop{},

	mSeo.Seo{},
	mFile.File{},
	mFile.FileGroup{},
	mLocation.Location{},
	mInventory.Inventory{},
	mInventory.InventoryChange{},
	mInventory.LogisticsProvider{},

	mTransaction.Transaction{},
	mPolicy.Policy{},

	// 商品
	mProduct.Product{},
	mProduct.ProductFiles{},
	mProduct.ProductOption{},
	mProduct.Supplier{}, // 供应商

	// 集合
	mCollection.CollectionProduct{},
	mCollection.ProductCondition{},
	mCollection.ProductCollection{},

	// 变体
	mProduct.Variant{},
	mProduct.VariantNameHandler{},

	// 采购单
	mPurchase.Purchase{},
	mPurchase.PurchaseItem{},

	mProduct.Vendor{},
	mOrder.Order{},
	mOrder.OrderAfterSale{},
	mOrder.OrderVariant{},
	mShipping.ShippingNote{},
	mShipping.ShippingNoteItem{},
	mAddress.Address{},

	// 库存转移
	mTransfer.Transfer{},
	mTransfer.TransferItem{},

	// 物流
	mShipping.Shipping{},
	mShipping.ShippingProduct{},
	mShipping.ShippingLocation{},
	mShipping.ShippingZone{},
	mShipping.ShippingZoneCode{},
	mShipping.ShippingZoneFee{},
	mShipping.ShippingZonFeeCondition{},

	// 本地配送
	mLocalDelivery.LocalDelivery{},
	mLocalDelivery.LocalDeliveryArea{},
	mLocalDelivery.LocalDeliveryFee{},

	// 到店自提
	mInStorePickup.InStorePickup{},
	mInStorePickup.InStorePickupBusinessHours{},

	// 税率
	mTax.Tax{},
	mTax.TaxZone{},
	mTax.CustomerTax{},
	mTax.CustomerTaxZone{},

	// 市场
	mMarket.Market{},
	mMarket.MarketCountry{},
	mMarket.MarketProduct{},
	mMarket.MarketPrice{},

	// 语言
	mLanguage.Language{},

	// 域名
	mDomains.Domain{},
	mDomains.DomainBlackIp{},
	mDomains.DomainBlockCountry{},

	// 客户
	mCustomer.Customer{},
	mCustomer.Cart{},
	mCustomer.CustomerAddress{},
	mCustomer.CustomerNoTaxArea{},

	// 在线商城
	mNav.Nav{},
}
