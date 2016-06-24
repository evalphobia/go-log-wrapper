go-log-wrapper
====

[![Build Status](https://travis-ci.org/evalphobia/go-log-wrapper.svg?branch=master)](https://travis-ci.org/evalphobia/go-log-wrapper) [![codecov](https://codecov.io/gh/evalphobia/go-log-wrapper/branch/master/graph/badge.svg)](https://codecov.io/gh/evalphobia/go-log-wrapper)
 [![GoDoc](https://godoc.org/github.com/evalphobia/go-log-wrapper?status.svg)](https://godoc.org/github.com/evalphobia/go-log-wrapper)

simple logrus wrapper to use easily

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
