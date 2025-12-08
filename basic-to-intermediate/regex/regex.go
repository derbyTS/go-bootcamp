package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Sample of backtick \n ")
	fmt.Println(`Sample of backtick \n `)

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "user"

	fmt.Println("Email1 : ", re.MatchString(email1))
	fmt.Println("Email2 : ", re.MatchString(email2))

	// Capturing groups
	// Compile a regex pattern to capture date components

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	ymd1 := "2025-31-12"

	subMatches := re.FindStringSubmatch(ymd1)

	fmt.Println(subMatches)
	fmt.Println(subMatches[0])
	fmt.Println(subMatches[1])
	fmt.Println(subMatches[2])
	fmt.Println(subMatches[3])

	re = regexp.MustCompile(`[aeiou]`)

	fmt.Println(re.ReplaceAllString("Hello Golang", "*"))

	// Flags
	// i = case insensitive
	// m = multi line model
	// s = dot matches all

	re = regexp.MustCompile(`(?i)go`) // Searching for the string "go"
	str1 := "Golang is going good so far"
	str2 := "Java is greate"

	fmt.Println("Match: ", re.MatchString(str1))
	fmt.Println("Match: ", re.MatchString(str2))
}
