package utils

import (
	"log"
	"os"
)

// setup server logging for errors
func InitLogging() {
    file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }
    log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

