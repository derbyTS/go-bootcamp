package main

import (
	"fmt"
	"os"
)

func main() {
	// not giving it a directory will cause a creation of file in os temporary directory sample: /var/folders/5t/8ztqx2dj4l3dmwm3mcbsr_xh0000gn/T/temporaryFile59059905
	file, err := os.CreateTemp("", "temporaryFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Temporary file: ", file.Name())
	defer os.Remove(file.Name())
	defer file.Close()

	tempDir, err := os.MkdirTemp("", "GoTempDir")
	if err != nil {
		fmt.Println(err)
	}

	defer os.RemoveAll(tempDir)
	fmt.Println("Temporary directory: ", tempDir)
}
