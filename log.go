package klog

import (
	"fmt"
	stdLog "log"
	"strings"
	"time"
)

// Level of debug
type Level int32

/**
 * Define log levels and formats
 */
const (
	Verbose Level = iota
	Debug
	Info
	Warn
	Error
	level      = LogLevel
	LogFormat  = "[%s]:[%s]"
	LogFormatF = "[%s]:[%s] %s\n"
)

var (
	levelStrings = []string{"V", "D", "I", "W", "E"}
	// LogBufferLength buffer size of log message
	LogBufferLength = 32
	fileLogger      = NewFileLogWriter(FileLogPath)
)

// LogRecord log message struct
type LogRecord struct {
	Created time.Time // The time at which the log message created
	Message string    // The log message
}

// LogWriter is an interface for anything that should be able to write logs
type LogWriter interface {
	LogWrite(rec *LogRecord)
	Close()
}

// V Verbose log
func V(tag string, args ...interface{}) {
	if level <= Verbose {
		log(Verbose, tag, args...)
	}
}

// Vf Verbose formatted log
func Vf(tag string, format string, args ...interface{}) {
	if level <= Verbose {
		logf(Verbose, tag, format, args...)
	}
}

// D Debug log
func D(tag string, args ...interface{}) {
	if level <= Debug {
		log(Debug, tag, args...)
	}
}

// Df Debug formatted log
func Df(tag string, format string, args ...interface{}) {
	if level <= Debug {
		logf(Debug, tag, format, args...)
	}
}

// I Info log
func I(tag string, args ...interface{}) {
	if level <= Info {
		log(Info, tag, args...)
	}
}

// If Info formatted log
func If(tag string, format string, args ...interface{}) {
	if level <= Info {
		logf(Info, tag, format, args...)
	}
}

// W Warn log
func W(tag string, args ...interface{}) {
	if level <= Warn {
		log(Warn, tag, args...)
	}
}

// Wf Warn formatted log
func Wf(tag string, format string, args ...interface{}) {
	if level <= Warn {
		logf(Warn, tag, format, args...)
	}
}

// E Error log
func E(tag string, args ...interface{}) {
	if level <= Error {
		log(Error, tag, args...)
	}
}

// Ef Error formatted log
func Ef(tag string, format string, args ...interface{}) {
	if level <= Error {
		logf(Error, tag, format, args...)
	}
}

func log(level Level, tag string, args ...interface{}) {
	format := fmt.Sprintf(LogFormat, levelStrings[level], tag)
	msg := fmt.Sprintf(format+strings.Repeat(" %v", len(args))+"\n", args...)
	stdLog.Printf(msg)
	if FileLogEnable {
		fileLogger.LogWrite(&LogRecord{Created: time.Now(), Message: msg})
	}

}

/**
 * Log with custom msg format
 */
func logf(level Level, tag string, format string, args ...interface{}) {
	content := fmt.Sprintf(format, args...)
	msg := fmt.Sprintf(LogFormatF, levelStrings[level], tag, content)
	stdLog.Print(msg)
	if FileLogEnable {
		fileLogger.LogWrite(&LogRecord{Created: time.Now(), Message: msg})
	}

}
