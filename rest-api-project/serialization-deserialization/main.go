package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		Name:  "Alice",
		Email: "alicetest@email.com",
	}

	fmt.Println(user)

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	var user1 User

	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created from jsonData: ", user1)

	jsonData1 := `{"name": "John", "email": "johntest@email.com"}`
	reader := strings.NewReader(jsonData1)
	decoder := json.NewDecoder(reader)

	var user2 User

	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created from decoder: ", user2)

	var buf bytes.Buffer
	// encoder := json.NewEncoder(&buf)

	// err = encoder.Encode(user)
	err = json.NewEncoder(&buf).Encode(user) // This is the prefer way of doing it
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encoded json string: ", buf.String())
	fmt.Println(user1.Name)
}

// json.Marshal(user)
// Takes the user struct and outputs raw JSON bytes directly into memory.
//
// json.Unmarshal(jsonData, &user1)
// Takes JSON bytes already in memory and parses them into user1.
//
// json.NewDecoder(reader).Decode(&user2)
// Reads a continuous stream of data (in your code, from a strings.NewReader) and converts it directly into user2 as it reads.
//
// json.NewEncoder(&buf).Encode(user)
// Converts user into JSON and writes it directly to an output stream (in your code, a bytes.Buffer).

// Rule of Thumb
// Use Marshal / Unmarshal when working with strings or byte arrays in local code.
//
// Use Encode / Decode when working with HTTP APIs, files, or socket streams (http.Request.Body, http.ResponseWriter, os.File).
