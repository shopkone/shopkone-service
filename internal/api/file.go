package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/ali/sAli"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/file/sFile"
	ctx2 "shopkone-service/utility/ctx"
	"shopkone-service/utility/handle"
)

type aFile struct {
}

func NewFileApi() *aFile {
	return &aFile{}
}

// AddFile 添加文件
func (a *aFile) AddFile(ctx g.Ctx, req *vo.AddFileReq) (res vo.AddFileRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		s := sFile.NewFile(tx, shop.ID)
		res.ID, err = s.Add(*req)
		return err
	})
	return res, err
}

// FileList 文件列表
func (a *aFile) FileList(ctx g.Ctx, req *vo.FileListReq) (res handle.PageRes[vo.FileListRes], err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sFile.NewFile(sOrm.NewDb(), shop.ID).List(*req)
}

// FileDelete 删除文件
func (a *aFile) FilesDelete(ctx g.Ctx, req *vo.FilesDeleteReq) (res vo.FilesDeleteRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sFile.NewFile(tx, shop.ID).FilesDelete(req.Ids)
	})
	return res, err
}

// FileInfo 文件详情
func (a *aFile) FileInfo(ctx g.Ctx, req *vo.FileInfoReq) (res vo.FileInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sFile.NewFile(sOrm.NewDb(), shop.ID).FileInfo(req.Id)
}

// FileUpdateInfo 更新文件信息
func (a *aFile) FileUpdateInfo(ctx g.Ctx, req *vo.FileUpdateInfoReq) (res vo.FileUpdateInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sFile.NewFile(tx, shop.ID).FileUpdateInfo(*req)
	})
	return res, err
}

// FileListByIds 根据文件ID列表获取文件信息
func (a *aFile) FileListByIds(ctx g.Ctx, req *vo.FileListByIdsReq) (res []vo.FileListByIdsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sFile.NewFile(sOrm.NewDb(), shop.ID).FileListByIds(req.Ids)
}

// UpdateGroupIdByFileIds 更新文件分组
func (a *aFile) UpdateGroupIdByFileIds(ctx g.Ctx, req *vo.UpdateGroupIdByFileIdsReq) (res vo.UpdateGroupIdByFileIdsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sFile.NewFile(tx, shop.ID).UpdateGroupIdByFileIds(req.FileIds, req.GroupId)
	})
	return res, err
}

// FileGroupList 文件分组列表
func (a *aFile) FileGroupList(ctx g.Ctx, req *vo.FileGroupListReq) (res []vo.FileGroupListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	res, err = sFile.NewFileGroup(sOrm.NewDb(), auth.Shop.ID).List()
	return res, err
}

// FileGroupAdd 文件分组添加
func (a *aFile) FileGroupAdd(ctx g.Ctx, req *vo.FileGroupAddReq) (res vo.FileGroupAddRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		s := sFile.NewFileGroup(tx, auth.Shop.ID)
		res.ID, err = s.Add(*req)
		return err
	})
	return res, err
}

// FileGroupUpdate 文件分组更新
func (a *aFile) FileGroupUpdate(ctx g.Ctx, req *vo.FileGroupUpdateReq) (res vo.FileGroupUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sFile.NewFileGroup(tx, auth.Shop.ID).Update(req.Id, req.Name)
	})
	res.ID = req.Id
	return res, err
}

// FileGroupRemove 文件分组删除
func (a *aFile) FileGroupRemove(ctx g.Ctx, req *vo.FileGroupRemoveReq) (res vo.FileGroupRemoveRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sFile.NewFileGroup(tx, auth.Shop.ID).Delete(req.Id)
	})
	return res, err
}

// GetUploadToken 获取上传token
func (a *aFile) GetUploadToken(ctx g.Ctx, req *vo.GetUploadTokenReq) (res vo.GetUploadTokenRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	res.Token, err = sAli.NewOss().GetUploadToken(shop.Uuid)
	return res, err
}
