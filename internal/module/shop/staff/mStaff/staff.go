package mStaff

import (
	"shopkone-service/internal/module/base/orm/mOrm"
	"time"
)

type StaffStatus uint8

const (
	STAFF_ON_JOB         StaffStatus = iota + 1 // 在职
	STAFF_LEAVE                                 // 离职
	STAFF_REJECT_INVITE                         // 拒绝邀请
	STAFF_EXPIRED_INVITE                        // 邀请过期
	STAFF_SEND_INVITE                           // 等待接收邀请
)

/*Staff 员工*/
type Staff struct {
	mOrm.Model
	Name          string      `gorm:"size:120;index"`      // 姓名
	IsMaster      bool        `gorm:"default:false;index"` // 是否是管理员
	Status        StaffStatus `gorm:"index"`               // 状态
	InviteToken   string      `gorm:"size:200"`            // 邀请token
	Note          string      `gorm:"size:500"`            // 备注
	LastInvitedAt *time.Time  `gorm:"default:null"`        // 最后邀请时间
	OnJobAt       *time.Time  `gorm:"default:null"`        // 入职时间
	LeaveAt       *time.Time  `gorm:"default:null"`        // 离职时间
	UserId        uint        `gorm:"index;not null"`      // 用户ID
}
