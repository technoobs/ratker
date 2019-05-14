// The logging utilities for the application

// TODO: kick off a process to monitor log file size, once it's
// over the limit, create a new log file, and compress
// the existing log file at the same time. byte.buffer[]

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	LogBackUpSuffix = "2006-01-02_15-04-05.99"
)

type IAppLogger interface {
	WriteInfo(context string)
	WriteWarn(context string)
	WriteError(context string)
}

type AppLogger struct {
	info *log.Logger
	warn *log.Logger
	err  *log.Logger
}

func (logger *AppLogger) WriteInfo(context string) {
	logger.info.Println(context)
}

func (logger *AppLogger) WriteWarn(context string) {
	logger.warn.Println(context)
}

func (logger *AppLogger) WriteError(context string) {
	logger.err.Panicln(context)
}

// Setup the AppLogger
func Setup(logPath string) *AppLogger {
	file, createErr := os.Create(logPath)
	if createErr != nil {
		fmt.Printf("Failed to setup log file: %s", createErr.Error())
		os.Exit(1)
	}
	logger := prepareLogger(file)
	return &logger
}

// prepareLogger prepares loggers for Info, Warning, Error
func prepareLogger(writer io.Writer) AppLogger {
	infoLogger := log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(writer, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger := log.New(writer, "*ERROR*: ", log.Ldate|log.Ltime|log.Lshortfile)
	return AppLogger{infoLogger, warnLogger, errLogger}
}

// generateLogFile function is potentialy used for new log file
func generateLogFile() {

}

// createLogFileName function creates the log file name by following
// the name rule: log_YYYY-MM-DD_HH-MM-SS-SSS.log
func createLogFileName() string {
	now := time.Now().UTC()
	logTime := now.Format(LogBackUpSuffix)
	return "log_" + logTime + ".log"
}
