package log

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	defaultTraceDepth = 10
	defaultSkipDepth  = 4
)

// Packet is struct for log data
type Packet struct {
	*Logger

	Severity string
	Title    string
	Data     interface{}
	Err      error
	Request  *http.Request
	Context  context.Context
	SQL      string
	Engine   string
	UserID   string
	Tag      string

	Trace     int // stacktrace depth
	TraceSkip int
	TraceData interface{}
	NoTrace   bool

	DataList []interface{}
}

// SetLogger sets Logger.
func (p *Packet) SetLogger(l *Logger) *Packet {
	p.Logger = l
	return p
}

// Fatal logs fatal-level log.
func (p Packet) Fatal() {
	if p.Logger == nil {
		logrus.WithFields(p.createField("fatal")).Fatal(p.Title)
		return
	}
	p.Logger.WithFields(p.createField("fatal")).Fatal(p.Title)
}

// Panic logs panic-level log.
func (p Packet) Panic() {
	if p.Logger == nil {
		logrus.WithFields(p.createField("panic")).Panic(p.Title)
		return
	}
	p.Logger.WithFields(p.createField("panic")).Panic(p.Title)
}

// Error logs error-level log.
func (p Packet) Error() {
	if p.Logger == nil {
		logrus.WithFields(p.createField("error")).Error(p.Title)
		return
	}
	p.Logger.WithFields(p.createField("error")).Error(p.Title)
}

// Warn logs warn-level log.
func (p Packet) Warn() {
	if p.Logger == nil {
		logrus.WithFields(p.createField("warn")).Warn(p.Title)
		return
	}
	p.Logger.WithFields(p.createField("warn")).Warn(p.Title)
}

// Info logs info-level log.
func (p Packet) Info() {
	if p.Logger == nil {
		logrus.WithFields(p.createField("info")).Info(p.Title)
		return
	}
	p.Logger.WithFields(p.createField("info")).Info(p.Title)
}

// Debug logs debug-level log.
func (p Packet) Debug() {
	if p.Logger == nil {
		logrus.WithFields(p.createField("debug")).Debug(p.Title)
		return
	}
	p.Logger.WithFields(p.createField("debug")).Debug(p.Title)
}

// AddData adds data for logging
func (p *Packet) AddData(d ...interface{}) *Packet {
	p.DataList = append(p.DataList, d...)
	return p
}

func (p Packet) createField(severity string) logrus.Fields {
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

	f["severity"] = severity
	if p.Severity != "" {
		f["severity"] = p.Severity
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
