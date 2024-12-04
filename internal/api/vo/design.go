package vo

import "github.com/gogf/gf/v2/frame/g"

type PartData struct {
	Name     string                 `json:"name"`
	Type     string                 `json:"type"`
	Order    []string               `json:"order"`
	Sections map[string]SectionData `json:"sections"`
}

type SectionData struct {
	Type       string               `json:"type"`
	BlockOrder []string             `json:"block_order"`
	Settings   g.Map                `json:"settings"`
	Blocks     map[string]BlockData `json:"blocks"`
}

type BlockData struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Settings g.Map  `json:"settings"`
}

type DesignDataListReq struct {
	g.Meta `path:"/design/data/list" method:"post" tags:"Design" summary:"获取设计数据列表"`
}
type DesignDataListRes struct {
	g.Meta          `path:"/design/data/list" method:"post" tags:"Design" summary:"获取设计数据列表"`
	HeaderData      PartData `json:"header_data"`
	FooterData      PartData `json:"footer_data"`
	CurrentPageData PartData `json:"current_page_data"`
	SettingData     g.Map    `json:"setting_data"`
}

/*------------------------------- 分隔线 -------------------------------*/

type SectionSchema struct {
	Type     string          `json:"type"`
	Name     string          `json:"name"`
	Class    string          `json:"class"`
	Blocks   []BlockSchema   `json:"blocks"`
	Settings []SettingSchema `json:"settings"`
}

type BlockSchema struct {
	Type     string          `json:"type"`
	Name     string          `json:"name"`
	Settings []SettingSchema `json:"settings"`
}

type SettingSchema struct {
	Type    string                `json:"type,omitempty"`
	Name    string                `json:"name,omitempty"`
	ID      string                `json:"id,omitempty"`
	Max     int                   `json:"max,omitempty"`
	Min     int                   `json:"min,omitempty"`
	Step    float64               `json:"step,omitempty"`
	Unit    string                `json:"unit,omitempty"`
	Label   string                `json:"label,omitempty"`
	Options []SettingSchemaOption `json:"options,omitempty"`
	Default interface{}           `json:"default,omitempty"`
	Content string                `json:"content,omitempty"`
}

type SettingSchemaOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// 获取头部
type DesignSchemaListReq struct {
	g.Meta `path:"/design/schema/list" method:"post" tags:"Design" summary:"获取设计数据列表"`
	Type   []string `json:"type"`
}
type DesignSchemaListRes SectionSchema

// 更新block
type DesignUpdateBlockReq struct {
	g.Meta    `path:"/design/block/update" method:"post" tags:"Design" summary:"更新设计数据列表"`
	SectionID string      `json:"section_id" v:"required"`
	BlockID   string      `json:"block_id" v:"required"`
	PartName  string      `json:"part_name" v:"required"`
	Key       string      `json:"key" v:"required"`
	Value     interface{} `json:"value"`
}
type DesignUpdateBlockRes struct {
}

// 获取部分section
type DesignSectionRenderReq struct {
	g.Meta    `path:"/design/section/render" method:"post" tags:"Design" summary:"获取部分section"`
	PartName  string `json:"part_name" v:"required"`
	SectionID string `json:"section_id" v:"required"`
}
type DesignSectionRenderRes struct {
	HTML string `json:"html"`
}
