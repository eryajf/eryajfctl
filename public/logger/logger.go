package logger

import (
	"os"
	"sync"

	"github.com/charmbracelet/log"
)

var Logger *log.Logger
var once sync.Once

func InitLogger(level string) {
	once.Do(func() {
		Logger = log.NewWithOptions(os.Stderr, log.Options{ReportTimestamp: true})
	})
	if level == "debug" {
		Logger.SetLevel(log.DebugLevel)
	} else {
		Logger.SetLevel(log.InfoLevel)
	}
}

func Info(args ...any) {
	Logger.Info(args)
}

func Warning(args ...any) {
	Logger.Warn(args)
}

func Debug(args ...any) {
	Logger.Debug(args)
}

func Error(args ...any) {
	Logger.Error(args)
}

func Fatal(args ...any) {
	Logger.Fatal(args)
}
