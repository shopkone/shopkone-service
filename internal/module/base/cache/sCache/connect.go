package sCache

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"shopkone-service/hack"
)

var email *gcache.Cache
var auth *gcache.Cache
var user *gcache.Cache
var staff *gcache.Cache
var shop *gcache.Cache

type ICacheConnect interface {
	Connect() error                      // 连接Redis
	setCache(int) (*gcache.Cache, error) // 设置缓存
}

type sCacheConnect struct {
}

func NewCacheConnect() *sCacheConnect {
	return &sCacheConnect{}
}

func (s *sCacheConnect) Connect() (err error) {
	// 邮件缓存
	if email, err = s.setCache(0); err != nil {
		return err
	}
	// 鉴权缓存
	if auth, err = s.setCache(1); err != nil {
		return err
	}
	// 用户缓存
	if user, err = s.setCache(2); err != nil {
		return err
	}
	// 员工缓存
	if staff, err = s.setCache(3); err != nil {
		return err
	}
	// 店铺缓存
	if shop, err = s.setCache(4); err != nil {
		return err
	}
	return nil
}

func (s *sCacheConnect) setCache(db int) (c *gcache.Cache, err error) {
	// 获取配置
	conf, err := hack.GetConfig()
	if err != nil {
		return c, err
	}
	// 链接redis
	address := conf.Redis.Host + ":" + conf.Redis.Port
	redis, err := gredis.New(&gredis.Config{Address: address, Db: db})
	if err != nil {
		return c, err
	}
	// 设置适配器
	c = gcache.New()
	c.SetAdapter(gcache.NewAdapterRedis(redis))
	return c, nil
}
