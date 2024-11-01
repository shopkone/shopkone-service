package sSupplier

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sSupplier) List() (res []vo.SupplierListRes, err error) {
	// 查询供应商
	var list []mProduct.Supplier
	query := s.orm.Model(&list).Where("shop_id = ?", s.shopId)
	err = query.Select("id", "address_id").Find(&list).Error
	if err != nil {
		return
	}

	// 查询地址
	addressIds := slice.Map(list, func(index int, item mProduct.Supplier) uint {
		return item.AddressId
	})
	address, err := sAddress.NewAddress(s.orm, s.shopId).ListByIds(addressIds)
	if err != nil {
		return
	}

	res = slice.Map(list, func(index int, item mProduct.Supplier) vo.SupplierListRes {
		find, ok := slice.FindBy(address, func(index int, i mAddress.Address) bool {
			return i.ID == item.AddressId
		})
		i := vo.SupplierListRes{}
		i.Id = item.ID
		if ok {
			i.Address = find
		}
		return i
	})

	return res, err
}
