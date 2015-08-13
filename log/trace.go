package log

import (
	"fmt"
	"runtime"
	"strings"
)

type StackTrace struct {
	File     string
	Module   string
	Function string
	Line     int
	Path     string
}

func (s StackTrace) String() string {
	return fmt.Sprintf("%s:%d %s.%s", s.File, s.Line, s.Module, s.Function)
}

func GetTraces(depth, skip int) []StackTrace {
	depth++
	var traces []StackTrace
	for i := 0; i < depth; i++ {
		st, ok := Trace(skip + i)
		if !ok {
			break
		}
		traces = append(traces, st)
	}
	return traces
}

func Trace(depth int) (StackTrace, bool) {
	pt, file, line, ok := runtime.Caller(depth)
	if !ok {
		return StackTrace{}, false
	}

	trace := StackTrace{
		File: trimPath(file),
		Line: line,
		Path: file,
	}
	trace.Module, trace.Function = getFunctionName(pt)
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

func getFunctionName(pt uintptr) (string, string) {
	fn := runtime.FuncForPC(pt)
	if fn == nil {
		return "", ""
	}
	pack := ""
	name := fn.Name()
	if idx := strings.LastIndex(name, "."); idx != -1 {
		pack = name[:idx]
		name = name[idx+1:]
	}
	name = strings.Replace(name, "·", ".", -1)
	return pack, name
}
