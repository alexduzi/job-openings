package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug    *log.Logger
	info     *log.Logger
	warnning *log.Logger
	err      *log.Logger
	writer   io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:    log.New(writer, "DEBUG ", logger.Flags()),
		info:     log.New(writer, "INFO ", logger.Flags()),
		warnning: log.New(writer, "WARNING ", logger.Flags()),
		err:      log.New(writer, "ERROR ", logger.Flags()),
		writer:   writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warnning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warnning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
