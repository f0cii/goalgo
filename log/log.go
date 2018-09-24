package log

import (
	//"flag"
	"fmt"
)

var (
	logger Logger
)

func GetLogger() *Logger {
	return &logger
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
	id, _ := sf.NextID()
	e := NewEntry(id, level, message)
	logger.Log.Log(e)
}
