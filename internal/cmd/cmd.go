package cmd

import (
	"context"
	"shopkone-service/hack"
	"shopkone-service/internal/module/base/ali/sAli"
	"shopkone-service/internal/module/base/cache/sCache"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 获取配置
			conf, err := hack.GetConfig()
			if err != nil {
				return err
			}
			// 初始化基础模块
			if err = bootstrap(); err != nil {
				return err
			}
			// 初始化阿里云服务
			if err := sAli.AliYunClient(); err != nil {
				return err
			}
			// 初始化服务
			s := g.Server()
			// 加载通用中间件
			s.BindMiddlewareDefault(middleware.Compress)
			s.BindMiddlewareDefault(middleware.RequestLogger)
			// 注册account路由
			registerAccountRoutes(s)
			// 注册admin路由
			registerAdminRoutes(s)
			// 设置服务地址
			s.SetAddr(conf.Admin.Address)
			// 启动服务
			s.Run()
			return nil
		},
	}
)

func bootstrap() error {
	// 连接redis
	if err := sCache.NewCacheConnect().Connect(); err != nil {
		return err
	}
	// 连接mysql
	if err := sOrm.ConnectMysql(migrate); err != nil {
		return err
	}
	// 初始化资源
	if err := resource.InitResource(); err != nil {
		return err
	}
	return nil
}
