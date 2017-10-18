package fluent

import (
	"net/http"

	"github.com/evalphobia/logrus_fluent"
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

func Set(host string, port int) error {
	hook, err := logrus_fluent.New(host, port)
	if err != nil {
		return err
	}

	hook.SetLevels(hookLevel)
	hook.AddIgnore("context")
	hook.AddFilter("http_request", filterRequest)
	hook.AddFilter("error", logrus_fluent.FilterError)

	logrus.AddHook(hook)
	return nil
}

func filterRequest(v interface{}) interface{} {
	req, ok := v.(*http.Request)
	if !ok {
		return v
	}

	return request{
		Method:     req.Method,
		Host:       req.Host,
		RequestURI: req.RequestURI,
		RemoteAddr: req.RemoteAddr,
	}
}

type request struct {
	Method     string
	Host       string
	RequestURI string
	RemoteAddr string
}
