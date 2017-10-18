package log

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPacketError(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Packet{
		Title: "the title",
		Data:  999,
		Tag:   "my_tag",
	}.Error()
	output := buf.String()
	assert.Contains(output, `level=error`)
	assert.Contains(output, `msg="the title"`)
	assert.Contains(output, `tag=my_tag`)
	assert.Contains(output, `value=999`)
	trace := GetTraces(0, 2)
	assert.Contains(output, trace[0].Function)
}

func TestPacketInfo(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Packet{
		Title: "the title",
		Data:  999,
		Tag:   "my_tag",
	}.Info()
	output := buf.String()
	assert.Contains(output, `level=info`)
	assert.Contains(output, `msg="the title"`)
	assert.Contains(output, `tag=my_tag`)
	assert.Contains(output, `value=999`)
	trace := GetTraces(0, 2)
	assert.Contains(output, trace[0].Function)
}

func TestPacketDebug(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Packet{
		Title: "the title",
		Data:  999,
		Tag:   "my_tag",
	}.Debug()
	output := buf.String()
	assert.Contains(output, `level=debug`)
	assert.Contains(output, `msg="the title"`)
	assert.Contains(output, `tag=my_tag`)
	assert.Contains(output, `value=999`)
	trace := GetTraces(0, 2)
	assert.Contains(output, trace[0].Function)
}

func TestPacketCreateField(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	p := Packet{
		Title: "the title",
		Data:  999,
		Tag:   "my_tag",
	}
	f := p.createField()

	assert.Equal(p.Data, f["value"])
	assert.Equal(p.Tag, f["tag"])

	p.AddData("foo", "bar", 111)
	f = p.createField()
	assert.Equal([]interface{}{999, "foo", "bar", 111}, f["value"])
}

func TestPacketNoTrace(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Packet{
		Title:   "the title",
		Data:    999,
		Tag:     "my_tag",
		NoTrace: true,
	}.Error()
	output := buf.String()
	assert.Contains(output, `level=error`)
	assert.Contains(output, `msg="the title"`)
	assert.Contains(output, `tag=my_tag`)
	assert.Contains(output, `value=999`)
	trace := GetTraces(0, 2)
	assert.NotContains(output, trace[0].Function)
}
