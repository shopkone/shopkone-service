package cmd

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"shopkone-service/internal/api"
	"shopkone-service/utility/middleware"
)

func registerAdminRoutes(s *ghttp.Server) {
	shopApi := api.NewShopApi()
	userApi := api.NewUserApi()

	baseApi := api.NewBaseApi()
	fileApi := api.NewFileApi()
	locationApi := api.NewLocationApi()

	productApi := api.NewProductApi()
	inventoryApi := api.NewInventoryApi()
	productCollectionApi := api.NewProductCollectionApi()
	purchaseApi := api.NewPurchaseApi()
	transferApi := api.NewTransferApi()
	shippingApi := api.NewShippingApi()
	localDeliveryApi := api.NewLocalDeliveryApi()
	inStorePickUpApi := api.NewInStorePickup()
	taxApi := api.NewTaxApi()
	marketApi := api.NewMarketApi()
	langaugesApi := api.NewLanguageApi()
	domainApi := api.NewDomainApi()

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.UserMiddleware)
		group.Bind(shopApi.List) // 获取店铺配置
	})

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthMiddleware)
		// 店铺
		group.Bind(shopApi.Info)                        // 获取店铺信息
		group.Bind(shopApi.General)                     // 获取店铺设置
		group.Bind(shopApi.UpdateGeneral)               // 更新店铺设置
		group.Bind(shopApi.TaxSwitchShipping)           // 对运费收的税包含在运费中
		group.Bind(shopApi.ShopTaxSwitchShippingUpdate) // 更新对运费收的税包含在运费中

		// 用户
		group.Bind(userApi.Info)       // 获取用户信息
		group.Bind(userApi.SetColumns) // 设置员工列
		group.Bind(userApi.GetColumns) // 获取用户列

		// 基础应用
		group.Bind(baseApi.Countries)    // 获取国家列表
		group.Bind(baseApi.PhonePrefix)  // 获取手机区号列表
		group.Bind(baseApi.TimezoneList) // 获取时区列表
		group.Bind(baseApi.CurrencyList) // 获取货币列表
		group.Bind(baseApi.CategoryList) // 获取商品分类列表
		group.Bind(baseApi.CarrierList)  // 获取物流公司列表
		group.Bind(baseApi.LanguageList) // 获取语言列表

		// 文件
		group.Bind(fileApi.GetUploadToken)         // 获取上传token
		group.Bind(fileApi.AddFile)                // 添加文件
		group.Bind(fileApi.FileList)               // 获取文件列表
		group.Bind(fileApi.FilesDelete)            // 删除文件
		group.Bind(fileApi.FileInfo)               // 获取文件信息
		group.Bind(fileApi.FileUpdateInfo)         // 更新文件信息
		group.Bind(fileApi.FileListByIds)          // 根据文件ID获取文件列表
		group.Bind(fileApi.UpdateGroupIdByFileIds) // 更新文件分组
		group.Bind(fileApi.FileGroupList)          // 获取文件分组列表
		group.Bind(fileApi.FileGroupAdd)           //  添加文件分组
		group.Bind(fileApi.FileGroupUpdate)        //  更新文件分组
		group.Bind(fileApi.FileGroupRemove)        //  删除文件分组

		// 地点
		group.Bind(locationApi.List)
		group.Bind(locationApi.LocationAdd)
		group.Bind(locationApi.LocationInfo)
		group.Bind(locationApi.LocationUpdate)
		group.Bind(locationApi.LocationDelete)
		group.Bind(locationApi.LocationExistInventory)
		group.Bind(locationApi.LocationSetDefault)
		group.Bind(locationApi.SetLocationOrder)

		// 商品
		group.Bind(productApi.Create)
		group.Bind(productApi.Info)
		group.Bind(productApi.List)
		group.Bind(productApi.Update)
		group.Bind(productApi.ListByIds)
		group.Bind(productApi.VariantsByIDs)
		group.Bind(productApi.CreateSupplier)
		group.Bind(productApi.SupplierList)
		group.Bind(productApi.UpdateSupplier)

		// 采购单
		group.Bind(purchaseApi.Create)
		group.Bind(purchaseApi.Update)
		group.Bind(purchaseApi.List)
		group.Bind(purchaseApi.Info)
		group.Bind(purchaseApi.MarkToOrdered)
		group.Bind(purchaseApi.Remove)
		group.Bind(purchaseApi.Adjust)
		group.Bind(purchaseApi.Close)

		// 库存
		group.Bind(inventoryApi.List)
		group.Bind(inventoryApi.Move)
		group.Bind(inventoryApi.HistoryList)
		group.Bind(inventoryApi.InventoryListUnByVariantIds)

		// 商品集合
		group.Bind(productCollectionApi.Create)
		group.Bind(productCollectionApi.List)
		group.Bind(productCollectionApi.Update)
		group.Bind(productCollectionApi.CollectionOptions)
		group.Bind(productCollectionApi.Info)

		// 库存转移
		group.Bind(transferApi.CreateTransfer)
		group.Bind(transferApi.List)
		group.Bind(transferApi.Info)
		group.Bind(transferApi.Mark)
		group.Bind(transferApi.Adjust)
		group.Bind(transferApi.Remove)
		group.Bind(transferApi.Update)

		// 物流
		group.Bind(shippingApi.Create)
		group.Bind(shippingApi.Update)
		group.Bind(shippingApi.Info)
		group.Bind(shippingApi.List)
		group.Bind(shippingApi.ZoneListByCountries)

		// 本地配送
		group.Bind(localDeliveryApi.List)
		group.Bind(localDeliveryApi.Info)
		group.Bind(localDeliveryApi.Update)

		// 到店自取
		group.Bind(inStorePickUpApi.List)
		group.Bind(inStorePickUpApi.Info)
		group.Bind(inStorePickUpApi.Update)

		// 税收
		group.Bind(taxApi.List)
		group.Bind(taxApi.Info)
		group.Bind(taxApi.TaxUpdate)
		group.Bind(taxApi.TaxCreate)
		group.Bind(taxApi.TaxRemove)
		group.Bind(taxApi.TaxActive)

		// 市场
		group.Bind(marketApi.Create)
		group.Bind(marketApi.List)
		group.Bind(marketApi.Info)
		group.Bind(marketApi.Update)
		group.Bind(marketApi.Options)
		group.Bind(marketApi.UpDomain)
		group.Bind(marketApi.UpdateLang)
		group.Bind(marketApi.UpdateLangByLangID)
		group.Bind(marketApi.UpdateProduct)
		group.Bind(marketApi.GetProduct)
		group.Bind(marketApi.Simple)

		// 语言
		group.Bind(langaugesApi.List)
		group.Bind(langaugesApi.Create)

		// 域名
		group.Bind(domainApi.List)
		group.Bind(domainApi.PreCheck)
		group.Bind(domainApi.ConnectCheck)
		group.Bind(domainApi.BlockCountryUpdate)
		group.Bind(domainApi.BlockCountryList)
	})
}
