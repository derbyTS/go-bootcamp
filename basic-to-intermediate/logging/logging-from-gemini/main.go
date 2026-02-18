package main

import (
	"log"

	"logger/custom-logger"
)

func main() {
	closeLog, err := customlogger.Init("app.log")
	if err != nil {
		log.Fatal(err)
	}

	defer closeLog()

	customlogger.Debug.Println("Debug test ...")
	customlogger.Info.Println("Info test ...")
	customlogger.Warn.Println("Warn test ...")
	customlogger.Error.Println("Error test ...")
}
