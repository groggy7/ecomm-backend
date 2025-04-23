package pkg

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "INFO - ", log.Lshortfile)
var ErrorLogger = log.New(os.Stderr, "ERROR - ", log.Lshortfile)
var DebugLogger = log.New(os.Stdout, "DEBUG - ", log.Lshortfile)
var WarningLogger = log.New(os.Stdout, "WARNING - ", log.Lshortfile)
