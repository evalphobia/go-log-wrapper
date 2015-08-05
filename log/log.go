package log

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
)

// Nothing is dummy variable for import error
var Nothing int

func newLogField(v []interface{}) logrus.Fields {
	f := logrus.Fields{}
	f["trace"] = getTrace(0, 3)

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
func Dump(v interface{}) {
	spew.Dump(v)
}

// Print prints variable information in console
func Print(v interface{}) {
	fmt.Printf("%#v\n", v)
}

// Header prints separator in console
func Header(v ...interface{}) {
	if len(v) < 1 {
		fmt.Printf("=============================================\n")
		return
	}
	fmt.Printf("===================== %s =====================\n", v[0])
	return
}
