package log

import "github.com/sirupsen/logrus"

// log levels
var (
	LevelsPanic = []logrus.Level{
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
	LevelsError = append(LevelsPanic, logrus.ErrorLevel)
	LevelsWarn  = append(LevelsError, logrus.WarnLevel)
	LevelsInfo  = append(LevelsWarn, logrus.InfoLevel)
	LevelsDebug = append(LevelsInfo, logrus.DebugLevel)
)
