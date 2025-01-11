package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var infoColor = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
var debugColor = color.New(color.FgHiMagenta).Add(color.Bold).SprintFunc()
var traceColor = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
var errorColor = color.New(color.FgRed).Add(color.Bold).SprintFunc()

func syslogToStdout(level, format string, args ...any) {
	logMessage := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Fprintf(os.Stdout, "%s [%s] %s\n", timestamp, level, logMessage)
}

// Log information at the INFO severity level
func Info(format string, args ...any) {
	syslogToStdout(infoColor("INFO"), format, args...)
}

// Log information at the DEBUG severity level
func Debug(format string, args ...any) {
	syslogToStdout(debugColor("DEBUG"), format, args...)
}

// Log information at the DEBUG severity level
func Trace(format string, args ...any) {
	syslogToStdout(traceColor("TRACE"), format, args...)
}

// Log information at the ERROR severity level
func Error(format string, args ...any) {
	syslogToStdout(errorColor("ERROR"), format, args...)
}

func Fatal(format string, args ...any) {
	syslogToStdout(errorColor("FATAL: TERMINATING"), format, args...)
	os.Exit(1)
}
