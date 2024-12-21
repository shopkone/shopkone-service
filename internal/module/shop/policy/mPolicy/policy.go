package mPolicy

import "shopkone-service/internal/module/base/orm/mOrm"

type PolicyType uint8

const (
	PolicyTypePrivacy  PolicyType = iota + 1 // 隐私政策
	PolicyTypeRefund   PolicyType = iota + 2 // 退款政策
	PolicyTypeService  PolicyType = iota + 3 // 服务政策
	PolicyTypeShipping PolicyType = iota + 4 // 运费政策
)

type Policy struct {
	mOrm.Model
	Body  string
	Title string
	Url   string
	Type  PolicyType
}
