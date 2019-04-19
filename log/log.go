package log

import (
	//"flag"
	"fmt"
	slog "github.com/sirupsen/logrus"
)

var (
	logger   Logger
	logLevel = InfoLevel
)

func init() {
	slog.SetFormatter(&slog.TextFormatter{})
}

func GetLogger() *Logger {
	return &logger
}

func GetLevel() Level {
	return logLevel
}

func SetLevel(level Level) {
	logLevel = level
}

func Debug(msg string) {
	log(DebugLevel, msg)
}

func Info(msg string) {
	log(InfoLevel, msg)
}

func Warn(msg string) {
	log(WarnLevel, msg)
}

func Error(msg string) {
	log(ErrorLevel, msg)
}

func Fatal(msg string) {
	log(FatalLevel, msg)
}

func Debugf(msg string, v ...interface{}) {
	log(DebugLevel, fmt.Sprintf(msg, v...))
}

func Infof(msg string, v ...interface{}) {
	log(InfoLevel, fmt.Sprintf(msg, v...))
}

func Warnf(msg string, v ...interface{}) {
	log(WarnLevel, fmt.Sprintf(msg, v...))
}

func Errorf(msg string, v ...interface{}) {
	log(ErrorLevel, fmt.Sprintf(msg, v...))
}

func Fatalf(msg string, v ...interface{}) {
	log(FatalLevel, fmt.Sprintf(msg, v...))
}

func log(level Level, message string) {
	if level < logLevel {
		return
	}
	id, _ := sf.NextID()
	e := NewEntry(id, level, message)

	if logger.Log == nil {
		switch level {
		case DebugLevel:
			slog.Debug(message)
		case InfoLevel:
			slog.Info(message)
		case WarnLevel:
			slog.Warn(message)
		case ErrorLevel:
			slog.Error(message)
		case FatalLevel:
			slog.Fatal(message)
		}
		return
	}

	logger.Log.Log(e)
}
