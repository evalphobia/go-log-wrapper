package log

import (
	"runtime"
)

type StackTrace struct {
	File     string
	Function string
	Line     int
	Path     string
}

func getTrace(depth, skip int) []StackTrace {
	depth++
	var traces []StackTrace
	for i := 0; i < depth; i++ {
		st, ok := trace(skip + i)
		if !ok {
			continue
		}
		traces = append(traces, st)
	}
	return traces
}

func trace(depth int) (StackTrace, bool) {
	pt, file, line, ok := runtime.Caller(depth)
	if !ok {
		return StackTrace{}, false
	}
	trace := StackTrace{
		File:     trimPath(file),
		Line:     line,
		Path:     file,
		Function: getFunctionName(pt),
	}
	return trace, true
}

func trimPath(path string) string {
	trimed := path
	for i := len(path) - 1; i > 0; i-- {
		if path[i] == '/' {
			trimed = path[i+1:]
			break
		}
	}
	return trimed
}

func getFunctionName(pt uintptr) string {
	fn := runtime.FuncForPC(pt)
	if fn == nil {
		return ""
	}
	return fn.Name()
}
