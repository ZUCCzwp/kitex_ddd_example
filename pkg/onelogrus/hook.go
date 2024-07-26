package onelogrus

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Ref to https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/README.md#json-formats
const (
	traceIDKey    = "trace_id"
	spanIDKey     = "span_id"
	traceFlagsKey = "trace_flags"
)

var _ logrus.Hook = (*TraceHook)(nil)

type TraceHookConfig struct {
	recordStackTraceInSpan bool
	enableLevels           []logrus.Level
	errorSpanLevel         logrus.Level
	timezone               *time.Location
}

type TraceHook struct {
	cfg *TraceHookConfig
}

func NewTraceHook(cfg *TraceHookConfig) *TraceHook {
	return &TraceHook{cfg: cfg}
}

func (h *TraceHook) Levels() []logrus.Level {
	return h.cfg.enableLevels
}

func (h *TraceHook) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}

	span := trace.SpanFromContext(entry.Context)

	// check span context
	spanContext := span.SpanContext()
	if !spanContext.IsValid() {
		return nil
	}

	// attach span context to log entry data fields
	entry.Data[traceIDKey] = spanContext.TraceID()
	entry.Data[spanIDKey] = spanContext.SpanID()
	entry.Data[traceFlagsKey] = spanContext.TraceFlags()

	// non recording spans do not support modifying
	if !span.IsRecording() {
		return nil
	}

	if entry.Level <= h.cfg.errorSpanLevel {
		// set span status
		span.SetStatus(codes.Error, "")

		// record error with stack trace
		span.RecordError(errors.New(entry.Message), trace.WithStackTrace(h.cfg.recordStackTraceInSpan))
	}

	return nil
}
