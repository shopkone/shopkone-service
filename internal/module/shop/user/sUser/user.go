package sUser

import (
	"github.com/gogf/gf/v2/crypto/gsha1"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/shop/user/iUser"
	"shopkone-service/internal/module/shop/user/mUser"
	"shopkone-service/utility/code"
)

type IUser interface {
	PreRegister() (user mUser.User, err error)                                // 预注册用户
	Register(in vo.RegisterReq, language string) (user mUser.User, err error) // 注册
	Login(email, pwd string) (token string, err error)                        // 登录
	Logout(token string) (err error)                                          // 登出
	ResetPwd() (err error)                                                    // 重置密码
	InfoById(id uint) (user mUser.User, err error)                            // 根据id获取用户信息
	InfoByEmail(email string) (user mUser.User, err error)                    // 根据邮箱获取用户信息
	EncryptPwd(pwd string) string                                             // 加密密码
	UpdateById(id uint, query ...string) (user mUser.User, err error)         // 根据id更新用户信息
}

type sUser struct {
	orm *gorm.DB
}

func NewUser(orm *gorm.DB) *sUser {
	return &sUser{orm: orm}
}

func (s *sUser) PreRegister() (user mUser.User, err error) {
	return user, err
}

func (s *sUser) Register(in vo.RegisterReq, language string) (user mUser.User, err error) {
	// 获取用户信息
	if user, err = s.InfoByEmail(in.Email); err != nil {
		return user, err
	}
	// 判断用户是否已经注册
	if user.Status == mUser.UserStatusRegistered {
		return user, code.UserIsRegistered
	}
	// 组装数据
	user.Password = s.EncryptPwd(in.Password)
	user.Language = language
	user.Email = in.Email
	// 如果是预注册用户，则更新
	if user.Status == mUser.UserStatusPreRegister {
		user.Status = mUser.UserStatusRegistered
		user, err = s.UpdateById(user.ID, "status", "language", "password")
		if err != nil {
			return user, err
		}
	} else {
		// 否则创建用户
		user.Status = mUser.UserStatusRegistered
		if err = s.orm.Model(&user).Create(&user).Error; err != nil {
			return user, err
		}
	}
	// 创建表格列
	if err = NewUserColumn(s.orm).Create(user.ID); err != nil {
		return user, err
	}
	// 更新缓存
	if err = NewUserCache().UpdateUserCache(user.ID, s.orm); err != nil {
		return user, err
	}
	return user, err
}

func (s *sUser) Login(in iUser.LoginIn) (token string, err error) {
	user, err := s.InfoByEmail(in.Email)
	if err != nil {
		return "", err
	}
	// 是否存在
	if user.ID == 0 {
		return "", code.UserUnRegistered
	}
	// 校验密码
	if user.Password != s.EncryptPwd(in.Pwd) {
		return "", code.UserPwdError
	}
	// 是否被冻结
	if user.Status == mUser.UserStatusLocked {
		return "", code.UserIsBlocked
	}
	// 校验状态
	if user.Status != mUser.UserStatusRegistered {
		return "", code.UserUnRegistered
	}
	// 添加记录
	loginRecordIn := iUser.LoginRecordIn{
		UserId: user.ID,
		Ip:     in.Ip,
		Ua:     in.Ua,
	}
	token, err = NewUserLoginRecord(s.orm).AddRecord(loginRecordIn)
	if err != nil {
		return "", err
	}
	return token, err
}

func (s *sUser) Logout(token string) (err error) {
	return err
}

func (s *sUser) ResetPwd() (err error) {
	return err
}

func (s *sUser) InfoById(id uint) (user mUser.User, err error) {
	if err = s.orm.Model(&user).Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, err
}

func (s *sUser) InfoByEmail(email string) (user mUser.User, err error) {
	query := s.orm.Model(&user).Where("email = ?", email)
	err = query.Select("status", "id", "password").Find(&user).Error
	return user, err
}

func (s *sUser) EncryptPwd(pwd string) string {
	return gsha1.Encrypt(pwd)
}

func (s *sUser) UpdateById(id uint, query ...string) (user mUser.User, err error) {
	if err = s.orm.Model(&user).Where("id = ?", id).
		Select(query).Updates(&user).Error; err != nil {
		return user, err
	}
	return user, err
}
