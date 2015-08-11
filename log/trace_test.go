package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTrace(t *testing.T) {
	assert := assert.New(t)

	trace := getTrace(0, 0)
	assert.Equal(trace[0].File, "trace.go")
	assert.Contains(trace[0].Module, "go-log-wrapper/log")
	assert.Contains(trace[0].Function, "trace")

	trace = getTrace(0, 1)
	assert.Equal(trace[0].File, "trace.go")
	assert.Contains(trace[0].Module, "go-log-wrapper/log")
	assert.Contains(trace[0].Function, "getTrace")

	trace = getTrace(0, 2)
	assert.Equal(trace[0].File, "trace_test.go")
	assert.Contains(trace[0].Module, "go-log-wrapper/log")
	assert.Contains(trace[0].Function, "TestGetTrace")

}

func TestTrace(t *testing.T) {
	assert := assert.New(t)

	tc, _ := trace(0)
	assert.Equal(tc.File, "trace.go")
	assert.Contains(tc.Module, "go-log-wrapper/log")
	assert.Contains(tc.Function, "trace")

	tc, _ = trace(1)
	assert.Equal(tc.File, "trace_test.go")
	assert.Contains(tc.Module, "go-log-wrapper/log")
	assert.Contains(tc.Function, "TestTrace")
}
