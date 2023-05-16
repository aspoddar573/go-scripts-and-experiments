package logger

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

type LogConfig struct {
	coreLogger  *zap.Logger
	sugarLogger *zap.SugaredLogger
	host        string
}

var logConfig LogConfig

func getLevel(l LogLevel) (zapcore.Level, error) {
	switch l {
	case LogLevelDebug:
		return zap.DebugLevel, nil
	case LogLevelInfo:
		return zap.InfoLevel, nil
	case LogLevelWarn:
		return zap.WarnLevel, nil
	case LogLevelError:
		return zap.ErrorLevel, nil
	case LogLevelFatal:
		return zap.FatalLevel, nil
	default:
		return zapcore.PanicLevel, errors.New("unsupported level")
	}
}

func InitializeLogger(logLevel LogLevel, encoder zapcore.Encoder, zapWriteSyncers []zapcore.WriteSyncer, callerSkip int) *zap.Logger {
	host := os.Getenv("HOSTNAME")
	if len(host) == 0 {
		host = "default"
	}

	level, err := getLevel(logLevel)
	if err != nil {
		// overriding to debug
		level = zap.DebugLevel
	}

	core := zapcore.NewCore(encoder, zap.CombineWriteSyncers(zapWriteSyncers...), level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(callerSkip))
	logConfig = LogConfig{
		coreLogger:  logger,
		sugarLogger: logger.Sugar(),
		host:        host,
	}
	return logger
}

type ILogger interface {
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Audit(ctx context.Context, args ...interface{})
	Auditf(ctx context.Context, format string, args ...interface{})
	WithTag(key string, value interface{}) ILogger
	WithTags(tags map[string]interface{}) ILogger
	Sync() error
}
