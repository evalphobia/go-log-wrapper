package fluent

import (
	"net/http"

	"github.com/evalphobia/logrus_fluent"
	"github.com/sirupsen/logrus"

	"github.com/evalphobia/go-log-wrapper/log"
)

// default logging hook level
var hookLevel = log.LevelsInfo

// SetLevels sets log levels.
func SetLevels(levels []logrus.Level) {
	hookLevel = levels
}

// SetLevelsAsDebug sets log level as Debug.
func SetLevelsAsDebug() {
	hookLevel = log.LevelsDebug
}

// SetLevelsAsInfo sets log level as Info.
func SetLevelsAsInfo() {
	hookLevel = log.LevelsInfo
}

// SetLevelsAsWarn sets log level as Warn.
func SetLevelsAsWarn() {
	hookLevel = log.LevelsWarn
}

// SetLevelsAsError sets log level as Error.
func SetLevelsAsError() {
	hookLevel = log.LevelsError
}

// SetLevelsAsPanic sets log level as Panic.
func SetLevelsAsPanic() {
	hookLevel = log.LevelsPanic
}

// AddLevel adds log level.
func AddLevel(level logrus.Level) {
	hookLevel = append(hookLevel, level)
}

// New creates logging hook for Stackdriver logging.
func New(host string, port int) (logrus.Hook, error) {
	hook, err := logrus_fluent.New(host, port)
	if err != nil {
		return nil, err
	}

	hook.SetLevels(hookLevel)
	hook.AddIgnore("context")
	hook.AddFilter("http_request", filterRequest)
	hook.AddFilter("error", logrus_fluent.FilterError)
	return hook, nil
}

// Set sets fluent hook to global level logger.
func Set(host string, port int) error {
	hook, err := New(host, port)
	if err != nil {
		return err
	}

	logrus.AddHook(hook)
	return nil
}

// SetLogger sets fluent hook to given logger.
func SetLogger(logger log.Logger, host string, port int) error {
	hook, err := New(host, port)
	if err != nil {
		return err
	}
	logger.AddHook(hook)
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
