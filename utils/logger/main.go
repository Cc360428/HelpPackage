package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var logger *zap.Logger

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 	  {"level":"debug","t":"20200919_222053","caller":"logger/main.go:50","msg":"sadf","app":"test"}
	fmt.Println()
	enc.AppendString(fmt.Sprintf("%d-%d-%d %d:%d:%d ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func InitZapLog() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "t",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     formatEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"utils/logger/save/zap.log"},
		ErrorOutputPaths: []string{"utils/logger/save/zap.log"},
		InitialFields: map[string]interface{}{
			"app": "test",
		},
	}
	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
}

func main() {
	InitZapLog()
	defer logger.Sync()
	logger.Debug("sada")
	logger.Info("sada")
	logger.Error("error")
}
