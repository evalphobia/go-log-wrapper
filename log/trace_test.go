package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTrace(t *testing.T) {
	assert := assert.New(t)

	trace := getTrace(0, 0)
	assert.Equal(trace[0].File, "trace.go")
	assert.Contains(trace[0].Function, "log-wrapper/log.trace")

	trace = getTrace(0, 1)
	assert.Equal(trace[0].File, "trace.go")
	assert.Contains(trace[0].Function, "log-wrapper/log.getTrace")

	trace = getTrace(0, 2)
	assert.Equal(trace[0].File, "trace_test.go")
	assert.Contains(trace[0].Function, "log-wrapper/log.TestGetTrace")
}

func TestTrace(t *testing.T) {
	assert := assert.New(t)

	tc, _ := trace(0)
	assert.Equal(tc.File, "trace.go")
	assert.Contains(tc.Function, "log-wrapper/log.trace")

	tc, _ = trace(1)
	assert.Equal(tc.File, "trace_test.go")
	assert.Contains(tc.Function, "log-wrapper/log.TestTrace")
}
