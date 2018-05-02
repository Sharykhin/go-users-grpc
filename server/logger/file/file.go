package file

import (
	"io"
	"log"
	"os"
)

type (
	logger struct {
		Log *log.Logger
	}
)

// Logger is a logger that satisfies Logger interface and depending on environment may put logs into a file
var Logger logger

func init() {
	f, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	Logger = logger{Log: log.New(io.MultiWriter(f, os.Stdout), "", log.LstdFlags|log.Lshortfile)}

}

func (f logger) Errorf(format string, args ...interface{}) {
	f.Log.Printf(format, args...)
}
