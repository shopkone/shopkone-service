package sSupplier

import (
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sSupplier) Create(address mAddress.Address) (id uint, err error) {
	// 创建地址
	addressId, err := sAddress.NewAddress(s.orm, s.shopId).CreateAddress(address)
	if err != nil {
		return 0, err
	}

	// 创建供应商
	data := mProduct.Supplier{}
	data.ShopId = s.shopId
	data.AddressId = addressId
	return data.ID, s.orm.Model(&data).Create(&data).Error
}
