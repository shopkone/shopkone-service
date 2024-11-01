package sSupplier

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sSupplier) Update(in vo.SupplierUpdateReq) error {
	// 获取供应商
	var info mProduct.Supplier
	if err := s.orm.Model(&info).Where("shop_id = ? AND id = ?", s.shopId, in.Id).
		Select("id", "address_id").First(&info).Error; err != nil {
		return err
	}

	// 更新地址
	in.Address.ID = info.AddressId
	return sAddress.NewAddress(s.orm, s.shopId).UpdateById(in.Address)
}
