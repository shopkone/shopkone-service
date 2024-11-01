package handle

import "gorm.io/gorm"

/*Pagination 定义 Scope 分页查询*/
func Pagination(pageRequest PageReq) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pageRequest.Page
		pageSize := pageRequest.PageSize
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

/*PageRes 分页响应*/
type PageRes[T any] struct {
	Total int64   `json:"total"`
	List  []T     `json:"list"`
	Page  PageReq `json:"page"`
}

/*PageReq 分页请求*/
type PageReq struct {
	Page     int `json:"page" v:"required|min:1|max:100" dc:"页码"`
	PageSize int `json:"page_size" v:"required|min:1|max:100" dc:"每页数量"`
}
