package Logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
)

var SysLog = New("Engine", zapcore.InfoLevel)

func New(service string, level zapcore.Level) *zap.SugaredLogger {
	syncWriter := getLoggerSync(service)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, syncWriter, level)
	logger := zap.New(core)
	sugaredLogger := logger.Sugar()
	defer sugaredLogger.Sync()
	return sugaredLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLoggerSync(service string) zapcore.WriteSyncer {
	var ws []zapcore.WriteSyncer
	var logRoot string
	if runtime.GOOS == "linux" {
		logRoot = "/var/log"
	} else {
		ws = append(ws, zapcore.AddSync(os.Stdout))
		logRoot = "./log"
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", logRoot, service),
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	ws = append(ws, zapcore.AddSync(lumberJackLogger))
	return zapcore.NewMultiWriteSyncer(ws...)
}
