package sCustomer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

func (s *sCustomer) CustomerUpdateTax(in *vo.CustomerSetTaxReq) (err error) {
	// 更新客户税设置地区
	var customer mCustomer.Customer
	customer.TaxFreeAll = in.All
	customer.TaxFreeActive = in.Free
	if err = s.orm.Model(&customer).Where("id = ?", in.ID).
		Select("tax_free_active", "tax_free_all").
		Updates(&customer).Error; err != nil {
		return err
	}
	// 获取原有的免税地区
	var oldTax []mCustomer.CustomerNoTaxArea
	if err = s.orm.Model(&oldTax).Where("customer_id = ?", in.ID).
		Omit("created_at", "updated_at", "deleted_at").
		Find(&oldTax).Error; err != nil {
		return err
	}
	// 组装新的免税地区
	newTax := slice.Map(in.Areas, func(index int, item vo.CustomerSetTaxItem) mCustomer.CustomerNoTaxArea {
		i := mCustomer.CustomerNoTaxArea{}
		i.CountryCode = item.CountryCode
		i.CustomerID = in.ID
		i.Zones = item.Zones
		return i
	})
	// 删除
	removes := slice.Filter(oldTax, func(index int, area mCustomer.CustomerNoTaxArea) bool {
		_, ok := slice.FindBy(newTax, func(index int, item mCustomer.CustomerNoTaxArea) bool {
			return item.CountryCode == area.CountryCode
		})
		return !ok
	})
	if len(removes) > 0 {
		if err = s.orm.Delete(&removes).Error; err != nil {
			return err
		}
	}
	// 新增
	adds := slice.Filter(newTax, func(index int, item mCustomer.CustomerNoTaxArea) bool {
		_, ok := slice.FindBy(oldTax, func(index int, area mCustomer.CustomerNoTaxArea) bool {
			return area.CountryCode == item.CountryCode
		})
		return !ok
	})
	if len(adds) > 0 {
		if err = s.orm.Create(&adds).Error; err != nil {
			return err
		}
	}
	return err
}
