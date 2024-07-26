package onelogrus

import (
	"errors"
	"github.com/sirupsen/logrus"
)

type ForeverLevel int

const (
	LevelHigh   ForeverLevel = 10
	LevelMedium ForeverLevel = 5
	LevelLow    ForeverLevel = 1
)

type ForeverLevelConfig struct {
	MinLevel ForeverLevel
}

type ForeverLevelHook struct {
	cfg *ForeverLevelConfig
}

func NewForeverLevelHook(cfg *ForeverLevelConfig) *ForeverLevelHook {
	return &ForeverLevelHook{cfg: cfg}
}

func (h *ForeverLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *ForeverLevelHook) Fire(entry *logrus.Entry) error {
	if levelVal, ok := entry.Data["foreverLevel"]; ok {
		if level, ok := levelVal.(ForeverLevel); ok && level < h.cfg.MinLevel {
			return errors.New("log entry below minimum forever level")
		}
	}
	return nil
}
