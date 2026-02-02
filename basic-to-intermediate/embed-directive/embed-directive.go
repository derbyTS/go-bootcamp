package main

import (
	"embed"
	"fmt"
	"io/fs"

	// This is a blank import
	_ "side-effect/helloer"
	// _ "embed" // we use blank import for embed if we only want to embed file and we don't need to use functions inside embed
)

//go:embed example/example.txt

var content string

//go:embed example
var basicFolder embed.FS

func main() {
	fmt.Println("Embeded content: ", content)
	folderContent, err := basicFolder.ReadFile("example/print.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("folderContent: ", string(folderContent))

	err = fs.WalkDir(basicFolder, "example", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error in walkdir: ", err)
		}
		fmt.Println("Success walkdir path: ", path)
		return nil
	})
	if err != nil {
		fmt.Println("Error in walkdir outside error: ", err)
	}
}
