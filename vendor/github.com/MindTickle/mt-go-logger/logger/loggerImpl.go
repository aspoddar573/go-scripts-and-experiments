package logger

import (
	"context"
	"fmt"
	"github.com/MindTickle/mt-go-logger/constants"
	"github.com/MindTickle/mt-go-logger/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type LoggerConfig struct {
	IncludeNilTagValues bool
	MaxLogSize          int
	TruncateMsg         bool
	LogLevel            LogLevel
	Encoder             zapcore.Encoder
	WriteSyncers        []zapcore.WriteSyncer
	CallerSkip          int
}

type LoggerImpl struct {
	logger *zap.Logger
	tags   map[string]struct{}
	config *LoggerConfig
}

var Logger LoggerImpl

func (l *LoggerImpl) UpdateConfig(config *LoggerConfig) *LoggerImpl {
	l.config = config
	if l.config.MaxLogSize == 0 {
		l.config.MaxLogSize = constants.DefaultLogMsgSize
	}
	if l.config.Encoder == nil {
		l.config.Encoder = util.GetDefaultJsonEncoder()
	}
	if l.config.WriteSyncers == nil {
		l.config.WriteSyncers = util.GetDefaultWriteSyncers()
	}
	if l.config.CallerSkip == 0 {
		l.config.CallerSkip = 1
	}
	l.logger = InitializeLogger(l.config.LogLevel, l.config.Encoder, l.config.WriteSyncers, l.config.CallerSkip)
	return l
}

func (l *LoggerImpl) AddTag(tag string) *LoggerImpl {
	l.tags[tag] = struct{}{}
	return l
}

func (l *LoggerImpl) AddTags(tagsList []string) *LoggerImpl {
	for _, tag := range tagsList {
		l.tags[tag] = struct{}{}
	}
	return l
}

func (l *LoggerImpl) RemoveTag(tag string) *LoggerImpl {
	if _, ok := l.tags[tag]; ok {
		delete(l.tags, tag)
	}
	return l
}

func (l *LoggerImpl) RemoveTags(tagsList []string) *LoggerImpl {
	for _, tag := range tagsList {
		if _, ok := l.tags[tag]; ok {
			delete(l.tags, tag)
		}
	}
	return l
}

func (l *LoggerImpl) Debug(ctx context.Context, args ...interface{}) {
	msg := l.getMessage("", args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Debug(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Debug(msg)
	}
}

func (l *LoggerImpl) Debugf(ctx context.Context, format string, args ...interface{}) {
	msg := l.getMessage(format, args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Debug(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Debug(msg)
	}
}

func (l *LoggerImpl) Info(ctx context.Context, args ...interface{}) {
	msg := l.getMessage("", args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Info(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Info(msg)
	}
}

func (l *LoggerImpl) Infof(ctx context.Context, format string, args ...interface{}) {
	msg := l.getMessage(format, args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Info(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Info(msg)
	}
}

func (l *LoggerImpl) Fatal(ctx context.Context, args ...interface{}) {
	msg := l.getMessage("", args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Fatal(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Fatal(msg)
	}
}

func (l *LoggerImpl) Fatalf(ctx context.Context, format string, args ...interface{}) {
	msg := l.getMessage(format, args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Fatal(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Fatal(msg)
	}
}

func (l *LoggerImpl) Warn(ctx context.Context, args ...interface{}) {
	msg := l.getMessage("", args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Warn(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Warn(msg)
	}
}

func (l *LoggerImpl) Warnf(ctx context.Context, format string, args ...interface{}) {
	msg := l.getMessage(format, args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Warn(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Warn(msg)
	}
}

func (l *LoggerImpl) Error(ctx context.Context, args ...interface{}) {
	msg := l.getMessage("", args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Error(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Error(msg)
	}
}

func (l *LoggerImpl) Errorf(ctx context.Context, format string, args ...interface{}) {
	msg := l.getMessage(format, args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).Error(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).Error(msg)
	}
}

func (l *LoggerImpl) Audit(ctx context.Context, args ...interface{}) {
	msg := l.getMessage("", args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).With(zap.String(constants.Audit, "true")).Info(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).With(zap.String(constants.Audit, "true")).Info(msg)
	}
}

func (l *LoggerImpl) Auditf(ctx context.Context, format string, args ...interface{}) {
	msg := l.getMessage(format, args...)
	if !l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		chunks := l.getMessageChunks(msg)
		for _, chunk := range chunks {
			l.getTaggedLogger(ctx).With(zap.String(constants.Audit, "true")).Info(chunk)
		}
	} else {
		l.getTaggedLogger(ctx).With(zap.String(constants.Audit, "true")).Info(msg)
	}
}

func (l *LoggerImpl) WithTag(key string, value interface{}) ILogger {
	return &LoggerImpl{logger: l.logger.With(zap.Any(key, value)), tags: l.tags, config: l.config}
}

func (l *LoggerImpl) WithTags(tags map[string]interface{}) ILogger {
	taggedLogger := l.logger
	for key, val := range tags {
		taggedLogger = taggedLogger.With(zap.Any(key, val))
	}
	return &LoggerImpl{logger: taggedLogger, tags: l.tags, config: l.config}
}

func (l *LoggerImpl) Sync() error {
	return l.logger.Sync()
}

func init() {
	Logger = *NewLogger()
}

func NewLogger() *LoggerImpl {
	encoder := util.GetDefaultJsonEncoder()
	writeSyncers := util.GetDefaultWriteSyncers()
	newLogger := LoggerImpl{
		logger: InitializeLogger(LogLevelDebug, encoder, writeSyncers, 1),
		tags:   util.GetDefaultTags(),
		config: &LoggerConfig{
			IncludeNilTagValues: false,
			MaxLogSize:          constants.DefaultLogMsgSize,
			TruncateMsg:         false,
			LogLevel:            LogLevelDebug,
			Encoder:             encoder,
			WriteSyncers:        writeSyncers,
			CallerSkip:          1,
		},
	}
	return &newLogger
}

func (l *LoggerImpl) getTaggedLogger(ctx context.Context) *zap.Logger {
	taggedLogger := l.logger
	if ctx == nil {
		return taggedLogger
	}
	for tag := range l.tags {
		if l.config.IncludeNilTagValues || ctx.Value(tag) != nil {
			taggedLogger = taggedLogger.With(zap.Any(tag, ctx.Value(tag)))
		}
	}
	return taggedLogger
}

func (l *LoggerImpl) getMessage(format string, args ...interface{}) string {
	msg := format
	if len(msg) == 0 && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if len(msg) > 0 && len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	if l.config.TruncateMsg && len(msg) > l.config.MaxLogSize {
		var sb strings.Builder
		sb.WriteString(msg[:l.config.MaxLogSize])
		sb.WriteString(" ...[truncated]")
		msg = sb.String()
	}
	return msg
}

func (l *LoggerImpl) getMessageChunks(msg string) []string {
	msgLength := len(msg)
	var splitMsg []string
	for start := 0; start < msgLength; start += l.config.MaxLogSize {
		end := start + l.config.MaxLogSize
		if end > msgLength {
			end = msgLength
		}
		splitMsg = append(splitMsg, msg[start:end])
	}
	return splitMsg
}
