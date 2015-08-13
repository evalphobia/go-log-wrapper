package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTraces(t *testing.T) {
	assert := assert.New(t)

	trace := GetTraces(0, 0)
	assert.Equal(trace[0].File, "trace.go")
	assert.Contains(trace[0].Module, "go-log-wrapper/log")
	assert.Contains(trace[0].Function, "Trace")

	trace = GetTraces(0, 1)
	assert.Equal(trace[0].File, "trace.go")
	assert.Contains(trace[0].Module, "go-log-wrapper/log")
	assert.Contains(trace[0].Function, "GetTraces")

	trace = GetTraces(0, 2)
	assert.Equal(trace[0].File, "trace_test.go")
	assert.Contains(trace[0].Module, "go-log-wrapper/log")
	assert.Contains(trace[0].Function, "TestGetTrace")

}

func TestTrace(t *testing.T) {
	assert := assert.New(t)

	tc, _ := Trace(0)
	assert.Equal(tc.File, "trace.go")
	assert.Contains(tc.Module, "go-log-wrapper/log")
	assert.Contains(tc.Function, "Trace")

	tc, _ = Trace(1)
	assert.Equal(tc.File, "trace_test.go")
	assert.Contains(tc.Module, "go-log-wrapper/log")
	assert.Contains(tc.Function, "TestTrace")
}
