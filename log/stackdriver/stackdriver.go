package stackdriver

import (
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/google-api-go-wrapper/config"
	"github.com/evalphobia/logrus_stackdriver"
)

// default logging hook level
var hookLevel = []logrus.Level{
	logrus.PanicLevel,
	logrus.FatalLevel,
	logrus.ErrorLevel,
	logrus.WarnLevel,
	logrus.InfoLevel,
}

func SetLevels(levels []logrus.Level) {
	hookLevel = levels
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
