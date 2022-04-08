package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}
