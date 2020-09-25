package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "",
		EncoderConfig:     zapcore.EncoderConfig{},
		OutputPaths:       nil,
		ErrorOutputPaths:  nil,
		InitialFields:     nil,
	}
	fmt.Println(config)
}
