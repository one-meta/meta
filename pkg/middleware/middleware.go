package middleware

import (
	"os"

	"github.com/one-meta/meta/app/entity/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Middleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func Middleware(m *fiber.App) {
	logConfig := GetWebLogConfig()
	m.Use(
		SetSecurityHeader(),
		cors.New(),
		logger.New(
			*logConfig,
		),
		requestid.New(),
		compress.New(),
		//本地测试，如果空指针导致panic，可以取消recover中间件，就能定位到具体位置了
		recover.New(),
	)
}

func GetWebLogConfig() *logger.Config {
	logConfig := &logger.Config{}
	var webConfig = config.CFG.Log.Web
	var lumberConfig = config.CFG.Log.Web.Lumberjack
	if webConfig.Format != "" {
		logConfig.Format = webConfig.Format
	}
	if webConfig.TimeFormat != "" {
		logConfig.TimeFormat = webConfig.TimeFormat
	}
	if webConfig.TimeZone != "" {
		logConfig.TimeZone = webConfig.TimeZone
	}
	if webConfig.Output == "stderr" || webConfig.Output == "" {
		logConfig.Output = os.Stderr
	} else {
		// 使用lumberjack进行日志滚动
		writer := &lumberjack.Logger{
			// 日志名称
			Filename: lumberConfig.LogFile,
			// 日志大小限制，单位MB
			MaxSize: lumberConfig.MaxSize,
			// 历史日志文件保留天数
			MaxAge: lumberConfig.MaxAge,
			// 最大保留历史日志数量
			MaxBackups: lumberConfig.MaxBackup,
			// 本地时区
			LocalTime: lumberConfig.LocalTime,
			// 历史日志文件压缩
			Compress: lumberConfig.Compress,
		}
		logConfig.Output = writer
	}
	return logConfig
}
