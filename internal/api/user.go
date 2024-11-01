package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/shop/user/sUser"
	ctx2 "shopkone-service/utility/ctx"
)

type aUserApi struct {
}

func NewUserApi() *aUserApi {
	return &aUserApi{}
}

func (s *aUserApi) Info(ctx g.Ctx, req *vo.UserInfoReq) (res vo.UserInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	res.Email = auth.User.Email
	res.Language = auth.User.Language
	res.IsMaster = auth.Staff.IsMaster
	return res, err
}

// SetColumns 设置列
func (a *aUserApi) SetColumns(ctx g.Ctx, req *vo.SetColumnsReq) (res vo.SetColumnsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	user := auth.User
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		return sUser.NewUserColumn(tx).Set(user.ID, *req)
	})
	return res, err
}

// GetColumns 获取列
func (a *aUserApi) GetColumns(ctx g.Ctx, req *vo.GetColumnsReq) (res vo.GetColumnsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	user := auth.User
	res, err = sUser.NewUserColumn(sOrm.NewDb()).Get(user.ID, *req)
	return
}
