package middleware

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

// ANSI 颜色代码
const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m" // 用于分隔符和成功状态码
	colorYellow = "\033[33m" // 用于重定向状态码
	colorRed    = "\033[31m" // 用于错误状态码
	colorCyan   = "\033[36m" // 用于路径和其他信息
)

// RequestLogger 日志中间件
func RequestLogger(r *ghttp.Request) {
	start := time.Now()

	// 执行下一个中间件或处理程序
	r.Middleware.Next()

	// 计算请求处理时间
	duration := time.Since(start)
	requestTime := start.Format("2006-01-02 15:04:05") // 请求时间格式化

	// 根据状态码选择颜色
	statusColor := colorGreen
	if r.Response.Status >= 300 && r.Response.Status < 400 {
		statusColor = colorYellow
	} else if r.Response.Status >= 400 {
		statusColor = colorRed
	}

	// 格式化并输出带颜色的日志
	log := fmt.Sprintf(
		"\n%s===============================================================================%s\n"+
			"Time: %s%s%s | Path: %s%s%s | Status: %s%d%s | Duration: %v\n"+
			"%s===============================================================================%s\n",
		colorGreen, colorReset,
		colorCyan, requestTime, colorReset,
		colorCyan, r.URL.Path, colorReset,
		statusColor, r.Response.Status, colorReset,
		duration,
		colorGreen, colorReset,
	)

	fmt.Println(log)
}
