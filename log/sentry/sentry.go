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

// SetTimeout sets timeout for Sentry API.
func SetTimeout(i time.Duration) {
	timeout = i * time.Millisecond
}

// SetLevels sets log levels.
func SetLevels(levels []logrus.Level) {
	hookLevel = levels
}

// SetLevelsAsDebug sets log level as Debug.
func SetLevelsAsDebug() {
	hookLevel = log.LevelsDebug
}

// SetLevelsAsInfo sets log level as Info.
func SetLevelsAsInfo() {
	hookLevel = log.LevelsInfo
}

// SetLevelsAsWarn sets log level as Warn.
func SetLevelsAsWarn() {
	hookLevel = log.LevelsWarn
}

// SetLevelsAsError sets log level as Error.
func SetLevelsAsError() {
	hookLevel = log.LevelsError
}

// SetLevelsAsPanic sets log level as Panic.
func SetLevelsAsPanic() {
	hookLevel = log.LevelsPanic
}

// AddLevel adds log level.
func AddLevel(level logrus.Level) {
	hookLevel = append(hookLevel, level)
}

// New creates logging hook for Sentry.
func New(dsn string) (logrus.Hook, error) {
	hook, err := logrus_sentry.NewSentryHook(dsn, hookLevel)
	if err != nil {
		return nil, err
	}
	hook.Timeout = timeout
	hook.AddIgnore("context")
	return hook, nil
}

// Set sets Sentry hook to global level logger.
func Set(dsn string) {
	hook, err := New(dsn)
	if err != nil {
		return
	}
	logrus.AddHook(hook)
}

// SetLogger sets Sentry hook to given logger.
func SetLogger(logger log.Logger, dsn string) error {
	hook, err := New(dsn)
	if err != nil {
		return err
	}
	logger.AddHook(hook)
	return nil
}
