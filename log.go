package klog

import (
	"fmt"
	stdLog "log"
	"strings"
	"time"
)

type Level int32

const (
	VERBOSE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	level        = LOG_LEVEL
	LOG_FORMAT   = "[%s]:[%s]"
	LOG_FORMAT_F = "[%s]:[%s] %s\n"
)

var (
	levelStrings    = []string{"V", "D", "I", "W", "E"}
	LogBufferLength = 32
	fileLogger      = NewFileLogWriter(FILE_LOG_PATH)
)

/**
 */
type LogRecord struct {
	Created time.Time // The time at which the log message created
	Message string    // The log message
}

/**
 * This is an interface for anything that should be able to write logs
 */
type LogWriter interface {
	LogWrite(rec *LogRecord)
	Close()
}

func V(tag string, args ...interface{}) {
	if level <= VERBOSE {
		log(VERBOSE, tag, args...)
	}
}

func Vf(tag string, format string, args ...interface{}) {
	if level <= VERBOSE {
		logf(VERBOSE, tag, format, args...)
	}
}

func D(tag string, args ...interface{}) {
	if level <= DEBUG {
		log(DEBUG, tag, args...)
	}
}

func Df(tag string, format string, args ...interface{}) {
	if level <= DEBUG {
		logf(DEBUG, tag, format, args...)
	}
}

func I(tag string, args ...interface{}) {
	if level <= INFO {
		log(INFO, tag, args...)
	}
}

func If(tag string, format string, args ...interface{}) {
	if level <= INFO {
		logf(INFO, tag, format, args...)
	}
}

func W(tag string, args ...interface{}) {
	if level <= WARN {
		log(WARN, tag, args...)
	}
}

func Wf(tag string, format string, args ...interface{}) {
	if level <= WARN {
		logf(WARN, tag, format, args...)
	}
}

func E(tag string, args ...interface{}) {
	if level <= ERROR {
		log(ERROR, tag, args...)
	}
}

func Ef(tag string, format string, args ...interface{}) {
	if level <= ERROR {
		logf(ERROR, tag, format, args...)
	}
}

func log(level Level, tag string, args ...interface{}) {
	format := fmt.Sprintf(LOG_FORMAT, levelStrings[level], tag)
	msg := fmt.Sprintf(format+strings.Repeat(" %v", len(args))+"\n", args...)
	stdLog.Printf(msg)
	if FILE_LOG_ENABLE {
		fileLogger.LogWrite(&LogRecord{Created: time.Now(), Message: msg})
	}

}

/**
 * Log with custom msg format
 */
func logf(level Level, tag string, format string, args ...interface{}) {
	content := fmt.Sprintf(format, args...)
	msg := fmt.Sprintf(LOG_FORMAT_F, levelStrings[level], tag, content)
	stdLog.Print(msg)
	if FILE_LOG_ENABLE {
		fileLogger.LogWrite(&LogRecord{Created: time.Now(), Message: msg})
	}

}
