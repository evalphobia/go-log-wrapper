package log

import (
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

// NewPacket creates Packet with Logger.
func (p Logger) NewPacket() Packet {
	return Packet{
		Logger: p,
	}
}
