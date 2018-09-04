package notion

import "github.com/sirupsen/logrus"

// Logger defines the logger type this package uses.
type Logger interface {
	WithField(key string, value interface{}) Logger
	WithError(err error) Logger

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
}

// WrapLogrus wraps a logrus Logger to conform to the Logger interface defined in this package.
type WrapLogrus struct {
	logrus.FieldLogger
}

// WithField attaches a key-value pair to a log line.
func (wl WrapLogrus) WithField(key string, value interface{}) Logger {
	return &WrapLogrus{wl.FieldLogger.WithField(key, value)}
}

// WithError attaches a key-value pair to a log line.
func (wl WrapLogrus) WithError(err error) Logger {
	return wl.WithError(err)
}
