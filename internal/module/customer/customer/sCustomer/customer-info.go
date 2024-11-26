package sCustomer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
	"shopkone-service/internal/module/customer/customer/sCustomer/sCustomerAddress"
)

func (s *sCustomer) Info(id uint) (res vo.CustomerInfoRes, err error) {
	var customer mCustomer.Customer
	if err = s.orm.Model(&customer).Where("id = ?", id).
		Omit("created", "updated", "deleted_at").First(&customer).Error; err != nil {
		return res, err
	}
	res.ID = customer.ID
	res.FirstName = customer.FirstName
	res.LastName = customer.LastName
	res.Email = customer.Email
	res.Note = customer.Note
	res.Gender = customer.Gender
	res.Language = customer.Language
	res.Tax = vo.CustomerFreeTax{
		Free: customer.TaxFreeActive,
		All:  customer.TaxFreeAll,
	}
	if customer.Birthday != nil {
		res.Birthday = customer.Birthday.UnixMilli()
	}
	if customer.Phone != "" {
		if res.Phone, err = ParsePhoneStr(customer.Phone); err != nil {
			return vo.CustomerInfoRes{}, err
		}
	}
	res.Tags = customer.Tags
	// 获取地址
	address, err := sCustomerAddress.NewCustomerAddress(s.shopId, s.orm).List([]uint{customer.ID})
	if err != nil {
		return vo.CustomerInfoRes{}, err
	}
	if len(address) > 0 {
		res.Address = address[0].Address
		res.DefaultAddressID = address[0].DefaultAddressID
	}
	// 获取免税地区
	var taxAreas []mCustomer.CustomerNoTaxArea
	if err = s.orm.Model(&taxAreas).Where("customer_id = ?", customer.ID).
		Omit("created_at", "updated_at", "deleted_at").
		Find(&taxAreas).Error; err != nil {
		return res, err
	}
	res.Tax.Areas = slice.Map(taxAreas, func(index int, item mCustomer.CustomerNoTaxArea) vo.CustomerSetTaxItem {
		return vo.CustomerSetTaxItem{
			ID:          item.ID,
			CountryCode: item.CountryCode,
			Zones:       item.Zones,
		}
	})
	return res, err
}
