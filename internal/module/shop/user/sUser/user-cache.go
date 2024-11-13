package sUser

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"gorm.io/gorm"
	"shopkone-service/internal/module/base/cache/sCache"
	"shopkone-service/internal/module/shop/user/mUser"
	"shopkone-service/utility/code"
)

type IUserCache interface {
	// UpdateUserCache 更新用户缓存
	UpdateUserCache(userId uint, orm *gorm.DB) error
	// GetUserCache 获取用户缓存
	GetUserCache(userId uint) (user mUser.User, err error)
	// UpdateUserLoginCache 更新用户登录缓存
	UpdateUserLoginCache(token string, userId uint) (err error)
	// GetUserLoginCache 获取用户登录缓存
	GetUserLoginCache(token string) (user mUser.User, err error)
}

type sUserCache struct {
}

func NewUserCache() *sUserCache {
	return &sUserCache{}
}

func (s *sUserCache) UpdateUserCache(userId uint, orm *gorm.DB) error {
	key := sCache.USER_PREFIX_KEY + convertor.ToString(userId)
	user, err := NewUser(orm).InfoById(userId)
	if err != nil {
		return err
	}
	year100 := 60 * 24 * 365 * 100 //100年
	if err = sCache.NewUserCache().Set(key, user, uint(year100)); err != nil {
		return err
	}
	return nil
}

func (s *sUserCache) GetUserCache(userId uint) (user mUser.User, err error) {
	key := sCache.USER_PREFIX_KEY + convertor.ToString(userId)
	err = sCache.NewUserCache().Get(key, &user)
	return user, err
}

func (s *sUserCache) UpdateUserLoginCache(token string, userId uint) (err error) {
	// 7天
	day7 := 60 * 24 * 7
	// 使用md5加密，避免token暴露
	token, err = gmd5.Encrypt(token)
	if err != nil {
		return err
	}
	// 查询是否冲突，冲突直接不允许登录
	user, err := s.GetUserLoginCache(token)
	if err != nil {
		return err
	}
	if user.ID != 0 {
		return code.LoginSystemLock
	}
	if err = sCache.AuthCache().Set(sCache.TOKEN_PREFIX_KEY+token, userId, uint(day7)); err != nil {
		return err
	}
	return nil
}

func (s *sUserCache) GetUserLoginCache(token string) (user mUser.User, err error) {
	// md5加密
	token, err = gmd5.Encrypt(token)
	if err != nil {
		return user, err
	}
	var userId uint
	if err = sCache.AuthCache().Get(sCache.TOKEN_PREFIX_KEY+token, &userId); err != nil {
		return user, err
	}
	return s.GetUserCache(userId)
}

func (s *sUserCache) Removes(tokens []string) error {
	if len(tokens) == 0 {
		return nil
	}
	for _, token := range tokens {
		if err := sCache.AuthCache().Remove(sCache.TOKEN_PREFIX_KEY + token); err != nil {
			return err
		}
	}
	return nil
}
