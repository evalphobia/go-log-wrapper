package stackdriver

import (
	"github.com/evalphobia/google-api-go-wrapper/config"
	"github.com/evalphobia/logrus_stackdriver"
	"github.com/sirupsen/logrus"

	"github.com/evalphobia/go-log-wrapper/log"
)

// default logging hook level
var hookLevel = log.LevelsInfo

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

// New creates logging hook for Stackdriver logging.
func New(conf Config) (logrus.Hook, error) {
	hook, err := logrus_stackdriver.NewWithConfig(conf.ProjectID, conf.LogName, config.Config{
		Email:      conf.Email,
		PrivateKey: conf.PrivateKey,
	})
	if err != nil {
		return nil, err
	}

	hook.SetLabels(conf.Labels)
	hook.SetLevels(hookLevel)
	hook.AddIgnore("context")
	hook.Async()
	return hook, nil
}

// Set sets Stackdriver hook to global level logger.
func Set(conf Config) error {
	hook, err := New(conf)
	if err != nil {
		return err
	}
	logrus.AddHook(hook)
	return nil
}

// SetLogger sets Stackdriver hook to given logger.
func SetLogger(logger log.Logger, conf Config) error {
	hook, err := New(conf)
	if err != nil {
		return err
	}
	logger.AddHook(hook)
	return nil
}

// Config is struct for Stackdriver's configuration.
type Config struct {
	ProjectID  string
	LogName    string
	Email      string
	PrivateKey string
	Labels     map[string]string
}
