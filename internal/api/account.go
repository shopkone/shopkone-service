package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/email/iEmail"
	"shopkone-service/internal/module/base/email/sEmail"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/shop/shop/iShop"
	"shopkone-service/internal/module/shop/shop/sShop"
	"shopkone-service/internal/module/shop/user/iUser"
	"shopkone-service/internal/module/shop/user/sUser"
	"shopkone-service/utility/code"
	ctx2 "shopkone-service/utility/ctx"
)

type aAccountApi struct {
}

func NewAccountApi() *aAccountApi {
	return &aAccountApi{}
}

// Login 登录
func (s *aAccountApi) Login(ctx g.Ctx, req *vo.LoginReq) (res vo.LoginRes, err error) {
	c := ctx2.NewCtx(ctx)
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		loginIn := iUser.LoginIn{
			Email: req.Email,
			Pwd:   req.Password,
			Ip:    c.GetIp(),
			Ua:    c.GetUa(),
		}
		res.Token, err = sUser.NewUser(tx).Login(loginIn)
		return err
	})
	return res, err
}

// Logout 退出登录
func (s *aAccountApi) Logout(ctx g.Ctx, req *vo.LogoutReq) (res vo.LogoutRes, err error) {
	return res, err
}

// Register 注册
func (s *aAccountApi) Register(ctx g.Ctx, req *vo.RegisterReq) (res vo.RegisterRes, err error) {
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		// 校验验证码
		if err = sEmail.NewCaptcha().Verify(req.Email, req.Code, iEmail.RegisterSense); err != nil {
			return err
		}
		// 注册用户
		user, err := sUser.NewUser(tx).Register(*req, ctx2.NewCtx(ctx).GetLanguage())
		if err != nil {
			return err
		}
		// 创建试用店铺
		createTrialIn := iShop.CreateTrialIn{
			Email:   req.Email,
			UserId:  user.ID,
			Country: "CN",
			Zone:    "",
		}
		if _, err = sShop.NewShop(tx).CreateTrial(createTrialIn); err != nil {
			return err
		}
		// 删除缓存
		return sEmail.NewCaptcha().Remove(req.Email, iEmail.RegisterSense)
	})
	return res, err
}

// ResetPwd 重置密码
func (s *aAccountApi) ResetPwd(ctx g.Ctx, req *vo.ResetPwdReq) (res vo.ResetPwdRes, err error) {
	return res, err
}

// AcceptInvite 接受邀请
func (s *aAccountApi) AcceptInvite(ctx g.Ctx, req *vo.AcceptInviteReq) (res vo.AcceptInviteRes, err error) {
	return res, err
}

// CheckInvite 校验邀请
func (s *aAccountApi) CheckInvite(ctx g.Ctx, req *vo.CheckInviteReq) (res vo.CheckInviteRes, err error) {
	return res, err
}

// SendCode 发送验证码
func (s *aAccountApi) SendCode(ctx g.Ctx, req *vo.SendCodeReq) (res vo.SendCodeRes, err error) {
	// 校验场景值
	if req.Type != iEmail.RegisterSense && req.Type != iEmail.ResetSense {
		return res, code.SenseError
	}
	// 发送验证码
	sendIn := iEmail.CaptchaSendIn{
		Email:        req.Email,
		Sense:        req.Type,
		Template:     "您的验证码为{{.Code}}, 10分钟内有效",
		MaxSendTimes: 20,
		MaxTryTimes:  20,
	}
	if req.Type == iEmail.RegisterSense {
		sendIn.Subject = "正在注册Shopkone账号"
	} else {
		sendIn.Subject = "正在重置Shopkone密码"
	}
	if err = sEmail.NewCaptcha().Send(sendIn); err != nil {
		return res, err
	}
	return res, err
}
