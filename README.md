go-log-wrapper
----

[![GoDoc][1]][2] [![License: Apache 2.0][23]][24] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22]

[1]: https://godoc.org/github.com/evalphobia/go-log-wrapper?status.svg
[2]: https://godoc.org/github.com/evalphobia/go-log-wrapper
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/go-log-wrapper.svg
[6]: https://github.com/evalphobia/go-log-wrapper/releases/latest
[7]: https://travis-ci.org/evalphobia/go-log-wrapper.svg?branch=master
[8]: https://travis-ci.org/evalphobia/go-log-wrapper
[9]: https://coveralls.io/repos/evalphobia/go-log-wrapper/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/go-log-wrapper?branch=master
[11]: https://codecov.io/github/evalphobia/go-log-wrapper/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/go-log-wrapper?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/go-log-wrapper
[14]: https://goreportcard.com/report/github.com/evalphobia/go-log-wrapper
[15]: https://img.shields.io/github/downloads/evalphobia/go-log-wrapper/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/go-log-wrapper/releases
[17]: https://img.shields.io/github/stars/evalphobia/go-log-wrapper.svg
[18]: https://github.com/evalphobia/go-log-wrapper/stargazers
[19]: https://codeclimate.com/github/evalphobia/go-log-wrapper/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/go-log-wrapper
[21]: https://bettercodehub.com/edge/badge/evalphobia/go-log-wrapper?branch=master
[22]: https://bettercodehub.com/
[23]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[24]: LICENSE.md


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
