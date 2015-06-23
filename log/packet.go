package log

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Packet is struct for log data
type Packet struct {
	Title   string
	Data    interface{}
	Request *http.Request
	SQL     string
	Engine  string
	Trace   int // stacktrace depth
}

// Fatal logs fatal error
func (p Packet) Fatal() {
	logrus.WithFields(p.createField()).Fatal(p.Title)
}

// Error logs serious error
func (p Packet) Error() {
	logrus.WithFields(p.createField()).Error(p.Title)
}

// Info logs information
func (p Packet) Info() {
	logrus.WithFields(p.createField()).Info(p.Title)
}

// Debug logs development information
func (p Packet) Debug() {
	logrus.WithFields(p.createField()).Debug(p.Title)
}

func (p Packet) createField() logrus.Fields {
	f := logrus.Fields{}

	f["message"] = p.Title
	f["value"] = p.Data
	if p.Request != nil {
		f["http_request"] = p.Request
	}
	if p.SQL != "" {
		f["query"] = p.SQL
		if p.Engine != "" {
			f["engine"] = p.Engine
		}
	}
	traces := getTrace(p.Trace, 4)
	if len(traces) != 0 {
		f["trace"] = traces
	}
	return f
}
