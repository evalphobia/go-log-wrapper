package kinesis

import (
	"net/http"

	"github.com/evalphobia/logrus_kinesis"
	"github.com/sirupsen/logrus"

	"github.com/evalphobia/go-log-wrapper/log"
)

// default logging hook level
var hookLevel = log.LevelsError

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

// New creates logging hook for Kinesis.
func New(name string, conf Config) (logrus.Hook, error) {
	hook, err := logrus_kinesis.New(name, logrus_kinesis.Config{
		AccessKey: conf.AccessKey,
		SecretKey: conf.SecretKey,
		Region:    conf.Region,
	})
	if err != nil {
		return nil, err
	}

	hook.SetLevels(hookLevel)
	hook.AddIgnore("context")
	hook.AddFilter("trace", filterTrace)
	hook.AddFilter("http_request", filterRequest)
	hook.Async()
	return hook, nil
}

// Set sets Kinesis hook to global level logger.
func Set(name string, conf Config) error {
	hook, err := New(name, conf)
	if err != nil {
		return err
	}
	logrus.AddHook(hook)
	return nil
}

// SetLogger sets Kinesis hook to given logger.
func SetLogger(logger log.Logger, name string, conf Config) {
	hook, err := New(name, conf)
	if err != nil {
		return
	}
	logger.AddHook(hook)
}

// Config is struct for Kinesis's configuration.
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
