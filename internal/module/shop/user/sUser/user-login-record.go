package sUser

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"shopkone-service/internal/module/shop/user/iUser"
	"shopkone-service/internal/module/shop/user/mUser"
	"shopkone-service/utility/handle"
)

type IUserLoginRecord interface {
	// AddRecord 添加登录记录
	AddRecord(in iUser.LoginRecordIn) (token string, err error)
}

type sLoginRecord struct {
	orm *gorm.DB
}

func NewUserLoginRecord(orm *gorm.DB) *sLoginRecord {
	return &sLoginRecord{orm}
}

func (s *sLoginRecord) AddRecord(in iUser.LoginRecordIn) (token string, err error) {
	token = handle.GenUid()
	record := mUser.UserLoginRecord{
		Ip:       in.Ip,
		Ua:       in.Ua,
		Token:    token,
		IsActive: true,
		UserId:   in.UserId,
	}
	// 添加记录
	if err = s.orm.Create(&record).Error; err != nil {
		return "", err
	}
	// 添加缓存
	if err = NewUserCache().UpdateUserLoginCache(token, in.UserId); err != nil {
		return "", err
	}
	// 获取全部已登录缓存记录
	var records []mUser.UserLoginRecord
	if err = s.orm.Where("user_id = ? AND is_active = ?", in.UserId, true).Find(&records).Error; err != nil {
		return "", err
	}
	// 获取超出3个的记录
	overThree := slice.Filter(records, func(index int, item mUser.UserLoginRecord) bool {
		return index > 2
	})
	overThreeToken := slice.Map(overThree, func(index int, item mUser.UserLoginRecord) string {
		return item.Token
	})
	overThreeId := slice.Map(overThree, func(index int, item mUser.UserLoginRecord) uint {
		return item.ID
	})
	// 将超出3个之外的登录记录设为false
	if err = s.orm.Model(&mUser.UserLoginRecord{}).Where("id IN ?", overThreeId).Update("is_active", false).Error; err != nil {
		return "", err
	}
	// 删除超过3个之外的登录缓存
	if err = NewUserCache().Removes(overThreeToken); err != nil {
		return "", err
	}
	return token, err
}
