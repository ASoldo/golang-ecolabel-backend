// Package logger provides a simple logging utility for the golang-ecolabel-backend project.
// It defines two loggers: Info and Error, which log to stdout and stderr respectively.
package logger

import (
	"log"
	"os"
)

// Info is a logger that logs informational messages to stdout.
var Info *log.Logger

// Error is a logger that logs error messages to stderr.
var Error *log.Logger

// init initializes the Info and Error loggers when the package is imported.
func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
