package main

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func initLogger() {
	cfg := zap.NewProductionConfig()
	cfg.Encoding = "console"
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeTime = customTimeEncoder
	cfg.EncoderConfig.EncodeLevel = customLevelEncoder

	tmp, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	logger = tmp.Sugar()
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(time.RFC3339) + "]")
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
