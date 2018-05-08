package appneta

import (
	"github.com/evalphobia/logrus_appneta"
	"github.com/sirupsen/logrus"

	"github.com/evalphobia/go-log-wrapper/log"
)

// default logging hook level
var hookLevel = log.LevelsError

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

// New creates logging hook for TraceView logging.
func New() (logrus.Hook, error) {
	hook := logrus_appneta.NewHook()
	return hook, nil
}

// Set sets TraceView hook to global level logger.
func Set() {
	hook, err := New()
	if err != nil {
		return
	}
	logrus.AddHook(hook)
}

// SetLogger sets TraceView hook to given logger.
func SetLogger(logger log.Logger) error {
	hook, err := New()
	if err != nil {
		return err
	}
	logger.AddHook(hook)
	return nil
}
