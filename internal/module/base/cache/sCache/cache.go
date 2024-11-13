package sCache

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"shopkone-service/utility/code"
	"strings"
	"time"
)

type ICache interface {
	RemoveByKeys(keys []string) error                  // 删除多个缓存
	Set(key string, value interface{}, min uint) error // 更新缓存
	Remove(key string) error                           // 删除缓存
	Get(key string, out interface{}) error             // 获取缓存
}

type sCache struct {
	cache *gcache.Cache
}

func NewEmailCache() *sCache {
	return &sCache{email}
}
func AuthCache() *sCache {
	return &sCache{auth}
}
func NewUserCache() *sCache {
	return &sCache{user}
}
func NewShopCache() *sCache {
	return &sCache{shop}
}

func NewCache() *sCache {
	return &sCache{other}
}

func (s *sCache) RemoveByKeys(keys []string) error {
	if err := s.checkKeyPrefix(keys); err != nil {
		return err
	}
	return s.cache.Removes(context.TODO(), gconv.Interfaces(keys))
}

func (s *sCache) Set(key string, value interface{}, min uint) error {
	if err := s.checkKeyPrefix([]string{key}); err != nil {
		return err
	}
	t := time.Minute * time.Duration(min)
	return s.cache.Set(context.TODO(), key, value, t)
}

func (s *sCache) Remove(key string) error {
	if err := s.checkKeyPrefix([]string{key}); err != nil {
		return err
	}
	_, err := s.cache.Remove(context.TODO(), key)
	return err
}

func (s *sCache) Get(key string, out interface{}) error {
	if err := s.checkKeyPrefix([]string{key}); err != nil {
		return err
	}
	ret, err := s.cache.Get(context.TODO(), key)
	if err != nil {
		return err
	}
	if ret.Scan(&out) != nil {
		return err
	}
	return err
}

func (s *sCache) checkKeyPrefix(keys []string) error {
	for _, key := range keys {
		// 根据_进行分割
		prefix := strings.Split(key, "_")[0]
		if _, exists := allowedPrefixes[prefix+"_"]; !exists {
			g.Dump(prefix, "___")
			return code.CacheKeyError
		}
	}
	return nil
}
