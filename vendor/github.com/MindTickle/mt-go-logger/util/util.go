package util

import (
	"github.com/MindTickle/mt-go-logger/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetDefaultTags() map[string]struct{} {
	tagSet := make(map[string]struct{}, 0)
	tagSet[constants.ServiceReqIdTag] = struct{}{}
	for _, tag := range constants.TracingTags {
		tagSet[tag] = struct{}{}
	}
	return tagSet
}

func GetDefaultConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetDefaultJsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func GetDefaultWriteSyncers() []zapcore.WriteSyncer {
	zapWriteSyncers := make([]zapcore.WriteSyncer, 0)
	zapWriteSyncers = append(zapWriteSyncers, zapcore.AddSync(os.Stdout))
	return zapWriteSyncers
}
