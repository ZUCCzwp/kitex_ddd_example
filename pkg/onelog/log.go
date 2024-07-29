package onelog

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
)

// FormatLogger is a logger interfaces that output logs with a format.
type FormatLogger interface {
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Noticef(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

// Logger is a logger interfaces that provides logging function with levels.
type Logger interface {
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Notice(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Forever(v ...interface{})
}

// CtxLogger is a logger interfaces that accepts a context argument and output
// logs with a format.
type CtxLogger interface {
	CtxTracef(ctx context.Context, format string, v ...interface{})
	CtxTraceOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxTraceOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
	CtxInfof(ctx context.Context, format string, v ...interface{})
	CtxInfoOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxInfoOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
	CtxNoticef(ctx context.Context, format string, v ...interface{})
	CtxNoticeOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxNoticeOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
	CtxWarnf(ctx context.Context, format string, v ...interface{})
	CtxWarnOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxWarnOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
	CtxErrorf(ctx context.Context, format string, v ...interface{})
	CtxErrorOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxErrorOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
	CtxFatalf(ctx context.Context, format string, v ...interface{})
	CtxFatalOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxFatalOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
	CtxDebugf(ctx context.Context, format string, v ...interface{})
	CtxDebugOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxDebugOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
	CtxForeverOption(ctx context.Context, v interface{}, key string, value interface{})
	CtxForeverOptions(ctx context.Context, v interface{}, fields ...logrus.Fields)
}

// Control provides methods to config a logger.
type Control interface {
	SetLevel(Level)
	SetOutput(io.Writer)
}

// FullLogger is the combination of Logger, FormatLogger, CtxLogger and Control.
type FullLogger interface {
	Logger
	FormatLogger
	CtxLogger
	Control
}

// Level defines the priority of a log message.
// When a logger is configured with a level, any log message with a lower
// log level (smaller by integer comparison) will not be output.
type Level int

// The levels of logs.
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelFatal
	LevelForever
)

var strs = []string{
	"[Trace] ",
	"[Debug] ",
	"[Info] ",
	"[Notice] ",
	"[Warn] ",
	"[Error] ",
	"[Fatal] ",
	"[Forever] ",
}

//func toString(lv Level) string {
//	if lv >= LevelTrace && lv <= LevelForever {
//		return strs[lv]
//	}
//	return fmt.Sprintf("[?%d] ", lv)
//}

func toString(lv Level) string {
	switch lv {
	case LevelTrace:
		return "[TRACE] "
	case LevelDebug:
		return "[DEBUG] "
	case LevelInfo:
		return "[INFO] "
	case LevelNotice:
		return "[NOTICE] "
	case LevelWarn:
		return "[WARN] "
	case LevelError:
		return "[ERROR] "
	case LevelFatal:
		return "[FATAL] "
	case LevelForever:
		return "[FOREVER] " // 新添加的日志级别的字符串表示
	default:
		return "[UNKNOWN] "
	}
}
