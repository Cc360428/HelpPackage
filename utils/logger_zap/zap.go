package logger_zap

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 	  {"level":"debug","t":"20200919_222053","caller":"logger_zap/zap.go:50","msg":"sadf","app":"test"}
	fmt.Println()
	enc.AppendString(fmt.Sprintf("%d-%d-%d %d:%d:%d ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func InitZapLog() (*zap.Logger, error) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "t",
			LevelKey:       "level",
			NameKey:        "logger_zap",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     formatEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"Cc360428_log_test"},
		ErrorOutputPaths: []string{"Cc360428_log_test"},
		InitialFields: map[string]interface{}{
			"app": "test",
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
	return logger, nil
}

//
//func main() {
//	InitZapLog()
//	defer logger.Sync()
//	logger.Debug("sada")
//	logger.Info("sada")
//	logger.Error("error")
//}
