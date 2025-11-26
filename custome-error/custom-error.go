package main

import (
	"errors"
	"fmt"
)

func main() {
	err := closeFileMock()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type customError struct {
	code    int
	message string
	err     error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s %v\n", e.code, e.message, e.err)
}

func closeFileMock() error {
	err := saveFileMock()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something happen!",
			err:     err,
		}
	}
	return nil
}

func saveFileMock() error {
	return errors.New("Something wrong saving file")
}
