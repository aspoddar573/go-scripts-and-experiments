package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
}

func getLogrusLevel(level LogLevel) logrus.Level {
	logrusLevelMapper := map[LogLevel]logrus.Level{
		DEBUG: logrus.DebugLevel,
		INFO:  logrus.InfoLevel,
		ERROR: logrus.ErrorLevel,
		WARN:  logrus.WarnLevel,
		FATAL: logrus.FatalLevel,
	}
	if logrusLevel, ok := logrusLevelMapper[level]; !ok {
		return logrus.InfoLevel
	} else {
		return logrusLevel
	}

}

func newLogrusLogger(config Configuration) (Logger, error) {

	level := getLogrusLevel(config.LogLevel)
	stdOutHandler := os.Stdout
	fileHandler := &lumberjack.Logger{
		Filename: config.FileLocation,
		MaxSize:  100,
		Compress: true,
		MaxAge:   28,
	}
	lLogger := &logrus.Logger{
		Out:       stdOutHandler,
		Formatter: getFormatter(config.EnableJSON),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	if config.Output == CONSOLE {
		lLogger.SetOutput(io.MultiWriter(stdOutHandler))
	} else {
		if config.Output == FILE {
			lLogger.SetOutput(fileHandler)
			lLogger.SetFormatter(getFormatter(config.EnableJSON))
		}
	}

	return &logrusLogger{
		logger: lLogger,
	}, nil
}

func (l *logrusLogger) Debugf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Debug(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Debug(args...)
}

func (l *logrusLogger) Infof(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Info(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Info(args...)
}

func (l *logrusLogger) Warnf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Warn(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Warn(args...)
}

func (l *logrusLogger) Errorf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Error(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Error(args...)
}

func (l *logrusLogger) Fatalf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Fatal(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogger)
	l.logger.Fatal(args...)
}

func (l *logrusLogger) WithFields(facet Facet) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(facet)),
	}
}

func (l *logrusLogEntry) Debugf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Debugf(format, args...)
}

func (l *logrusLogEntry) Debug(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Debug(args...)
}

func (l *logrusLogEntry) Infof(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Infof(format, args...)
}

func (l *logrusLogEntry) Info(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Info(args...)
}

func (l *logrusLogEntry) Warnf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Warnf(format, args...)
}

func (l *logrusLogEntry) Warn(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Warn(args...)
}

func (l *logrusLogEntry) Errorf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Errorf(format, args...)
}

func (l *logrusLogEntry) Error(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Error(args...)
}

func (l *logrusLogEntry) Fatalf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogEntry) Fatal(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*logrusLogEntry)
	l.entry.Fatal(args...)
}

func (l *logrusLogEntry) WithFields(facet Facet) Logger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(facet)),
	}
}

func convertToLogrusFields(facet Facet) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range facet.GetFields() {
		logrusFields[index] = val
	}
	return logrusFields
}
