package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/setting/file/mFile"
	"shopkone-service/utility/handle"
)

// GetUploadToken 获取上传token
type GetUploadTokenReq struct {
	g.Meta `path:"/file/upload/token" method:"post" summary:"获取上传token" tags:"Base"`
}
type GetUploadTokenRes struct {
	Token string `json:"token"`
}

// AddFile 添加文件
type AddFileReq struct {
	g.Meta   `path:"/file/add" method:"post" summary:"添加文件" tags:"File"`
	Name     string         `json:"name" v:"required" dc:"文件名"`
	Path     string         `json:"path" v:"required" dc:"文件路径"`
	Source   mFile.Source   `json:"source" v:"required" dc:"文件来源"`
	Size     uint           `json:"size" v:"required" dc:"文件大小"`
	Alt      string         `json:"alt" dc:"文件描述"`
	Type     mFile.FileType `json:"type" v:"required" dc:"文件类型"`
	Width    uint           `json:"width" dc:"文件宽度"`
	Height   uint           `json:"height" dc:"文件高度"`
	Duration uint           `json:"duration" dc:"文件时长"`
	Suffix   string         `json:"suffix"  dc:"文件后缀"`
	GroupId  uint           `json:"group_id"`
	Cover    string         `json:"cover"`
}
type AddFileRes struct {
	ID uint `json:"id"`
}

type FileSizeRange struct {
	Max float32 `json:"max"`
	Min float32 `json:"min"`
}

// FileList 文件列表
type FileListReq struct {
	g.Meta   `path:"/file/list" method:"post" summary:"获取文件列表" tags:"文件"`
	GroupId  uint             `json:"group_id"`
	FileSize FileSizeRange    `json:"file_size"`
	FileType []mFile.FileType `json:"file_type"`
	Keyword  string           `json:"keyword"`
	handle.PageReq
}
type FileListRes struct {
	Id         uint           `json:"id"`
	FileName   string         `json:"file_name"`
	DataAdded  int64          `json:"data_added"`
	Suffix     string         `json:"suffix"`
	Src        string         `json:"src"`
	Size       uint           `json:"size"`
	References uint           `json:"references"`
	GroupId    uint           `json:"group_id"`
	Cover      string         `json:"cover"`
	Type       mFile.FileType `json:"type"`
}

// FilesDelete 删除文件
type FilesDeleteReq struct {
	g.Meta `path:"/files/delete" method:"post" summary:"删除文件" tags:"文件"`
	Ids    []uint `json:"ids" v:"required" dc:"文件ID列表"`
}
type FilesDeleteRes struct {
}

// FileInfo 文件信息
type FileInfoReq struct {
	g.Meta `path:"/file/info" method:"post" summary:"获取文件信息" tags:"文件"`
	Id     uint `json:"id" v:"required" dc:"文件ID"`
}
type FileInfoRes struct {
	Name      string         `json:"name" v:"required" dc:"文件名"`
	Path      string         `json:"path" v:"required" dc:"文件路径"`
	Source    mFile.Source   `json:"source" v:"required" dc:"文件来源"`
	Size      uint           `json:"size" v:"required" dc:"文件大小"`
	Alt       string         `json:"alt" dc:"文件描述"`
	Type      mFile.FileType `json:"type" v:"required" dc:"文件类型"`
	Width     uint           `json:"width" dc:"文件宽度"`
	Height    uint           `json:"height" dc:"文件高度"`
	Duration  uint           `json:"duration" dc:"文件时长"`
	Suffix    string         `json:"suffix"  dc:"文件后缀"`
	GroupId   uint           `json:"group_id"`
	Cover     string         `json:"cover"`
	CreatedAt int64          `json:"created_at"`
}

// FileUpdateInfo 更新文件信息
type FileUpdateInfoReq struct {
	g.Meta `path:"/file/update/info" method:"post" summary:"更新文件信息" tags:"文件"`
	Id     uint   `json:"id" v:"required" dc:"文件ID"`
	Name   string `json:"name" dc:"文件名"`
	Alt    string `json:"alt" dc:"文件描述"`
	Cover  string `json:"cover"`
	Src    string `json:"src"`
}
type FileUpdateInfoRes struct {
}

// FileListByIds 根据文件ID列表获取文件信息
type FileListByIdsReq struct {
	g.Meta `path:"/file/list/by_ids" method:"post" summary:"根据文件ID列表获取文件信息" tags:"文件"`
	Ids    []uint `json:"ids" v:"required" dc:"文件ID列表"`
}
type FileListByIdsRes struct {
	Id    uint           `json:"id"`
	Path  string         `json:"path"`
	Type  mFile.FileType `json:"type"`
	Cover string         `json:"cover"`
}

// UpdateGroupIdByFileIds 更新文件分组
type UpdateGroupIdByFileIdsReq struct {
	g.Meta  `path:"/file/update/group_id" method:"post" summary:"更新文件分组" tags:"文件"`
	FileIds []uint `json:"file_ids" v:"required" dc:"文件ID列表"`
	GroupId uint   `json:"group_id" v:"required" dc:"分组ID"`
}
type UpdateGroupIdByFileIdsRes struct {
}

// FileGroupList 文件分组列表
type FileGroupListReq struct {
	g.Meta `path:"/file/group/list" method:"post" summary:"获取文件分组列表" tags:"文件"`
}
type FileGroupListRes struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Count uint   `json:"count"`
}

// FileGroupAdd 添加分组
type FileGroupAddReq struct {
	g.Meta `path:"/file/group/add" method:"post" summary:"添加文件分组" tags:"文件"`
	Name   string `json:"name" v:"required" dc:"分组名称"`
}
type FileGroupAddRes struct {
	ID uint `json:"id"`
}

// FileGroupUpdate 更新分组
type FileGroupUpdateReq struct {
	g.Meta `path:"/file/group/update" method:"post" summary:"更新文件分组" tags:"文件"`
	Id     uint   `json:"id" v:"required" dc:"分组ID"`
	Name   string `json:"name" v:"required" dc:"分组名称"`
}
type FileGroupUpdateRes struct {
	ID uint `json:"id"`
}

// FileGroupRemove 删除分组
type FileGroupRemoveReq struct {
	g.Meta `path:"/file/group/remove" method:"post" summary:"删除文件分组" tags:"文件"`
	Id     uint `json:"id" v:"required" dc:"分组ID"`
}
type FileGroupRemoveRes struct {
}
