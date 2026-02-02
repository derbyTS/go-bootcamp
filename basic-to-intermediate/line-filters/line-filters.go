package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	keyword := "brown"

	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			// if strings.EqualFold(
			// 	line,
			// 	keyword,
			// ) {
			updatedLine := strings.ReplaceAll(line, keyword, "green")
			fmt.Printf("%d %s \n", lineNumber, line)
			fmt.Printf("%d %s \n", lineNumber, updatedLine)
			lineNumber++
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
