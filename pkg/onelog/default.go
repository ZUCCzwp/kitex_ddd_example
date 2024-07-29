package onelog

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

//type LogOption func(*logOptions)

type LogOptions struct {
	fields map[string]interface{}
}

var logger FullLogger = &defaultLogger{
	level:  LevelInfo,
	stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
}

// SetOutput sets the output of default logger. By default, it is stderr.
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

// SetLevel sets the level of logs below which logs will not be output.
// The default log level is LevelTrace.
// Note that this method is not concurrent-safe.
func SetLevel(lv Level) {
	logger.SetLevel(lv)
}

// DefaultLogger return the default logger for kitex.
func DefaultLogger() FullLogger {
	return logger
}

// SetLogger sets the default logger.
// Note that this method is not concurrent-safe and must not be called
// after the use of DefaultLogger and global functions in this package.
func SetLogger(v FullLogger) {
	logger = v
}

// Fatal calls the default logger's Fatal method and then os.Exit(1).
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Error calls the default logger's Error method.
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Warn calls the default logger's Warn method.
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Notice calls the default logger's Notice method.
func Notice(v ...interface{}) {
	logger.Notice(v...)
}

// Info calls the default logger's Info method.
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Forever calls the default logger's Info method.
func Forever(v ...interface{}) {
	logger.Forever(v)
}

// Debug calls the default logger's Debug method
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Trace calls the default logger's Trace method.
func Trace(v ...interface{}) {
	logger.Trace(v...)
}

// Fatalf calls the default logger's Fatalf method and then os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

// Errorf calls the default logger's Errorf method.
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Warnf calls the default logger's Warnf method.
func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Noticef calls the default logger's Noticef method.
func Noticef(format string, v ...interface{}) {
	logger.Noticef(format, v...)
}

// Infof calls the default logger's Infof method.
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Debugf calls the default logger's Debugf method.
func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Tracef calls the default logger's Tracef method.
func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

// CtxFatalf calls the default logger's CtxFatalf method and then os.Exit(1).
func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxFatalf(ctx, format, v...)
}

func CtxFatalOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxFatalOption(ctx, v, key, value)
}

func CtxFatalOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxFatalOptions(ctx, v, fields...)
}

// CtxErrorf calls the default logger's CtxErrorf method.
func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxErrorf(ctx, format, v...)
}

func CtxErrorOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxErrorOption(ctx, v, key, value)
}

func CtxErrorOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxErrorOptions(ctx, v, fields...)
}

// CtxWarnf calls the default logger's CtxWarnf method.
func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxWarnf(ctx, format, v...)
}

func CtxWarnOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxWarnOption(ctx, v, key, value)
}

func CtxWarnOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxWarnOptions(ctx, v, fields...)
}

// CtxNoticef calls the default logger's CtxNoticef method.
func CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	logger.CtxNoticef(ctx, format, v...)
}

func CtxNoticeOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxNoticeOption(ctx, v, key, value)
}

func CtxNoticeOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxNoticeOptions(ctx, v, fields...)
}

// CtxInfof calls the default logger's CtxInfof method.
func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	logger.CtxInfof(ctx, format, v...)
}

func CtxInfoOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxInfoOption(ctx, v, key, value)
}

func CtxInfoOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxInfoOptions(ctx, v, fields...)
}

// CtxDebugf calls the default logger's CtxDebugf method.
//func CtxDebugf(ctx context.Context, opts []LogOption, format string, v ...interfaces{}) {
//	logger.CtxDebugf(ctx, opts, format, v...)
//}

func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxDebugf(ctx, format, v...)
}

func CtxDebugOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxDebugOption(ctx, v, key, value)
}

func CtxDebugOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxDebugOptions(ctx, v, fields...)
}

func WithCustomOption(customKey string, customValue interface{}) logrus.Fields {
	return logrus.Fields{
		customKey: customValue,
	}
}

// CtxTracef calls the default logger's CtxTracef method.
func CtxTracef(ctx context.Context, format string, v ...interface{}) {
	logger.CtxTracef(ctx, format, v...)
}

func CtxForeverf(ctx context.Context, format string, v ...interface{}) {
	logger.CtxInfof(ctx, format, v...)
}

func CtxTraceOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxTraceOption(ctx, v, key, value)
}

func CtxTraceOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxTraceOptions(ctx, v, fields...)
}

func CtxForeverOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxForeverOption(ctx, v, key, value)
}

func CtxForeverOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxForeverOptions(ctx, v, fields...)
}

type defaultLogger struct {
	stdlog *log.Logger
	level  Level
}

func (ll *defaultLogger) CtxForeverOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxInfoOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxForeverOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxInfoOptions(ctx, v, fields...)
}

