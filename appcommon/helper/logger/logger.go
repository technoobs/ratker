// The logging utilities for the application

package logger

import (
	"io"
	"log"
)

type IAppLogger interface {
	WriteInfo(context string)
	WriteWarn(context string)
	WriteError(context string)
}

type AppLogger struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

func (logger *AppLogger) WriteInfo(context string) {
	logger.Info.Println(context)
}

func (logger *AppLogger) WriteWarn(context string) {
	logger.Warn.Println(context)
}

func (logger *AppLogger) WriteError(context string) {
	logger.Error.Panicln(context)
}

// // Setup the AppLogger
// func (logger AppLogger) Setup(logPath string) *AppLogger {
// 	if _, err := os.Stat(logPath); os.IsNotExist(err) {
// 		file, createErr := os.Create(logPath)
// 		if createErr != nil {
// 			// Place to return error to system
// 			fmt.Printf("Failed to open log file: %v\n", createErr)
// 			os.Exit(1)
// 		}
// 	} else {
// 		fmt.Println("Found file......")
// 		file, openErr := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
// 		if openErr != nil {
// 			fmt.Println("Opening file failed")
// 			fmt.Println(openErr)
// 		}
// 		if openErr == nil {
// 			Info, Warning, Error = prepareLogger(logPath)
// 		}
// 	}
// 	return &logger
// }

// prepareLogger prepares loggers for Info, Warning, Error
func prepareLogger(writer io.Writer) *AppLogger {
	infoLogger := log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(writer, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger := log.New(writer, "*ERROR*: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &AppLogger{infoLogger, warnLogger, errLogger}
}
