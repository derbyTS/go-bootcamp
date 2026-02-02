package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./data/file.txt"
	absolutPath := "~/Development/Gogogo/go-bootcamp/basic-to-intermediate/file-path/"

	joinedPath := filepath.Join(absolutPath, relativePath)

	fmt.Println(joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")

	fmt.Println(normalizedPath)

	dir, file := filepath.Split(joinedPath)

	fmt.Println(dir)
	fmt.Println(file)
	fmt.Println(filepath.Base(joinedPath))

	fmt.Println("Is absolute path: ", filepath.IsAbs(relativePath))
	fmt.Println("Is absolute path: ", filepath.IsAbs(absolutPath))
	fmt.Println("Extension: ", filepath.Ext(file))
	fmt.Println("TrimSuffix: ", strings.TrimSuffix(file, filepath.Ext(file)))
	rel, err := filepath.Rel(absolutPath, joinedPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Return relative path: ", rel)

	rel, err = filepath.Rel("a/b/c", joinedPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Return relative path: ", rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(absPath)
}
