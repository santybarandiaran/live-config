package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// ServiceLogger holds the implementation for the logger interface
type ServiceLogger struct {
	Logger *log.Logger
}

var (
	// Instance provides methods to do all the logging in different levels
	Instance = newLogger()
)

// New gets a ServiceLogger instance with standard configuration
func newLogger() *ServiceLogger {
	Logger := log.New()
	var serviceLogger ServiceLogger
	// Log as JSON instead of the default ASCII formatter.
	Logger.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	Logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	Logger.SetLevel(log.InfoLevel)

	serviceLogger.Logger = Logger

	return &serviceLogger
}

// Error logs error level messages
func (l ServiceLogger) Error(message string) {
	l.Logger.Error(message)
}

// Info logs info level messages
func (l ServiceLogger) Info(message string) {
	l.Logger.Info(message)
}

// Warn logs warn level messages
func (l ServiceLogger) Warn(message string) {
	l.Logger.Warn(message)
}
