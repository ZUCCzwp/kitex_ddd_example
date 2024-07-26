package onelogrus

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Option interface {
	apply(cfg *config)
}

type option func(cfg *config)

func (fn option) apply(cfg *config) {
	fn(cfg)
}

type config struct {
	logger *logrus.Logger
	hooks  []logrus.Hook

	traceHookConfig    *TraceHookConfig
	timezoneLocation   *time.Location
	foreverLevelConfig *ForeverLevelConfig
}

func defaultConfig() *config {
	// new logger
	logger := logrus.New()
	// default json format
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC1123Z,
	})

	return &config{
		logger: logger,
		hooks:  []logrus.Hook{},
		traceHookConfig: &TraceHookConfig{
			recordStackTraceInSpan: true,
			enableLevels:           logrus.AllLevels,
			errorSpanLevel:         logrus.ErrorLevel,
		},
	}
}

func WithLogger(logger *logrus.Logger) Option {
	return option(func(cfg *config) {
		cfg.logger = logger
	})
}

func WithHook(hook logrus.Hook) Option {
	return option(func(cfg *config) {
		cfg.hooks = append(cfg.hooks, hook)
	})
}

func WithTraceHookConfig(hookConfig *TraceHookConfig) Option {
	return option(func(cfg *config) {
		cfg.traceHookConfig = hookConfig
	})
}

func WithTraceHookLevels(levels []logrus.Level) Option {
	return option(func(cfg *config) {
		cfg.traceHookConfig.enableLevels = levels
	})
}

func WithTraceHookErrorSpanLevel(level logrus.Level) Option {
	return option(func(cfg *config) {
		cfg.traceHookConfig.errorSpanLevel = level
	})
}

func WithRecordStackTraceInSpan(recordStackTraceInSpan bool) Option {
	return option(func(cfg *config) {
		cfg.traceHookConfig.recordStackTraceInSpan = recordStackTraceInSpan
	})
}

func WithTimezone(location *time.Location) Option {
	return option(func(cfg *config) {
		cfg.timezoneLocation = location
	})
}

func WithForeverLevel(minLevel ForeverLevel) Option {
	return option(func(cfg *config) {
		if cfg.foreverLevelConfig == nil {
			cfg.foreverLevelConfig = &ForeverLevelConfig{}
		}
		cfg.foreverLevelConfig.MinLevel = minLevel
	})
}
