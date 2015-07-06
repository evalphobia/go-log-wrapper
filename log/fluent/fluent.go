package fluent

import (
	"github.com/Sirupsen/logrus"
	// logrus_fluent "github.com/Sirupsen/logrus/hooks/fluent"
	logrus_fluent "github.com/evalphobia/logrus/hooks/fluent"
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
	logrus.AddHook(hook)
}
