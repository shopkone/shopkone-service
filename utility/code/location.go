package code

var (
	ErrLocationNameExist                   = locationError(1, "地点名称已存在")
	ErrLocationDefaultDisable              = locationError(2, "默认地点不可禁用")
	ErrLocationDefaultDelete               = locationError(3, "默认地点不可删除")
	ErrLocationDefaultUnFulfillmentDetails = locationError(4, "默认地点不可修改配送信息")
	ErrLocationNotAllActive                = locationError(5, "部分地点未启用")
	ErrLocationDelete                      = locationError(6, "请将点地点设为禁用后再删除")
	NoDeActiveByHasInventory               = productError(7, "存在库存，无法停用地点")
	ErrLocationNotFound                    = locationError(8, "地点不存在或已停用")
)
