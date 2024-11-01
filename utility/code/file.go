package code

var (
	ErrFileGroupNameExist = fileError(1, "文件分组名称已存在")
	ErrFileGroupNotExist  = fileError(2, "文件分组不存在")
)
