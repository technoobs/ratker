package logger

// The logging utilities for the application

// IAppLogger is the interface for logger function
type IAppLogger interface {
	Info()
	Warn()
	Error()
}
