package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// err := os.Mkdir("data", 0755)
	// checkError(err)
	// // defer os.RemoveAll("data")
	//
	// os.WriteFile("data/sample.txt", []byte("Hello Word"), 0755)
	//
	// err = os.MkdirAll("data1/test/a/b/sample.txt", 0755)
	// checkError(err)

	// os.RemoveAll("data")
	// os.RemoveAll("data1") //os.Remove will not remove a directory if there is a file inside and you have to remove the file using the exact path

	// err := os.MkdirAll("subdir/parent/child1", 0755)
	// checkError(err)
	// err = os.MkdirAll("subdir/parent/child2", 0755)
	// checkError(err)
	// err = os.MkdirAll("subdir/parent/child3", 0755)
	// checkError(err)
	// err = os.MkdirAll("subdir/parent/child4", 0755)
	// checkError(err)
	// os.WriteFile("subdir/parent/child1/file", []byte("hello file"), 0775)
	// os.WriteFile("subdir/parent/child2/file", []byte("hello file"), 0775)
	// os.WriteFile("subdir/parent/file", []byte("hello file"), 0775)

	result, err := os.ReadDir("subdir/parent")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range result {
		fmt.Printf(
			"entry: %s, entry.Name: %s, entry.IsDir: %t, entry.Type: %s\n",
			entry,
			entry.Name(),
			entry.IsDir(),
			entry.Type(),
		)
	}

	checkError(os.Chdir("subdir/parent/child1"))

	readDirectory, err := os.ReadDir(".")
	checkError(err)

	fmt.Println(
		"************************************************************************************************",
	)
	for _, entry := range readDirectory {
		fmt.Printf(
			"entry: %s, entry.Name: %s, entry.IsDir: %t, entry.Type: %s\n",
			entry,
			entry.Name(),
			entry.IsDir(),
			entry.Type(),
		)
	}

	checkError(os.Chdir("../")) // Dont forget it cause an error in WalkDir

	readDirectory1, err := os.ReadDir(".")
	checkError(err)

	currentDir, err := os.Getwd()
	checkError(err)
	fmt.Println("Current Directory: ", currentDir)
	checkError(os.Chdir("../../")) // I cd again to parent folder

	fmt.Println(
		"************************************************************************************************",
	)
	for _, entry := range readDirectory1 {
		fmt.Printf(
			"entry: %s, entry.Name: %s, entry.IsDir: %t, entry.Type: %s\n",
			entry,
			entry.Name(),
			entry.IsDir(),
			entry.Type(),
		)
	}

	fmt.Println(
		"**************************************walkdir**********************************************************",
	)

	pathFile := "subdir/parent/child1"
	pathFile = strings.TrimSpace(pathFile)
	err = filepath.WalkDir(pathFile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
		}
		fmt.Println("Successfully visited: ", path)
		return nil
	})
	if err != nil {
		fmt.Println("Error from WalkDir: ", err)
	}
}

func checkError(err error) {
	if err != nil {
		// panic(err) //Panic is common use here I guess
		fmt.Println(err)
	}
}
