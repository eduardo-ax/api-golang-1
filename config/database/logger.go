package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, p+": DEBUG: ", logger.Flags()),
		info:    log.New(writer, p+": INFO: ", logger.Flags()),
		warning: log.New(writer, p+": WARNING: ", logger.Flags()),
		err:     log.New(writer, p+": ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Warm(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.debug.Println(v...)
}

// Format Enable Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Warmf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
