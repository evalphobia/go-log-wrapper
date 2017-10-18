package sentry

import (
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"

	"github.com/evalphobia/go-log-wrapper/log"
)

// default timeout
var timeout = 200 * time.Millisecond

// default logging hook level
var hookLevel = log.LevelsError

func SetTimeout(i time.Duration) {
	timeout = i * time.Millisecond
}

func SetLevels(levels []logrus.Level) {
	hookLevel = levels
}

func SetLevelsAsDebug() {
	hookLevel = log.LevelsDebug
}

func SetLevelsAsInfo() {
	hookLevel = log.LevelsInfo
}

func SetLevelsAsWarn() {
	hookLevel = log.LevelsWarn
}

func SetLevelsAsError() {
	hookLevel = log.LevelsError
}

func SetLevelsAsPanic() {
	hookLevel = log.LevelsPanic
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
	hook.AddIgnore("context")
	logrus.AddHook(hook)
}
