package mFile

import "shopkone-service/internal/module/base/orm/mOrm"

type Source uint8

const (
	SourceLocal  Source = iota + 1 // 1. 本地
	SourceRemote                   // 2. 远程
	SourceOther                    // 3. 其他
)

type FileType uint8

const (
	FileTypeImage FileType = iota + 1 // 1. 图片
	FileTypeVideo                     // 2. 视频
	FileTypeAudio                     // 3. 音频
)

type File struct {
	mOrm.Model
	Name     string   `gorm:"size:500"`  // 文件名
	Path     string   `gorm:"size:500"`  // 文件路径
	Source   Source   `gorm:"default:1"` // 文件来源
	GroupId  uint     `gorm:"index"`     // 文件分组
	Size     uint     `gorm:"not null"`  // 文件大小
	Alt      string   `gorm:"size:500"`  // 文件描述
	Type     FileType `gorm:"default:1"` // 文件类型
	Width    uint     `gorm:"default:0"` // 图片宽度
	Height   uint     `gorm:"default:0"` // 图片高度
	Duration uint     `gorm:"default:0"` // 音频时长
	Suffix   string   `gorm:"size:20"`   // 文件后缀
	Cover    string   `gorm:"size:1000"` // 封面
}
