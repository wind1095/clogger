package clogger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var myConsoleLogger *zap.Logger
var myFileLogger *zap.Logger

func Init(logpath string, appname string, logday time.Duration) {

	//logFile := "./var/log/app-%Y-%m-%d.log"
	logFile := logpath + "/" + appname + "/" + appname + "-%Y-%m-%d.log"
	linkFile := logpath + "/" + appname + ".log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithLinkName(linkFile),
		rotatelogs.WithMaxAge(logday),
		rotatelogs.WithRotationTime(24*time.Hour))

	if err != nil {
		panic(err)
	}
	//파일 로거 정의
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "timestamp",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	logCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(rotator),
		zap.InfoLevel)

	myFileLogger = zap.New(logCore)

	//콘솔 로거 정의
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	myConsoleLogger, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	myConsoleLogger.Info(message, fields...)
	myFileLogger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	myConsoleLogger.Debug(message, fields...)
	myFileLogger.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	myConsoleLogger.Warn(message, fields...)
	myFileLogger.Warn(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	myConsoleLogger.Panic(message, fields...)
	myFileLogger.Panic(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	myConsoleLogger.Error(message, fields...)
	myFileLogger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	myConsoleLogger.Fatal(message, fields...)
	myFileLogger.Fatal(message, fields...)
}
