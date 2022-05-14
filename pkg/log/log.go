package log

import (
	"github.com/JIeeiroSst/store/config"
	logrus_logstash "github.com/sima-land/logrus-logstash-hook"

	"github.com/sirupsen/logrus"
)

type Log interface {
	Trace(msg interface{})
	Debug(msg interface{})
	Info(msg interface{})
	Warn(msg interface{})
	Error(msg interface{})
	Fatal(msg interface{})
	Panic(msg interface{})
}

type log struct {
	config config.Config
	logrus logrus.Logger
}

func NewLog(config config.Config) Log {
	logrus := logrus.New()
	hook, err := logrus_logstash.NewHook(config.Logstash.Tranport, config.Logstash.Host, config.Logstash.NameApp)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Hooks.Add(hook)

	return &log{
		logrus: *logrus,
		config: config,
	}
}

func (l *log) Trace(msg interface{}) {
	l.logrus.Trace(msg)
}

func (l *log) Debug(msg interface{}) {
	l.logrus.Debug(msg)
}

func (l *log) Info(msg interface{}) {
	l.logrus.Info(msg)
}

func (l *log) Warn(msg interface{}) {
	l.logrus.Warn(msg)
}

func (l *log) Error(msg interface{}) {
	l.logrus.Error(msg)
}

func (l *log) Fatal(msg interface{}) {
	l.logrus.Fatal(msg)
}

func (l *log) Panic(msg interface{}) {
	l.logrus.Panic(msg)
}
