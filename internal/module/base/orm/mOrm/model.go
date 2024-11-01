package mOrm

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID          uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	ShopId      uint           `json:"shop_id,omitempty" gorm:"index;not null"`
	CanCreateId bool           `json:"-" gorm:"-"`
}

func (m Model) GetID() uint {
	return m.ID
}

/*Base Hooks */
func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if m.ShopId == 0 {
		return gerror.New("THE SHOP ID CAN NOT ZERO!")
	}
	if m.ID != 0 && !m.CanCreateId {
		return gerror.New("THE ID NEED IS ZERO!")
	}
	return nil
}
