package bootstrap

import (
	"ginserver/global"
	"ginserver/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level   zapcore.Level
	options []zap.Option
)

func InitializeLog() *zap.Logger {
	createRootDir()
	setLogLevel()
	if global.AppConfig.Config.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	return zap.New(getZapCore(), options...)
}

func createRootDir() {
	if ok, _ := utils.PathExists(global.AppConfig.Config.Log.RootDir); !ok {
		_ = os.Mkdir(global.AppConfig.Config.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.AppConfig.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.AppConfig.Config.App.Env + "." + l.String())
	}

	// 设置编码器
	if global.AppConfig.Config.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.AppConfig.Config.Log.RootDir + "/" + global.AppConfig.Config.Log.Filename,
		MaxSize:    global.AppConfig.Config.Log.MaxSize,
		MaxBackups: global.AppConfig.Config.Log.MaxBackups,
		MaxAge:     global.AppConfig.Config.Log.MaxAge,
		Compress:   global.AppConfig.Config.Log.Compress,
	}

	return zapcore.AddSync(file)
}
