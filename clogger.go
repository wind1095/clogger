package clogger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var myConsoleLogger *zap.SugaredLogger
var myFileLogger *zap.SugaredLogger

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
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(rotator),
		zap.InfoLevel)

//	myFileLogger = zap.New(logCore)
	log := zap.New(logCore)
	myFileLogger = log.Sugar()
	myFileLogger = myFileLogger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))

	//콘솔 로거 정의
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000")
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log2, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	myConsoleLogger = log2.Sugar()
			
}
func Info(template string, args ...interface{}) {
	myConsoleLogger.Infof(template, args...)
	myFileLogger.Infof(template, args...)
}

func Debug(template string, args ...interface{}) {
	myConsoleLogger.Debugf(template, args...)
	myFileLogger.Debugf(template, args...)
}

func Warn(template string, args ...interface{}) {
	myConsoleLogger.Warnf(template, args...)
	myFileLogger.Warnf(template, args...)
}

func Panic(template string, args ...interface{}) {
	myConsoleLogger.Panicf(template, args...)
	myFileLogger.Panicf(template, args...)
}

func Error(template string, args ...interface{}) {
	myConsoleLogger.Errorf(template, args...)
	myFileLogger.Errorf(template, args...)
}

func Infoln(args ...interface{}) {
	myConsoleLogger.Infoln(args...)
	myFileLogger.Infoln(args...)
}

func Debugln(args ...interface{}) {
	myConsoleLogger.Debugln(args...)
	myFileLogger.Debugln(args...)
}

func Warnln(args ...interface{}) {
	myConsoleLogger.Warnln(args...)
	myFileLogger.Warnln(args...)
}

func Panicln(args ...interface{}) {
	myConsoleLogger.Panicln(args...)
	myFileLogger.Panicln(args...)
}

func Errorln(args ...interface{}) {
	myConsoleLogger.Errorln(args...)
	myFileLogger.Errorln(args...)
}

