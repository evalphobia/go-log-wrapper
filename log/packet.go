package log

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Packet is struct for log data
type Packet struct {
	Title   string
	Data    interface{}
	Err     error
	Request *http.Request
	SQL     string
	Engine  string
	Trace   int // stacktrace depth
	NoTrace bool

	data []interface{}
}

// Fatal logs fatal error
func (p Packet) Fatal() {
	logrus.WithFields(p.createField()).Fatal(p.Title)
}

// Panic logs fatal error
func (p Packet) Panic() {
	logrus.WithFields(p.createField()).Panic(p.Title)
}

// Error logs serious error
func (p Packet) Error() {
	logrus.WithFields(p.createField()).Error(p.Title)
}

// Warn logs warning
func (p Packet) Warn() {
	logrus.WithFields(p.createField()).Warn(p.Title)
}

// Info logs information
func (p Packet) Info() {
	logrus.WithFields(p.createField()).Info(p.Title)
}

// Debug logs development information
func (p Packet) Debug() {
	logrus.WithFields(p.createField()).Debug(p.Title)
}

// AddData adds data for logging
func (p *Packet) AddData(d ...interface{}) {
	p.data = append(p.data, d...)
}

func (p Packet) createField() logrus.Fields {
	f := logrus.Fields{}

	f["message"] = p.Title
	f["value"] = createLogValue(p.Data, p.data)

	if p.Request != nil {
		f["http_request"] = p.Request
	}
	if p.SQL != "" {
		f["query"] = p.SQL
		if p.Engine != "" {
			f["engine"] = p.Engine
		}
	}
	if p.Err != nil {
		f["error"] = p.Err
	}
	if !p.NoTrace {
		traces := getTrace(p.Trace, 4)
		if len(traces) != 0 {
			f["trace"] = traces
		}
	}
	return f
}

func createLogValue(v interface{}, extra []interface{}) interface{} {
	switch {
	case len(extra) == 0:
		return v
	case v == nil:
		return extra
	}
	var list []interface{}
	list = append(list, v)
	return append(list, extra...)
}
