package sUser

import (
	"errors"
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/shop/user/mUser"
	"shopkone-service/utility/code"
)

type IUserColumn interface {
	// Create 创建列
	Create(staffId uint) (err error)
	// GetAllColumnsTypes 获取所有列类型
	GetAllColumnsTypes() []mUser.UserColumnType
	// Set 设置列
	Set(userId uint, req vo.SetColumnsReq) (err error)
	// Get 获取列
	Get(userId uint, req vo.GetColumnsReq) (res vo.GetColumnsRes, err error)
}

type iUserColumn struct {
	orm *gorm.DB
}

func NewUserColumn(orm *gorm.DB) *iUserColumn {
	return &iUserColumn{orm: orm}
}

func (s *iUserColumn) Create(userId uint) (err error) {
	// 获取所有列类型
	allTypes := s.GetAllColumnsTypes()
	// 获取该用户下所有列类型
	var oldList []mUser.UserColumn
	if err = s.orm.Model(&oldList).Where("user_id = ?", userId).Find(&oldList).Error; err != nil {
		return err
	}
	// 过滤掉已存在的列类型
	allTypes = slice.Filter(allTypes, func(index int, item mUser.UserColumnType) bool {
		_, ok := slice.FindBy(oldList, func(index int, i mUser.UserColumn) bool {
			return i.Type == item
		})
		return !ok
	})
	// 生成要创建的数据
	if len(allTypes) == 0 {
		return err
	}
	list := slice.Map(allTypes, func(index int, item mUser.UserColumnType) mUser.UserColumn {
		temp := mUser.UserColumn{}
		temp.UserId = userId
		temp.Type = item
		temp.Items = []mUser.UserColumnItem{}
		return temp
	})
	// 批量创建
	return s.orm.Create(&list).Error
}

func (s *iUserColumn) GetAllColumnsTypes() []mUser.UserColumnType {
	return []mUser.UserColumnType{
		mUser.UserColumnTypeProduct,
		mUser.UserColumnTypeVariant,
	}
}

func (s *iUserColumn) Set(userId uint, req vo.SetColumnsReq) (err error) {
	// 判断类型是否存在
	if !slice.Contain(s.GetAllColumnsTypes(), req.Type) {
		return code.UserColumnTypeErr
	}
	data := mUser.UserColumn{
		Items: req.Columns,
	}
	// name不允许重复
	names := slice.Map(req.Columns, func(index int, item mUser.UserColumnItem) string {
		return item.Name
	})
	if len(slice.Unique(names)) != len(req.Columns) {
		return code.UserColumnNameRepeatErr
	}
	// 列表不准为空
	if len(req.Columns) == 0 {
		return code.UserColumnListEmptyErr
	}
	// 列表长度不能超过100
	if len(req.Columns) > 100 {
		return code.UserColumnListTooLongErr
	}
	// 更新
	return s.orm.Model(mUser.UserColumn{}).
		Where("user_id = ? AND type = ?", userId, req.Type).
		Select("items").Updates(data).Error
}

func (s *iUserColumn) Get(userId uint, req vo.GetColumnsReq) (res vo.GetColumnsRes, err error) {
	// 判断类型是否存在
	if !slice.Contain(s.GetAllColumnsTypes(), req.Type) {
		return res, code.UserColumnTypeErr
	}
	// 查询
	var column mUser.UserColumn
	err = s.orm.Model(mUser.UserColumn{}).
		Where("user_id = ? AND type = ?", userId, req.Type).
		First(&column).Error
	res.Columns = column.Items
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 创建类型
		return res, s.Create(userId)
	}
	return res, err
}
