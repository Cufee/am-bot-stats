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
