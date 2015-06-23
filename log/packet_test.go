package log

import (
	"bytes"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPacketError(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Packet{
		Title: "the title",
		Data:  999,
	}.Error()
	output := buf.String()
	assert.Contains(output, `level=error`)
	assert.Contains(output, `message="the title"`)
	assert.Contains(output, `value=999`)
	trace := getTrace(0, 2)
	assert.Contains(output, trace[0].Function)
}

func TestPacketInfo(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Packet{
		Title: "the title",
		Data:  999,
	}.Info()
	output := buf.String()
	assert.Contains(output, `level=info`)
	assert.Contains(output, `message="the title"`)
	assert.Contains(output, `value=999`)
	trace := getTrace(0, 2)
	assert.Contains(output, trace[0].Function)
}

func TestPacketDebug(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Packet{
		Title: "the title",
		Data:  999,
	}.Debug()
	output := buf.String()
	assert.Contains(output, `level=debug`)
	assert.Contains(output, `message="the title"`)
	assert.Contains(output, `value=999`)
	trace := getTrace(0, 2)
	assert.Contains(output, trace[0].Function)
}

func TestPacketCreateField(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	p := Packet{
		Title: "the title",
		Data:  999,
	}
	f := p.createField()

	assert.Equal(p.Title, f["message"])
	assert.Equal(p.Data, f["value"])
}
