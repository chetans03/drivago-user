package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitializeLogging() error {
	fmt.Println("F01: Initializing logger")
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel("info")
	if err != nil {
		return fmt.Errorf("F02: Failed to parse log level: %v", err)
	}
	Logger.SetLevel(level)
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: false,
		TimestampFormat:  time.RFC3339Nano,
		DisableSorting:   true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
		},
	})
	return nil
}
