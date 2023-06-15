package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jimlawless/whereami"
)

const (
	INFO  = "Info"
	ERROR = "Error"
)

func PrintLog(debugLevel string, args ...interface{}) (err error) {
	var logFile *os.File

	// creating file name based on current date
	fileName := "logs/log-" + time.Now().UTC().Format("2006-01-02") + ".log"
	// opening current log file
	if logFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		fmt.Println("Error in opening log file", err)
		return
	}
	// appending new log inside current log file
	logger := log.New(logFile, debugLevel, log.Ldate|log.Ltime|log.Lshortfile)
	args = append(args, whereami.WhereAmI(3))
	logger.Println(args...)

	// closing file before return
	defer logFile.Close()
	return
}

// InfoLog ...
// function will receive information related logs to print inside log file
func InfoLog(args ...interface{}) {
	// call for print log
	PrintLog(INFO, args...)
}

// ErrorLog ...
// function will receive error related logs to print inside log file
func ErrorLog(args ...interface{}) {
	// call for print log
	PrintLog(ERROR, args...)
}
