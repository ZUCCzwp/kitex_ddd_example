package onelogrus

import (
	"github.com/sirupsen/logrus"
	"time"
)

var _ logrus.Hook = (*TraceHook)(nil)

type TimezoneHook struct {
	Location *time.Location
}

// Levels 返回这个 hook 适用的日志级别
func (hook *TimezoneHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 调整日志条目的时间
func (hook *TimezoneHook) Fire(entry *logrus.Entry) error {
	entry.Time = entry.Time.In(hook.Location)
	return nil
}
