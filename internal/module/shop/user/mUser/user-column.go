package mUser

type UserColumnType string

const (
	UserColumnTypeProduct UserColumnType = "product"
	UserColumnTypeVariant UserColumnType = "variant"
)

type UserColumnItem struct {
	Name     string `json:"name,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Lock     bool   `json:"lock,omitempty"`
	Hidden   bool   `json:"hidden,omitempty"`
	Required bool   `json:"required,omitempty"`
}

type UserColumn struct {
	Type   UserColumnType   `gorm:"index;not null"`
	UserId uint             `gorm:"index;not null"`
	Items  []UserColumnItem `gorm:"serializer:json"`
}
