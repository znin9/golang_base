package main

import (
	"log"
	"os"
)

func main() {
	log.Println("this log is output by std variable by package log")

	logFile, err := os.OpenFile("log_pkg.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, "Output =>:", log.LstdFlags)
	logger.Println("this log message is output into log file")
}