func (ll *defaultLogger) CtxTraceOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxTraceOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxTraceOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxTraceOptions(ctx, v, fields...)
}

func (ll *defaultLogger) CtxInfoOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxInfoOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxInfoOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxInfoOptions(ctx, v, fields...)
}

func (ll *defaultLogger) CtxNoticeOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxNoticeOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxNoticeOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxNoticeOptions(ctx, v, fields...)
}

func (ll *defaultLogger) CtxWarnOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxWarnOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxWarnOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxWarnOptions(ctx, v, fields...)
}

func (ll *defaultLogger) CtxErrorOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxErrorOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxErrorOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxErrorOptions(ctx, v, fields...)
}

func (ll *defaultLogger) CtxFatalOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxFatalOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxFatalOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxFatalOptions(ctx, v, fields...)
}

func (ll *defaultLogger) SetOutput(w io.Writer) {
	ll.stdlog.SetOutput(w)
}

func (ll *defaultLogger) SetLevel(lv Level) {
	ll.level = lv
}

func (ll *defaultLogger) logf(lv Level, format *string, v ...interface{}) {
	if ll.level > lv {
		return
	}
	msg := toString(lv)
	if format != nil {
		msg += fmt.Sprintf(*format, v...)
	} else {
		msg += fmt.Sprint(v...)
	}
	ll.stdlog.Output(4, msg)
	if lv == LevelFatal {
		os.Exit(1)
	}
}

func (ll *defaultLogger) logf1(lv Level, options interface{}) {
	if ll.level > lv {
		return
	}
	msg := toString(lv)

	fieldsJson, err := json.Marshal(options)
	if err != nil {
		ll.stdlog.Output(4, fmt.Sprintf("Failed to marshal fields: %v", err))
		return
	}

	// 将 fields 添加到消息中
	msg += fmt.Sprintf(", Fields: %s", string(fieldsJson))

	// Add fields to message
	ll.stdlog.Output(4, msg)
	if lv == LevelFatal {
		os.Exit(1)
	}
}

func (ll *defaultLogger) Fatal(v ...interface{}) {
	ll.logf(LevelFatal, nil, v...)
}

func (ll *defaultLogger) Error(v ...interface{}) {
	ll.logf(LevelError, nil, v...)
}

func (ll *defaultLogger) Warn(v ...interface{}) {
	ll.logf(LevelWarn, nil, v...)
}

func (ll *defaultLogger) Notice(v ...interface{}) {
	ll.logf(LevelNotice, nil, v...)
}

func (ll *defaultLogger) Info(v ...interface{}) {
	ll.logf(LevelInfo, nil, v...)
}

func (ll *defaultLogger) Forever(v ...interface{}) {
	ll.logf(LevelForever, nil, v...)
}

func (ll *defaultLogger) Debug(v ...interface{}) {
	ll.logf(LevelDebug, nil, v...)
}

func (ll *defaultLogger) Trace(v ...interface{}) {
	ll.logf(LevelTrace, nil, v...)
}

func (ll *defaultLogger) Fatalf(format string, v ...interface{}) {
	ll.logf(LevelFatal, &format, v...)
}

func (ll *defaultLogger) Errorf(format string, v ...interface{}) {
	ll.logf(LevelError, &format, v...)
}

func (ll *defaultLogger) Warnf(format string, v ...interface{}) {
	ll.logf(LevelWarn, &format, v...)
}

func (ll *defaultLogger) Noticef(format string, v ...interface{}) {
	ll.logf(LevelNotice, &format, v...)
}

func (ll *defaultLogger) Infof(format string, v ...interface{}) {
	ll.logf(LevelInfo, &format, v...)
}

func (ll *defaultLogger) Debugf(format string, v ...interface{}) {
	ll.logf(LevelDebug, &format, v...)
}

func (ll *defaultLogger) Tracef(format string, v ...interface{}) {
	ll.logf(LevelTrace, &format, v...)
}

func (ll *defaultLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelFatal, &format, v...)
}

func (ll *defaultLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelError, &format, v...)
}

func (ll *defaultLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelWarn, &format, v...)
}

func (ll *defaultLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelNotice, &format, v...)
}

func (ll *defaultLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelInfo, &format, v...)
}

func (ll *defaultLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelDebug, &format, v...)
}

func (ll *defaultLogger) CtxDebugOption(ctx context.Context, v interface{}, key string, value interface{}) {
	logger.CtxDebugOption(ctx, v, key, value)
}

func (ll *defaultLogger) CtxDebugOptions(ctx context.Context, v interface{}, fields ...logrus.Fields) {
	logger.CtxDebugOptions(ctx, v, fields...)
}

func (ll *defaultLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelTrace, &format, v...)
}
