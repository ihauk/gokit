package logger

import (
	"log"
	"os"
)

var (
	ErrorLog = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLog  = log.New(os.Stderr, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
)
