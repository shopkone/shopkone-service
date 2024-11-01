package sFile

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/file/mFile"
	"shopkone-service/internal/module/shop/shop/sShop"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

type IFile interface {
	// Add 添加文件
	Add(in vo.AddFileReq) (id uint, err error)
	// List 获取文件列表
	List(in vo.FileListReq) (res handle.PageRes[vo.FileListRes], err error)
	// UpdateGroupIdsByOldGroupId 根据老分组id更新分组id
	UpdateGroupIdsByOldGroupId(oldGroupId, newGroupId uint) error
	// UpdateGroupIdByFileIds 根据文件id更新分组id
	UpdateGroupIdByFileIds(fileIds []uint, groupId uint) error
	// FilesDelete 删除文件
	FilesDelete(ids []uint) error
	// FileInfo 获取文件信息
	FileInfo(id uint) (res vo.FileInfoRes, err error)
	// UpdateFileInfo 更新文件信息
	FileUpdateInfo(in vo.FileUpdateInfoReq) (err error)
	// FileListByIds 根据文件id获取文件信息
	FileListByIds(ids []uint) ([]vo.FileListByIdsRes, error)
}

type sFile struct {
	orm    *gorm.DB
	shopId uint
}

func NewFile(orm *gorm.DB, shopId uint) *sFile {
	return &sFile{orm: orm, shopId: shopId}
}

func (s *sFile) Add(in vo.AddFileReq) (id uint, err error) {
	if in.Type == mFile.FileTypeImage {
		in.Cover = ""
	}
	data := mFile.File{
		Name:     in.Name,
		Path:     in.Path,
		Source:   in.Source,
		Size:     in.Size,
		Alt:      in.Alt,
		Type:     in.Type,
		Width:    in.Width,
		Height:   in.Height,
		Duration: in.Duration,
		Suffix:   in.Suffix,
		GroupId:  in.GroupId,
		Cover:    in.Cover,
	}
	data.ShopId = s.shopId
	if err = s.orm.Create(&data).Error; err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (s *sFile) List(in vo.FileListReq) (res handle.PageRes[vo.FileListRes], err error) {
	var list []mFile.File
	query := s.orm.Model(&mFile.File{}).Where("shop_id = ?", s.shopId)
	query = query.Select("id", "name", "path", "suffix", "size", "created_at", "group_id", "cover", "type")
	// 筛选分组
	if in.GroupId > 0 {
		query = query.Where("group_id = ?", in.GroupId)
	}
	// 筛选文件类型
	if in.FileType != nil && len(in.FileType) > 0 {
		query = query.Where("type in ?", in.FileType)
	} else {
		res.List = []vo.FileListRes{}
		return res, err
	}
	// 筛选文件大小
	if in.FileSize.Min != 0 {
		query = query.Where("size >= ?", in.FileSize.Min*1024*1024)
	}
	// 搜索
	if in.Keyword != "" {
		query = query.Where("name like ?", "%"+in.Keyword+"%")
	}
	if in.FileSize.Min != 0 {
		query = query.Where("size <= ?", in.FileSize.Max*1024*1024)
	}
	// 查找总数
	if err = query.Count(&res.Total).Error; err != nil {
		return res, err
	}
	// 分页
	query = query.Scopes(handle.Pagination(in.PageReq))
	// 根据创建时间倒序
	query = query.Order("created_at desc")
	// 查找数据
	if err = query.Find(&list).Error; err != nil {
		return res, err
	}
	res.List = slice.Map(list, func(index int, item mFile.File) vo.FileListRes {
		temp := vo.FileListRes{}
		temp.Id = item.ID
		temp.FileName = item.Name
		temp.DataAdded = item.CreatedAt.UnixMilli()
		temp.Suffix = item.Suffix
		temp.Src = item.Path
		temp.Size = item.Size
		temp.GroupId = item.GroupId
		temp.Cover = item.Cover
		temp.Type = item.Type
		return temp
	})
	res.Page = in.PageReq
	return res, err
}

func (s *sFile) UpdateGroupIdsByOldGroupId(oldGroupId, newGroupId uint) error {
	query := s.orm.Model(&mFile.File{}).Where("shop_id = ?", s.shopId)
	query = query.Where("group_id = ?", oldGroupId)
	return query.Update("group_id", newGroupId).Error
}

func (s *sFile) FilesDelete(ids []uint) error {
	if err := sShop.NewShop(s.orm).RemoveCoverByFileIds(ids, s.shopId); err != nil {
		return err
	}
	query := s.orm.Model(&mFile.File{}).Where("shop_id = ?", s.shopId)
	query = query.Where("id IN ?", ids)
	return query.Delete(&mFile.File{}).Error
}

func (s *sFile) FileInfo(id uint) (res vo.FileInfoRes, err error) {
	var file mFile.File
	query := s.orm.Model(&mFile.File{}).Where("shop_id = ?", s.shopId)
	query = query.Where("id = ?", id)
	if err = query.First(&file).Error; err != nil {
		return res, err
	}
	return vo.FileInfoRes{
		Name:      file.Name,
		Path:      file.Path,
		Size:      file.Size,
		Source:    file.Source,
		Suffix:    file.Suffix,
		Type:      file.Type,
		Width:     file.Width,
		Height:    file.Height,
		Duration:  file.Duration,
		Alt:       file.Alt,
		GroupId:   file.GroupId,
		Cover:     file.Cover,
		CreatedAt: file.CreatedAt.UnixMilli(),
	}, nil
}

func (s *sFile) FileUpdateInfo(in vo.FileUpdateInfoReq) (err error) {
	var file mFile.File
	query := s.orm.Model(&mFile.File{}).Where("shop_id = ?", s.shopId)
	query = query.Where("id = ?", in.Id)
	if err = query.First(&file).Error; err != nil {
		return err
	}
	file.Alt = in.Alt
	if in.Name != "" {
		file.Name = in.Name
	}
	if in.Cover != "" {
		file.Cover = in.Cover
	}
	if in.Src != "" {
		file.Path = in.Src
	}
	if file.Type == mFile.FileTypeImage {
		file.Cover = ""
	}
	if err = s.orm.Model(&file).Where("shop_id = ? AND id = ?", s.shopId, in.Id).
		Select("name", "alt", "cover", "path").Updates(&file).Error; err != nil {
		return err
	}
	return nil
}

func (s *sFile) UpdateGroupIdByFileIds(fileIds []uint, groupId uint) error {
	// 校验groupId是否存在
	exist, err := NewFileGroup(s.orm, s.shopId).Exist(groupId)
	if err != nil || !exist {
		return code.ErrFileGroupNotExist
	}
	// 更新文件分组
	query := s.orm.Model(&mFile.File{}).Where("shop_id = ?", s.shopId)
	query = query.Where("id IN ?", fileIds)
	return query.Update("group_id", groupId).Error
}

func (s *sFile) FileListByIds(ids []uint) ([]vo.FileListByIdsRes, error) {
	var list []mFile.File
	query := s.orm.Model(&mFile.File{}).Where("shop_id = ?", s.shopId)
	query = query.Where("id IN ?", ids)
	if err := query.Select("id", "path", "type", "cover").Find(&list).Error; err != nil {
		return nil, err
	}
	return slice.Map(list, func(index int, item mFile.File) vo.FileListByIdsRes {
		return vo.FileListByIdsRes{
			Id:    item.ID,
			Path:  item.Path,
			Type:  item.Type,
			Cover: item.Cover,
		}
	}), nil
}
