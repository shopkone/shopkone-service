package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/seo/mSeo"
	"shopkone-service/internal/module/product/collection/mCollection"
	"shopkone-service/utility/handle"
)

type BaseCondition struct {
	Id     uint   `json:"id" dc:"条件ID"`
	Action string `json:"action" dc:"操作类型" v:"required"`
	Key    string `json:"key" dc:"条件键" v:"required"`
	Value  string `json:"value" dc:"条件值" v:"required"`
}

// CreateProductCollectionReq 创建商品集合
type CreateProductCollectionReq struct {
	g.Meta         `path:"/product/collection/create" method:"post" tags:"Product" summary:"创建商品集合"`
	Title          string                          `json:"title" v:"required" dc:"集合标题"`
	Description    string                          `json:"description" dc:"集合描述"`
	CollectionType mCollection.CollectionType      `json:"collection_type" v:"required" dc:"集合模式"`
	MatchMode      mCollection.CollectionMatchMode `json:"match_mode" dc:"集合匹配模式"`
	CoverId        uint                            `json:"cover_id" dc:"集合封面"`
	Seo            mSeo.Seo                        `json:"seo" v:"required" dc:"seo信息"`
	Conditions     []BaseCondition                 `json:"conditions" dc:"集合条件"`
	ProductIds     []uint                          `json:"product_ids"`
}
type CreateProductCollectionRes struct {
	Id uint `json:"id" dc:"集合ID"`
}

// ProductCollectionInfoReq 获取商品集合详情
type ProductCollectionInfoReq struct {
	Id     uint `json:"id" v:"required" dc:"集合ID"`
	g.Meta `path:"/product/collection/info" method:"post" tags:"Product" summary:"获取商品集合详情"`
}
type ProductCollectionInfoRes struct {
	CreateProductCollectionReq
	Id uint `json:"id" dc:"集合ID"`
}

// ProductCollectionListReq 获取商品集合列表
type ProductCollectionListReq struct {
	g.Meta `path:"/product/collection/list" method:"post" tags:"Product" summary:"获取商品集合列表"`
	handle.PageReq
}
type ProductCollectionListRes struct {
	Id              uint                       `json:"id" dc:"集合ID"`
	Title           string                     `json:"title" dc:"标题"`
	Cover           string                     `json:"cover" dc:"封面"`
	CollectionType  mCollection.CollectionType `json:"collection_type" dc:"类型"`
	ProductQuantity int                        `json:"product_quantity" dc:"商品数量"`
}

// UpdateProductCollectionReq 更新商品集合
type UpdateProductCollectionReq struct {
	g.Meta `path:"/product/collection/update" method:"post" tags:"Product" summary:"更新商品集合"`
	Id     uint `json:"id" v:"required" dc:"集合ID"`
	CreateProductCollectionReq
}
type UpdateProductCollectionRes struct {
}

// RemoveProductCollectionReq 删除商品集合
type RemoveProductCollectionReq struct {
	g.Meta `path:"/product/collection/remove" method:"post" tags:"Product" summary:"删除商品集合"`
	Id     uint `json:"id" v:"required" dc:"集合ID"`
}
type RemoveProductCollectionRes struct {
}

type CollectionOptionsReq struct {
	g.Meta `path:"/product/collection/options" method:"post" tags:"Product" summary:"获取商品集合选项"`
}
type CollectionOptionsRes struct {
	Value    uint   `json:"value"`
	Label    string `json:"label"`
	Disabled bool   `json:"disabled"`
}
