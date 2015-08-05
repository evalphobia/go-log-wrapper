package log

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type StdoutMock struct {
	stdout *os.File
	stderr *os.File
	writer *os.File
	output chan string
}

func (m *StdoutMock) Set() {
	backupOut := os.Stdout
	backupErr := os.Stderr
	r, w, _ := os.Pipe()

	os.Stdout = w
	os.Stderr = w
	opChan := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		opChan <- buf.String()
	}()

	m.stdout = backupOut
	m.stderr = backupErr
	m.writer = w
	m.output = opChan
}

func (m *StdoutMock) Get() string {
	m.writer.Close()
	os.Stdout = m.stdout
	os.Stderr = m.stderr
	return <-m.output
}

func TestNewLogField(t *testing.T) {
	assert := assert.New(t)

	f := newLogField([]interface{}{"message from me"})
	assert.Equal("message from me", f["value"])

	f = newLogField([]interface{}{"the title", "message from me"})
	assert.Equal("the title", f["message"])
	assert.Equal("message from me", f["value"])
}

func TestError(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Error("message from me")
	output := buf.String()
	assert.Contains(output, `level=error`)
	assert.Contains(output, `value="message from me"`)
}

func TestWarn(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Warn("message from me")
	output := buf.String()
	assert.Contains(output, `level=warn`)
	assert.Contains(output, `value="message from me"`)
}

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Info("message from me")
	output := buf.String()
	assert.Contains(output, `level=info`)
	assert.Contains(output, `value="message from me"`)
}

func TestDebug(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	Debug("message from me")
	output := buf.String()
	assert.Empty(output)

	logrus.SetLevel(logrus.DebugLevel)
	Debug("message from me")
	output = buf.String()
	assert.Contains(output, `level=debug`)
	assert.Contains(output, `value="message from me"`)
}

func TestDump(t *testing.T) {
	assert := assert.New(t)

	m := &StdoutMock{}
	m.Set()

	Dump("message from me")

	op := m.Get()
	assert.Contains(op, `(string) (len=15) "message from me"`)
}

func TestPrint(t *testing.T) {
	assert := assert.New(t)

	m := &StdoutMock{}
	m.Set()

	Print("message from me")

	op := m.Get()
	assert.Contains(op, "message from me")
}

func TestHeader(t *testing.T) {
	assert := assert.New(t)

	m := &StdoutMock{}
	m.Set()

	Header()
	op := m.Get()
	assert.Contains(op, "=============================================\n")

	m.Set()
	Header("message header")
	op = m.Get()
	assert.Contains(op, "===================== message header =====================\n")
}
