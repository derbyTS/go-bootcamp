package customlogger

import (
	"io"
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func Init(debugFilePath string) (func() error, error) {
	file, err := os.OpenFile(debugFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}

	flags := log.Ldate | log.Ltime | log.Lshortfile
	debugOutput := io.MultiWriter(file, os.Stdout)

	// Debug = log.New(file, "DEBUG: ", flags) // This will only output in file
	Debug = log.New(debugOutput, "DEBUG: ", flags)
	Info = log.New(os.Stdout, "INFO: ", flags)
	Warn = log.New(os.Stdout, "WARN: ", flags)
	Error = log.New(os.Stdout, "ERROR: ", flags)

	return file.Close, nil
}
