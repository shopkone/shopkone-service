package vo

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/shop/staff/mStaff"
	"shopkone-service/utility/handle"
)

type StaffListReq struct {
	g.Meta  `path:"/staff/list" method:"post" tags:"员工" summary:"获取员工列表"`
	Status  mStaff.StaffStatus `json:"status"`
	Keyword string             `json:"keyword"`
	handle.PageReq
}
type StaffListRes struct {
	ID            uint               `json:"id"`
	Name          string             `json:"name"`
	Status        mStaff.StaffStatus `json:"status"`
	Note          string             `json:"note"`
	IsMaster      bool               `json:"is_master"`
	Email         string             `json:"email"`
	LastInvitedAt int64              `json:"last_invited_at"`
	OnJobAt       int64              `json:"on_job_at"`
	LeaveAt       int64              `json:"leave_at"`
}

// StaffCreate 创建员工
type StaffCreateReq struct {
	g.Meta `path:"/staff/create" method:"post" tags:"员工" summary:"创建员工"`
	Name   string `json:"name" v:"required#员工名称不能为空"`
	Email  string `json:"email" v:"required#邮箱不能为空"`
	Note   string `json:"note"`
}
type StaffCreateRes struct {
	ID uint `json:"id"`
}

// SendInvite 发送邀请邮件
type StaffSendInviteReq struct {
	g.Meta `path:"/staff/send-invite" method:"post" tags:"员工" summary:"发送邀请邮件"`
	Email  string `json:"email" v:"required#邮箱不能为空"`
	Note   string `json:"note"`
}
type StaffSendInviteRes struct {
}

// StaffInfoReq 获取员工信息
type StaffInfoReq struct {
	g.Meta `path:"/staff/info" method:"post" tags:"员工" summary:"获取员工信息"`
	ID     uint `json:"id" v:"required#员工ID不能为空"`
}
type StaffInfoRes struct {
	ID            uint               `json:"id"`
	Name          string             `json:"name"`
	Status        mStaff.StaffStatus `json:"status"`
	Note          string             `json:"note"`
	IsMaster      bool               `json:"is_master"`
	Email         string             `json:"email"`
	LastInvitedAt int64              `json:"last_invited_at"`
}

// StaffUpdate 更新员工信息
type StaffUpdateReq struct {
	g.Meta `path:"/staff/update" method:"post" tags:"员工" summary:"更新员工信息"`
	ID     uint   `json:"id" v:"required#员工ID不能为空"`
	Name   string `json:"name" v:"required#员工名称不能为空"`
	Note   string `json:"note"`
}
type StaffUpdateRes struct{}

// Leave 离职员工
type StaffLeaveReq struct {
	g.Meta `path:"/staff/leave" method:"post" tags:"员工" summary:"离职员工"`
	ID     uint `json:"id" v:"required#员工ID不能为空"`
}
type StaffLeaveRes struct{}

// Remove 删除员工
type StaffRemoveReq struct {
	g.Meta `path:"/staff/remove" method:"post" tags:"员工" summary:"删除员工"`
	ID     uint `json:"id" v:"required#员工ID不能为空"`
}
type StaffRemoveRes struct{}

// SendChangeCode 发送转让验证码
type SendChangeCodeReq struct {
	g.Meta `path:"/staff/send-change-code" method:"post" tags:"员工" summary:"发送转让验证码"`
	Email  string `json:"email" v:"required#邮箱不能为空"`
}
type SendChangeCodeRes struct {
}

// ChangeStaffReq 转让店铺
type ChangeStaffReq struct {
	g.Meta `path:"/staff/change" method:"post" tags:"员工" summary:"转让员工"`
	Email  string `json:"email" v:"required#邮箱不能为空"`
	Code   string `json:"code" v:"required#验证码不能为空"`
}
type ChangeStaffRes struct{}
