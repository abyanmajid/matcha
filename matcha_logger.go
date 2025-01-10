package matcha

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var infoColor = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
var debugColor = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
var errorColor = color.New(color.FgRed).Add(color.Bold).SprintFunc()

func syslogToStdout(level, format string, args ...any) {
	logMessage := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02T15:04:05Z07:00")
	fmt.Fprintf(os.Stdout, "%s [%s] %s\n", timestamp, level, logMessage)
}

// Logger represents a structured logger with levels: Info, Debug, and Error.
type Logger struct{}

// Log information at the INFO severity level
func (l *Logger) Info(format string, args ...any) {
	syslogToStdout(infoColor("INFO"), format, args...)
}

// Log information at the DEBUG severity level
func (l *Logger) Debug(format string, args ...any) {
	syslogToStdout(debugColor("DEBUG"), format, args...)
}

// Log information at the ERROR severity level
func (l *Logger) Error(format string, args ...any) {
	syslogToStdout(errorColor("ERROR"), format, args...)
}
