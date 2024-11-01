package sStaff

import (
	"github.com/duke-git/lancet/v2/convertor"
	"gorm.io/gorm"
	"shopkone-service/internal/module/base/cache/sCache"
	"shopkone-service/internal/module/shop/staff/mStaff"
)

type IStaffCache interface {
	GetCache(shopId, userId uint) (staff mStaff.Staff, err error)
	UpdateCache(staffId uint, orm *gorm.DB) (err error)
	getKey(shopId, userId uint) string
}

type sStaffCache struct{}

func NewStaffCache() *sStaffCache {
	return &sStaffCache{}
}

func (s *sStaffCache) GetCache(shopId, userId uint) (staff mStaff.Staff, err error) {
	key := s.getKey(shopId, userId)
	err = sCache.NewStaffCache().Get(key, &staff)
	return staff, err
}

func (s *sStaffCache) UpdateCache(staffId uint, orm *gorm.DB) (err error) {
	var staff mStaff.Staff
	if err = orm.Where("id = ?", staffId).First(&staff).Error; err != nil {
		return err
	}
	key := s.getKey(staff.ShopId, staff.UserId)
	year100 := 60 * 24 * 365 * 100 //100年
	return sCache.NewStaffCache().Set(key, staff, uint(year100))
}

func (s *sStaffCache) getKey(shopId, userId uint) string {
	return "staff_" + convertor.ToString(shopId) + "_" + convertor.ToString(userId)
}
