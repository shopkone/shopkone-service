package sShop

import (
	"gorm.io/gorm"
	"shopkone-service/internal/module/base/cache/sCache"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	"shopkone-service/internal/module/shop/shop/mShop"
)

type IShopCache interface {
	// UpdateShopCache 更新店铺缓存
	UpdateShopCache(shopId uint, orm *gorm.DB) error
	// GetShopCache 获取店铺缓存
	GetShopCache(uuid string) (shop mShop.Shop, err error)
}

type sShopCache struct{}

func NewShopCache() *sShopCache {
	return &sShopCache{}
}

func (s *sShopCache) UpdateShopCache(shopId uint, orm *gorm.DB) error {
	// 更新店铺缓存
	shop := mShop.Shop{}
	query := orm.Model(&shop).Where("id = ?", shopId)
	if err := query.First(&shop).Error; err != nil {
		return err
	}
	key := sCache.SHOP_PREFIX_KEY + shop.Uuid
	year100 := 60 * 24 * 365 * 100 //100年
	if err := sCache.NewShopCache().Set(key, shop, uint(year100)); err != nil {
		return err
	}
	// 更新主要市场货币
	mps := sMarket.NewMarket(orm, shop.ID)
	if err := mps.PriceUpdateMainCurrency(shop.StoreCurrency); err != nil {
		return err
	}
	return nil
}

func (s *sShopCache) GetShopCache(uuid string) (shop mShop.Shop, err error) {
	key := sCache.SHOP_PREFIX_KEY + uuid
	err = sCache.NewShopCache().Get(key, &shop)
	return shop, err
}
