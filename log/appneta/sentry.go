package appneta

import (
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_appneta"
)

// default logging hook level
var hookLevel = []logrus.Level{
	logrus.PanicLevel,
	logrus.ErrorLevel,
}

func SetLevels(levels []logrus.Level) {
	hookLevel = levels
}

func AddLevel(level logrus.Level) {
	hookLevel = append(hookLevel, level)
}

func Set() {
	hook := logrus_appneta.NewHook()
	logrus.AddHook(hook)
}
