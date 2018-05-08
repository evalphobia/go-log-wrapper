package log

import (
	"io"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// Logger is wrapper struct of *logrus.Logger.
type Logger struct {
	*logrus.Logger
}

// NewLogger returns initialized *Logger.
func NewLogger() *Logger {
	return &Logger{
		Logger: logrus.New(),
	}
}

// SetOutput sets output writer.
func (l *Logger) SetOutput(w io.Writer) {
	l.Logger.Out = w
}

// DisableOutput disables output.
func (l *Logger) DisableOutput() {
	l.SetOutput(ioutil.Discard)
}

// SetFormatter sets Fomatter.
func (l *Logger) SetFormatter(f logrus.Formatter) {
	l.Logger.Formatter = f
}

// SetLogLevel sets log level.
func (l *Logger) SetLogLevel(lv logrus.Level) {
	l.Logger.SetLevel(lv)
}

// NewPacket creates Packet with Logger.
func (l *Logger) NewPacket() Packet {
	return Packet{
		Logger: l,
	}
}
