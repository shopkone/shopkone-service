package sShop

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/util/guid"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/internal/module/setting/language/sLanguage"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	"shopkone-service/internal/module/setting/tax/sTax/sTax"
	"shopkone-service/internal/module/shop/shop/iShop"
	"shopkone-service/internal/module/shop/shop/mShop"
	"shopkone-service/internal/module/shop/staff/mStaff"
	"shopkone-service/internal/module/shop/staff/sStaff"
	"shopkone-service/utility/handle"
)

type IShop interface {
	// CreateTrial 创建试用店铺
	CreateTrial(in iShop.CreateTrialIn) (shopId uint, err error)
	// ShopListByUserId 根据用户id获取店铺列表
	ShopListByUserId(userId uint) (out []mShop.Shop, err error)
	// UpdateShopGeneral 更新店铺信息
	UpdateShopGeneral(shopId uint, in vo.ShopUpdateGeneralReq) (err error)
	// RemoveCoverByFileIds 删除封面图片
	RemoveCoverByFileIds(fileIds []uint, shopId uint) (err error)
}

type sShop struct {
	orm *gorm.DB
}

func NewShop(orm *gorm.DB) *sShop {
	return &sShop{orm: orm}
}

func (s *sShop) CreateTrial(in iShop.CreateTrialIn) (shopId uint, err error) {
	// 获取国家
	country, err := sResource.NewCountry().CountryByCode(in.Country)
	if err != nil {
		return shopId, err
	}
	timezone, err := sResource.NewTimezone().TimezoneByCountry(in.Country)
	if err != nil {
		return shopId, err
	}
	shop := mShop.Shop{
		Model:                gorm.Model{},
		StoreName:            "My Store",
		CustomerServiceEmail: in.Email,
		StoreOwnerEmail:      in.Email,
		Status:               mShop.SHOP_TRIAL,
		PasswordProtection:   true,
		Password:             guid.S()[0:6],
		PasswordMessage:      "This store is password protected. Use the password to enter the store",
		Uuid:                 handle.GenUid(),
		OrderIdPrefix:        "#",
		StoreCurrency:        country.Currencies[0],
		TimeZone:             timezone,
		Country:              country.Code,
	}
	// 创建店铺
	if err = s.orm.Create(&shop).Error; err != nil {
		return shopId, err
	}
	// 创建默认税费
	if err = sTax.NewTax(s.orm, shop.ID).TaxCreate([]string{shop.Country}); err != nil {
		return 0, err
	}
	// 初始化员工
	if _, err = sStaff.NewStaff(s.orm, shop.ID).CreateOnJobStaff(in.UserId, "admin", true); err != nil {
		return 0, err
	}
	// 初始化地址
	address, err := s.CreateInitAddress(in.Country, in.Zone, shop.ID)
	if err != nil {
		return 0, err
	}
	// 创建默认语言
	if _, err = sLanguage.NewLanguage(s.orm, shop.ID).LanguageCreate([]string{"en"}, true); err != nil {
		return 0, err
	}
	// 创建主要市场
	marketCreateIn := vo.MarketCreateReq{Name: shop.Country, CountryCodes: []string{shop.Country}, IsMain: true, Force: true}
	if _, err = sMarket.NewMarket(s.orm, shop.ID).MarketCreate(marketCreateIn); err != nil {
		return 0, err
	}
	// 初始化地点(这里的address与上面的address是两个不同的addrsss，只是初始化的时候使用为了方便使用了同一个)
	locationCreateIn := vo.LocationAddReq{
		Name:    "Default",
		Address: address,
	}
	if _, err = sLocation.NewLocation(s.orm, shop.ID).Create(locationCreateIn, shop.TimeZone); err != nil {
		return 0, err
	}
	// 更新缓存
	if err = NewShopCache().UpdateShopCache(shop.ID, s.orm); err != nil {
		return 0, err
	}
	return shop.ID, nil
}

func (s *sShop) ShopListByUserId(userId uint) (out []mShop.Shop, err error) {
	// 获取员工信息
	staffs, err := sStaff.NewStaff(s.orm, 0).StaffsByUserId(userId)
	if err != nil {
		return out, err
	}
	// 获取店铺id
	shopIds := slice.Map(staffs, func(_ int, item mStaff.Staff) uint {
		return item.ShopId
	})
	// 根据店铺id获取店铺信息
	if err = s.orm.Where("id IN ?", shopIds).
		Select("status", "website_favicon_id", "uuid", "store_name").Find(&out).Error; err != nil {
		return out, err
	}
	return out, nil
}

func (s *sShop) CreateInitAddress(country, zone string, shopId uint) (address mAddress.Address, err error) {
	// 创建地址
	address = mAddress.Address{Country: country, Zone: zone}
	addressId, err := sAddress.NewAddress(s.orm, shopId).CreateAddress(address)
	if err != nil {
		return address, err
	}
	// 更新店铺地址id
	if err = s.orm.Model(&mShop.Shop{}).Where("id = ?", shopId).Update("address_id", addressId).Error; err != nil {
		return address, err
	}
	return address, err
}

func (s *sShop) UpdateShopGeneral(shopId uint, in vo.ShopUpdateGeneralReq, addressId uint) (err error) {
	// 更新地址信息
	in.Address.ID = addressId
	if err = sAddress.NewAddress(s.orm, shopId).UpdateById(in.Address); err != nil {
		return err
	}
	// 更新店铺信息
	data := mShop.Shop{}
	data.StoreName = in.StoreName
	data.StoreOwnerEmail = in.StoreOwnerEmail
	data.CustomerServiceEmail = in.CustomerServiceEmail
	data.WebsiteFaviconId = in.WebsiteFaviconId
	data.StoreCurrency = in.StoreCurrency
	data.CurrencyFormatting = in.CurrencyFormatting
	data.TimeZone = in.Timezone
	data.Password = in.Password
	data.PasswordMessage = in.PasswordMessage
	data.PasswordProtection = in.PasswordProtection
	data.OrderIdPrefix = in.OrderIdPrefix
	data.OrderIdSuffix = in.OrderIdSuffix
	if err = s.orm.Model(&mShop.Shop{}).Where("id = ?", shopId).
		Select(
			"store_name",
			"store_owner_email",
			"customer_service_email",
			"website_favicon_id",
			"store_currency",
			"currency_formatting",
			"time_zone",
			"password",
			"password_message",
			"password_protection",
			"order_id_prefix",
			"order_id_suffix",
		).
		Updates(data).Error; err != nil {
		return err
	}
	// 更新缓存
	return NewShopCache().UpdateShopCache(shopId, s.orm)
}

func (s *sShop) RemoveCoverByFileIds(fileIds []uint, shopId uint) (err error) {
	var data mShop.Shop
	query := s.orm.Model(&data).Where("id = ? AND website_favicon_id IN (?)", shopId, fileIds)
	if err = query.Update("website_favicon_id", nil).Error; err != nil {
		return err
	}
	// 更新缓存
	return NewShopCache().UpdateShopCache(shopId, s.orm)
}
