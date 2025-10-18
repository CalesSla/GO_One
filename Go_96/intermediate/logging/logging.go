package main

import (
	"log"
	"os"
)

func main() {

	log.Println("This is a log message.")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message.")

	log.SetFlags(log.Ldate)
	log.Println("This is a log message with date.")

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("This is a log message with date and time.")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time, and file info.")

	infoLogger.Println("This is an info log message.")
	warnLogger.Println("This is a warning log message.")
	errorLogger.Println("This is an error log message.")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: ", err)
	}
	defer file.Close()

	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug message.")
	warnLogger1.Println("This is a warning log message written to file.")
	errorLogger1.Println("This is an error log message written to file.")
	infoLogger1.Println("This is an info log message written to file.")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
