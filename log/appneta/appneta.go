package appneta

import (
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_appneta"

	"github.com/evalphobia/go-log-wrapper/log"
)

// default logging hook level
var hookLevel = log.LevelsError

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

func Set() {
	hook := logrus_appneta.NewHook()
	logrus.AddHook(hook)
}
