package utils

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

// InitLogger initializes loggers for info, warning, and error levels
func InitLogger() {
	infoLogFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open info log file: %v", err)
	}
	warnLogFile, err := os.OpenFile("warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open warn log file: %v", err)
	}
	errorLogFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}

	Info = log.New(infoLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(warnLogFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
