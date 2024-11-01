package sFile

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/file/mFile"
	"shopkone-service/utility/code"
)

type IFileGroup interface {
	// 添加文件分组
	Add(in vo.FileGroupAddReq) (id uint, err error)
	// 获取文件分组列表
	List() (res []vo.FileGroupListRes, err error)
	// 更新文件分组名称
	Update(id uint, name string) (err error)
	// 删除文件分组
	Delete(id uint) (err error)
	// 判断分组是否存在
	Exist(id uint) (bool, error)
}

type sFileGroup struct {
	gorm   *gorm.DB
	shopId uint
}

func NewFileGroup(orm *gorm.DB, shopId uint) *sFileGroup {
	return &sFileGroup{gorm: orm, shopId: shopId}
}

func (s *sFileGroup) Add(in vo.FileGroupAddReq) (id uint, err error) {
	// 校验名称
	var count int64
	if err = s.gorm.Model(&mFile.FileGroup{}).
		Where("shop_id = ? AND name = ?", s.shopId, in.Name).
		Count(&count).Error; err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, code.ErrFileGroupNameExist
	}
	// 添加分组
	data := mFile.FileGroup{}
	data.ShopId = s.shopId
	data.Name = in.Name
	if err = s.gorm.Create(&data).Error; err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (s *sFileGroup) Update(id uint, name string) (err error) {
	var count int64
	if err = s.gorm.Model(&mFile.FileGroup{}).
		Where("shop_id = ? AND name = ? AND id != ?", s.shopId, name, id).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return code.ErrFileGroupNameExist
	}
	return s.gorm.Model(&mFile.FileGroup{}).
		Where("shop_id = ? AND id = ?", s.shopId, id).
		Update("name", name).Error
}

func (s *sFileGroup) List() (res []vo.FileGroupListRes, err error) {
	var list []mFile.FileGroup
	if err = s.gorm.Model(&mFile.FileGroup{}).
		Where("shop_id = ?", s.shopId).
		Select("id", "name").
		Find(&list).Error; err != nil {
		return res, err
	}
	res = slice.Map(list, func(index int, item mFile.FileGroup) vo.FileGroupListRes {
		return vo.FileGroupListRes{
			Id:   item.ID,
			Name: item.Name,
		}
	})
	return res, nil
}

func (s *sFileGroup) Delete(id uint) (err error) {
	err = s.gorm.Model(&mFile.FileGroup{}).Where("shop_id = ? AND id = ?", s.shopId, id).Delete(&mFile.FileGroup{}).Error
	if err != nil {
		return err
	}
	return NewFile(s.gorm, s.shopId).UpdateGroupIdsByOldGroupId(id, 0)
}

func (s *sFileGroup) Exist(id uint) (bool, error) {
	if id == 0 {
		return true, nil
	}
	var count int64
	if err := s.gorm.Model(&mFile.FileGroup{}).
		Where("shop_id = ? AND id = ?", s.shopId, id).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
