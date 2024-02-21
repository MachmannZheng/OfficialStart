package logfunc

import (
	"log"
	"os"
)

func BasicLogger() {
	myLogger := log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	myLogger.Println("this is the logger message from BasicLogger function")
}
