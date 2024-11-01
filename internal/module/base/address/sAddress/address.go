package sAddress

import (
	"gorm.io/gorm"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/utility/code"
)

type IAddress interface {
	// CreateAddress 创建地址
	CreateAddress(address mAddress.Address) (addressId uint, err error)
	// 获取地址
	GetAddress(addressId uint) (address mAddress.Address, err error)
	// UpdateById 更新地址
	UpdateById(address mAddress.Address) (err error)
	// ListByIds 批量获取地址
	ListByIds(addressIds []uint) ([]mAddress.Address, error)
}

type sAddress struct {
	orm    *gorm.DB
	shopId uint
}

func NewAddress(orm *gorm.DB, shopId uint) *sAddress {
	return &sAddress{orm: orm, shopId: shopId}
}

func (s *sAddress) CreateAddress(address mAddress.Address) (addressId uint, err error) {
	sCountry := sResource.NewCountry()
	// 获取国家信息
	country, err := sCountry.CountryByCode(address.Country)
	if err != nil {
		return 0, err
	}
	// 获取省份信息
	zone, err := sCountry.ZoneByCode(address.Zone)
	if err != nil {
		return 0, err
	}
	// 如果没有设置省份,且该国家下有省份，则默认设置为第一个省份
	if zone.Name == "" && len(country.Zones) > 0 {
		address.Zone = country.Zones[0].Code
	}
	// 如果没有设置电话号码区域，则根据国家取设置
	if address.Phone.Country == "" {
		address.Phone.Prefix = country.PhoneNumberPrefix
		address.Phone.Country = country.Code
	}
	// 创建地址
	address.ShopId = s.shopId
	address.ID = 0
	if err = s.orm.Create(&address).Error; err != nil {
		return 0, err
	}
	// 获取地址id
	return address.ID, nil
}

func (s *sAddress) GetAddress(addressId uint) (address mAddress.Address, err error) {
	err = s.orm.Model(&address).
		Omit("created_at", "deleted_at", "updated_at").
		Where("shop_id = ? and id = ?", s.shopId, addressId).
		Find(&address).Error
	address.CreatedAt = nil
	address.UpdatedAt = nil
	if err != nil {
		return address, err
	}
	return address, nil
}

func (s *sAddress) UpdateById(address mAddress.Address) (err error) {
	sCountry := sResource.NewCountry()
	// 获取国家信息
	country, err := sCountry.CountryByCode(address.Country)
	if err != nil {
		return err
	}
	// 获取省份信息
	zone, err := sCountry.ZoneByCode(address.Zone)
	if err != nil {
		return err
	}
	// 如果没有设置省份,且该国家下有省份，则默认设置为第一个省份
	if zone.Name == "" && len(country.Zones) > 0 {
		address.Zone = country.Zones[0].Code
	}
	// 如果没有设置电话号码区域，则根据国家取设置
	if address.Phone.Country == "" {
		address.Phone.Prefix = country.PhoneNumberPrefix
		address.Phone.Country = country.Code
	}
	if address.ID == 0 {
		return code.ErrAddressId
	}
	err = s.orm.Model(&mAddress.Address{}).
		Where("shop_id = ? AND id = ?", s.shopId, address.ID).
		Select(
			"address1",
			"address2",
			"city",
			"company",
			"country",
			"first_name",
			"last_name",
			"phone",
			"postal_code",
			"zone",
			"legal_business_name",
		).
		Updates(&address).Error
	return err
}

func (s *sAddress) ListByIds(addressIds []uint) ([]mAddress.Address, error) {
	var addresses []mAddress.Address
	return addresses, s.orm.Model(&addresses).
		Omit("updated_at", "deleted_at", "created_at").
		Where("shop_id = ? AND id in (?)", s.shopId, addressIds).
		Find(&addresses).Error
}
