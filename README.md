go-log-wrapper
====

[![Build Status](https://travis-ci.org/evalphobia/go-log-wrapper.svg?branch=master)](https://travis-ci.org/evalphobia/go-log-wrapper) [![codecov](https://codecov.io/gh/evalphobia/go-log-wrapper/branch/master/graph/badge.svg)](https://codecov.io/gh/evalphobia/go-log-wrapper)
 [![GoDoc](https://godoc.org/github.com/evalphobia/go-log-wrapper?status.svg)](https://godoc.org/github.com/evalphobia/go-log-wrapper)

simple logrus wrapper to use easily

# Supported Hooks

- Appneta TraceView
- Fluent
- Kinesis
- Sentry
- Stackdriver Logging

# Quick Usage

## Print

```go
import (
	"github.com/evalphobia/go-log-wrapper/log"
)

func main() {
	// Print
	log.Print("print me")
	// "print me"
	log.Dump("dump me")
	// "dump me"

	// Header
	log.Header()
	// "============================================="
	log.Header("header")
	// "===================== header ====================="

	// Mark
	log.Mark()
	// "main"
}
```


## Logrus wrapper

```go
import (
	"github.com/evalphobia/go-log-wrapper/log"
)

func main(){
	user := getUser()

	err := someFunc(user)
	if err != nil {
		log.Packet{
			Title: "error on someFunc",
			Data: user,
			Err: err,
			UserID: user.ID, // used for logrus_sentry
			Tag: "error", // used for logrus_fluent
		}.Error()

		//     Send as below:
		// logrus.WithFields(logrus.Fields{
		// 	"value":   user,
		// 	"error":   err,
		// 	"user_id": user.ID,
		// 	"tag":     "error",
		// }).Error("error on someFunc")
	}

}
```

## Logrus Hook

```go
import(
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/go-log-wrapper/log/fluent"
	"github.com/evalphobia/go-log-wrapper/log/sentry"
)

func main(){
	// fluentd
	fluent.AddLevel(logrus.DebugLevel)
	fluent.Set(host, port)

	// Sentry
	sentry.AddLevel(logrus.WarnLevel)
	sentry.Set("https://....:....@app.getsentry.com/99999")

	// AppNeta TraceView
	appneta.AddLevel(logrus.WarnLevel)
	appneta.Set()
}
```

## Logger Instance

```go
package main

import (
	"github.com/evalphobia/go-log-wrapper/log"
)

var logger1 = log.NewLogger()
var logger2 = log.NewLogger()


func main() {
	// global log level sets Error-level
	log.SetGlobalLogLevel(log.Error)

	// logger1 write log to file
	file, _ := os.Create(`/path/to/file`)
	logger1.SetOutput(file)
	logger1.SetLogLevel(log.Debug)

	// logger2 does not write to file, use Sentry hook
	logger2.DisableOutput()
	sentry.SetLogger(logger2, dsn)

	// global logging
	log.Packet{
		Title: "Error by global level",
	}.Error()

	// use logger1
	log.Packet{
		Logger: logger1,
		Title:  "Info by logger1",
	}.Info()

	// use logger2
	log.Packet{
		Logger: logger2,
		Title:  "logger2 does not print, only hooks",
	}.Error()
}
```
