package mNav

import "shopkone-service/internal/module/base/orm/mOrm"

type Nav struct {
	mOrm.Model
	Handle string
	Title  string    `gorm:"index;not null"`
	Links  []NavItem `gorm:"serializer:json"`
}

type NavItem struct {
	ID    string    `json:"id" v:"required"`
	Title string    `json:"title" v:"required"`
	Url   string    `json:"url"`
	Links []NavItem `json:"links"`
}
