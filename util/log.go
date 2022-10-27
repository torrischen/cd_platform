package util

import (
	"go.uber.org/zap"
)

type logger struct {
	l *zap.Logger
}

var Logger logger

func InitLogger() {
	log, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("Fail to init Logger!")
	}

	Logger.l = log
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.l.Sugar().Errorf(template, args)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.l.Sugar().Infof(template, args)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.l.Sugar().Warnf(template, args)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.l.Sugar().Fatalf(template, args)
}
