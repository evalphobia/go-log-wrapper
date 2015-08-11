package fluent

import (
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_fluent"
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

func Set(host string, port int) {
	hook := logrus_fluent.NewHook(host, port)
	hook.SetLevels(hookLevel)
	logrus.AddHook(hook)
}
