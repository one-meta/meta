package wireset

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/one-meta/meta/app/entity/config"
)

func InitLog() (*zap.Logger, func(), error) {
	atomicLevel := zap.NewAtomicLevel()
	switch config.CFG.Log.App.LogLevel {
	case 1:
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case 2:
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case 3:
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case 4:
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	case 5:
		atomicLevel.SetLevel(zapcore.DPanicLevel)
	case 6:
		atomicLevel.SetLevel(zapcore.PanicLevel)
	case 7:
		atomicLevel.SetLevel(zapcore.FatalLevel)
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var lumberConfig = config.CFG.Log.App.Lumberjack
	// 日志轮转
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

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(writer),
		atomicLevel,
	)
	logger := zap.New(zapCore, zap.AddCaller())
	return logger, func() {
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil {

			}
		}(logger)
	}, nil
}
