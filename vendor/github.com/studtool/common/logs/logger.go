package logs

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: func() *logrus.Logger {
			log := logrus.StandardLogger()
			log.SetFormatter(&logrus.JSONFormatter{})
			return log
		}(),
	}
}

type LogFields struct {
	Service   string
	Component string
	Function  string
}

func (log *Logger) Debug(f *LogFields, args ...interface{}) {
	log.logger.WithFields(log.mapFields(f)).Debug(args...)
}

func (log *Logger) Info(f *LogFields, args ...interface{}) {
	log.logger.WithFields(log.mapFields(f)).Info(args...)
}

func (log *Logger) Warning(f *LogFields, args ...interface{}) {
	log.logger.WithFields(log.mapFields(f)).Warn(args...)
}

func (log *Logger) Error(f *LogFields, args ...interface{}) {
	log.logger.WithFields(log.mapFields(f)).Error(args...)
}

func (log *Logger) Fatal(f *LogFields, args ...interface{}) {
	log.logger.WithFields(log.mapFields(f)).Fatal(args...)
}

func (log *Logger) mapFields(f *LogFields) logrus.Fields {
	return logrus.Fields{
		"service":   f.Service,
		"component": f.Component,
		"function":  f.Function,
	}
}
