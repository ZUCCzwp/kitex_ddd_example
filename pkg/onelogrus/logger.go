package onelogrus

import (
	"context"
	"github.com/ZUCCzwp/kitex_ddd_example/pkg/onelog"
	"github.com/sirupsen/logrus"
	"io"
)

var _ onelog.FullLogger = (*Logger)(nil)

type Logger struct {
	l *logrus.Logger
}

func NewLogger(opts ...Option) *Logger {
	cfg := defaultConfig()

	// apply options
	for _, opt := range opts {
		opt.apply(cfg)
	}

	// default trace hooks
	cfg.hooks = append(cfg.hooks, NewTraceHook(cfg.traceHookConfig))

	if cfg.timezoneLocation != nil {
		cfg.hooks = append(cfg.hooks, &TimezoneHook{Location: cfg.timezoneLocation})
	}
	if cfg.foreverLevelConfig != nil {
		cfg.hooks = append(cfg.hooks, NewForeverLevelHook(cfg.foreverLevelConfig))
	}
	// attach hook
	for _, hook := range cfg.hooks {
		cfg.logger.AddHook(hook)
	}

	return &Logger{
		l: cfg.logger,
	}
}

func (l *Logger) Logger() *logrus.Logger {
	return l.l
}

func (l *Logger) Trace(v ...interface{}) {
	l.l.Trace(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.l.Debug(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.l.Info(v...)
}

func (l *Logger) Forever(v ...interface{}) {
	l.l.Log(logrus.Level(onelog.LevelForever), v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.l.Error(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.l.Fatal(v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.l.Tracef(format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.l.Debugf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.l.Infof(format, v...)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.l.Errorf(format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.l.Fatalf(format, v...)
}

func (l *Logger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Tracef(format, v...)
}

func (l *Logger) CtxTraceOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Trace(v)
}

func (l *Logger) CtxTraceOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.WithContext(ctx).WithFields(mergedFields).Trace(v)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Debugf(format, v...)
}

func (l *Logger) CtxDebugOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Debug(v)
}

func (l *Logger) CtxDebugOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.WithContext(ctx).WithFields(mergedFields).Debug(v)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Infof(format, v...)
}

func (l *Logger) CtxInfoOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Info(v)
}

func (l *Logger) CtxInfoOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.WithContext(ctx).WithFields(mergedFields).Info(v)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Warnf(format, v...)
}

func (l *Logger) CtxNoticeOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Warn(v)
}

func (l *Logger) CtxNoticeOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.WithContext(ctx).WithFields(mergedFields).Warn(v)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Warnf(format, v...)
}

func (l *Logger) CtxWarnOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Warn(v)
}

func (l *Logger) CtxWarnOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.WithContext(ctx).WithFields(mergedFields).Warn(v)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Errorf(format, v...)
}

func (l *Logger) CtxErrorOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Error(v)
}

func (l *Logger) CtxErrorOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.WithContext(ctx).WithFields(mergedFields).Error(v)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Fatalf(format, v...)
}

func (l *Logger) CtxFatalOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Fatal(v)
}

func (l *Logger) CtxFatalOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.WithContext(ctx).WithFields(mergedFields).Fatal(v)
}

func (l *Logger) CtxForeverOption(ctx context.Context, v interface{}, key string, value interface{}) {
	l.l.WithContext(ctx).WithField(key, value).Info(v)
}

func (l *Logger) CtxForeverOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	mergedFields := make(logrus.Fields)

	for _, f := range fields {
		for key, value := range f {
			mergedFields[key] = value
		}
	}
	l.l.Log(logrus.InfoLevel)
	l.l.WithContext(ctx).WithFields(mergedFields).Info(v)
}

func (l *Logger) SetLevel(level onelog.Level) {
	var lv logrus.Level
	switch level {
	case onelog.LevelTrace:
		lv = logrus.TraceLevel
	case onelog.LevelDebug:
		lv = logrus.DebugLevel
	case onelog.LevelInfo:
		lv = logrus.InfoLevel
	case onelog.LevelWarn, onelog.LevelNotice:
		lv = logrus.WarnLevel
	case onelog.LevelError:
		lv = logrus.ErrorLevel
	case onelog.LevelFatal:
		lv = logrus.FatalLevel
	case onelog.LevelForever:
		lv = logrus.InfoLevel
	default:
		lv = logrus.WarnLevel
	}
	l.l.SetLevel(lv)
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.l.SetOutput(writer)
}
