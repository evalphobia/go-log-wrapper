package sentry

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
)

// default timeout
var timeout = 200 * time.Millisecond

// default logging hook level
var hookLevel = []logrus.Level{
	logrus.PanicLevel,
	logrus.ErrorLevel,
}

func SetTimeout(i time.Duration) {
	timeout = i * time.Millisecond
}

func SetLevels(levels []logrus.Level) {
	hookLevel = levels
}

func AddLevel(level logrus.Level) {
	hookLevel = append(hookLevel, level)
}

func Set(dsn string) {
	hook, err := logrus_sentry.NewSentryHook(dsn, hookLevel)
	if err != nil {
		return
	}
	hook.Timeout = timeout
	logrus.AddHook(hook)
}
