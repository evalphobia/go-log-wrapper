package log

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

// Nothing is dummy variable for import error
var Nothing int

// SetGlobalLogLevel sets log level.
func SetGlobalLogLevel(l logrus.Level) {
	logrus.SetLevel(l)
}

// SetGlobalFormatter sets Fomatter.
func SetGlobalFormatter(f logrus.Formatter) {
	logrus.SetFormatter(f)
}

func newLogField(v []interface{}) logrus.Fields {
	f := logrus.Fields{}
	f["trace"] = GetTraces(0, 3)

	if len(v) > 1 {
		f["message"] = v[0].(string)
		f["value"] = v[1]
	} else {
		f["value"] = v[0]
	}
	return f
}

// Error logs serious error
func Error(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Error("error")
}

// Error logs warning
func Warn(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Warn("warn")
}

// Info logs information
func Info(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Info("info")
}

// Debug logs development information
func Debug(v ...interface{}) {
	f := newLogField(v)
	logrus.WithFields(f).Debug("debug")
}

// Dump prints dump variable in console
func Dump(v ...interface{}) {
	spew.Dump(v)
}

// Json prints data on json format in console
func Json(v ...interface{}) {
	for _, vv := range v {
		byt, err := json.Marshal(vv)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(byt))
	}
}

// Print prints variable information in console
func Print(v interface{}) {
	fmt.Printf("%+v\n", v)
}

// Header prints separator in console
func Header(v ...interface{}) {
	if len(v) < 1 {
		fmt.Printf("=============================================\n")
		return
	}
	fmt.Printf("===================== %+v =====================\n", v[0])
	return
}

// Mark prints trace info in console
func Mark(i ...int) {
	depth := 2
	if len(i) > 0 {
		depth = i[0]
	}
	v, _ := Trace(depth)
	Header(v)
}
