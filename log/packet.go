package log

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	defaultTraceDepth = 10
	defaultSkipDepth  = 4
)

// Packet is struct for log data
type Packet struct {
	Title   string
	Data    interface{}
	Err     error
	Request *http.Request
	Context context.Context
	SQL     string
	Engine  string
	UserID  string
	Tag     string

	Trace     int // stacktrace depth
	TraceSkip int
	TraceData interface{}
	NoTrace   bool

	DataList []interface{}
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
func (p *Packet) AddData(d ...interface{}) *Packet {
	p.DataList = append(p.DataList, d...)
	return p
}

func (p Packet) createField() logrus.Fields {
	f := logrus.Fields{}

	f["tag"] = p.Tag
	f["value"] = createLogValue(p.Data, p.DataList)

	if p.Request != nil {
		f["http_request"] = p.Request
	}
	if p.Context != nil {
		f["context"] = p.Context
	}
	if p.SQL != "" {
		f["query"] = p.SQL
		if p.Engine != "" {
			f["engine"] = p.Engine
		}
	}
	if p.UserID != "" {
		f["user_id"] = p.UserID
	}

	if p.Err != nil {
		f["error"] = p.Err
	}

	switch {
	case p.NoTrace:
		return f
	case p.TraceData != nil:
		f["trace"] = p.TraceData
		return f
	}

	// assign default value
	if p.Trace == 0 {
		p.Trace = defaultTraceDepth
	}
	if p.TraceSkip == 0 {
		p.TraceSkip = defaultSkipDepth
	}

	traces := GetTraces(p.Trace, p.TraceSkip)
	if len(traces) != 0 {
		f["trace"] = traces
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
