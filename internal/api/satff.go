package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/utility/handle"
)

type aStaff struct {
}

func NewStaffApi() *aStaff {
	return &aStaff{}
}

// ListStaff 获取员工列表
func (a *aStaff) ListStaff(ctx g.Ctx, req *vo.StaffListReq) (res handle.PageRes[vo.StaffListRes], err error) {
	return res, err
}

// CreateStaff 创建员工
func (a *aStaff) CreateStaff(ctx g.Ctx, req *vo.StaffCreateReq) (res vo.StaffCreateRes, err error) {
	return res, err
}

// SendInvite 发送邀请邮件
func (a *aStaff) SendInvite(ctx g.Ctx, req *vo.StaffSendInviteReq) (res vo.StaffSendInviteRes, err error) {
	return res, err
}

// Info 获取员工信息
func (a *aStaff) Info(ctx g.Ctx, req *vo.StaffInfoReq) (res vo.StaffInfoRes, err error) {
	return res, err
}

// StaffUpdate 更新员工信息
func (a *aStaff) StaffUpdate(ctx g.Ctx, req *vo.StaffUpdateReq) (res vo.StaffUpdateRes, err error) {
	return res, err
}

// Leave 离职员工
func (a *aStaff) Leave(ctx g.Ctx, req *vo.StaffLeaveReq) (res vo.StaffLeaveRes, err error) {
	return res, err
}

// Remove 删除员工
func (a *aStaff) Remove(ctx g.Ctx, req *vo.StaffRemoveReq) (res vo.StaffRemoveRes, err error) {
	return res, err
}

// SendChangeCode 发送转让验证码
func (a *aStaff) SendChangeCode(ctx g.Ctx, req *vo.SendChangeCodeReq) (res vo.SendChangeCodeRes, err error) {
	return res, err
}

// ChangeStaff 转让店铺
func (a *aStaff) ChangeStaff(ctx g.Ctx, req *vo.ChangeStaffReq) (res vo.ChangeStaffRes, err error) {
	return res, err
}
