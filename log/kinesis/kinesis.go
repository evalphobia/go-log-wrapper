package kinesis

import (
	"net/http"

	"github.com/evalphobia/logrus_kinesis"
	"github.com/sirupsen/logrus"

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

func Set(name string, conf Config) error {
	hook, err := logrus_kinesis.New(name, logrus_kinesis.Config{
		AccessKey: conf.AccessKey,
		SecretKey: conf.SecretKey,
		Region:    conf.Region,
	})
	if err != nil {
		return err
	}

	hook.SetLevels(hookLevel)
	hook.AddIgnore("context")
	hook.AddFilter("trace", filterTrace)
	hook.AddFilter("http_request", filterRequest)
	hook.Async()
	logrus.AddHook(hook)
	return nil
}

type Config struct {
	AccessKey string
	SecretKey string
	Region    string
}

func filterTrace(v interface{}) interface{} {
	trace, ok := v.([]log.StackTrace)
	if !ok || len(trace) == 0 {
		return v
	}

	return trace[0]
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
