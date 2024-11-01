package sStaff

import (
	"gorm.io/gorm"
	"shopkone-service/internal/module/shop/staff/mStaff"
	"shopkone-service/utility/handle"
)

type IStaff interface {
	// CreateOnJobStaff 创建在职员工
	CreateOnJobStaff(userID uint, name string, isMaster bool) (staff mStaff.Staff, err error)
	// StaffsByUserId 根据用户ID获取员工列表
	StaffsByUserId(userId uint) (staffs []mStaff.Staff, err error)
}

type sStaff struct {
	orm    *gorm.DB
	shopId uint
}

func NewStaff(orm *gorm.DB, shopId uint) *sStaff {
	return &sStaff{orm, shopId}
}

func (s *sStaff) CreateOnJobStaff(userID uint, name string, isMaster bool) (staff mStaff.Staff, err error) {
	data := mStaff.Staff{}
	data.UserId = userID
	data.ShopId = s.shopId
	data.IsMaster = isMaster
	data.Name = name
	data.OnJobAt = handle.GetNowTime()
	data.Status = mStaff.STAFF_ON_JOB
	// 添加员工
	err = s.orm.Model(&data).Create(&data).Error
	if err != nil {
		return data, err
	}
	// 更新缓存
	if err = NewStaffCache().UpdateCache(data.ID, s.orm); err != nil {
		return data, err
	}
	return staff, err
}

func (s *sStaff) StaffsByUserId(userId uint) (staffs []mStaff.Staff, err error) {
	query := s.orm.Model(&mStaff.Staff{}).Where("user_id = ?", userId)
	err = query.Select("shop_id").Find(&staffs).Error
	return staffs, err
}
