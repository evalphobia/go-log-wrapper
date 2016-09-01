package fluent

import (
	"net/http"

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
