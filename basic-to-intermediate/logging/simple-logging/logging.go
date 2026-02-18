package main

import (
	"log"
	"os"
)

func main() {
	log.Println("First log")

	log.SetPrefix("INFO: ")
	log.Println("This is INFO")

	log.SetFlags(log.Ldate)
	log.Println("This is a log message with date only YYYY/MM/DD")

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("This is a log message with date and time")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time and file")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Llongfile)
	log.Println("This is a log message with date, time, short file and long file")

	infoLogger.Println("INFO log test")
	warnLogger.Println("WARN log test")
	errorLogger.Println("ERROR log test")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is debug message")
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
