package stackdriver

import (
	"github.com/evalphobia/google-api-go-wrapper/config"
	"github.com/evalphobia/logrus_stackdriver"
	"github.com/sirupsen/logrus"

	"github.com/evalphobia/go-log-wrapper/log"
)

// default logging hook level
var hookLevel = log.LevelsInfo

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

func Set(conf Config) error {
	hook, err := logrus_stackdriver.NewWithConfig(conf.ProjectID, conf.LogName, config.Config{
		Email:      conf.Email,
		PrivateKey: conf.PrivateKey,
	})
	if err != nil {
		return err
	}

	hook.SetLabels(conf.Labels)
	hook.SetLevels(hookLevel)
	hook.AddIgnore("context")
	hook.Async()
	logrus.AddHook(hook)
	return nil
}

type Config struct {
	ProjectID  string
	LogName    string
	Email      string
	PrivateKey string
	Labels     map[string]string
}
