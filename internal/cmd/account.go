package cmd

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"shopkone-service/internal/api"
)

func registerAccountRoutes(s *ghttp.Server) {
	accountApi := api.NewAccountApi()
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Bind(accountApi)
	})
}
