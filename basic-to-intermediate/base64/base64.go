package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, Base64 Encoding")

	fmt.Printf("%T\n", data)
	// Encode
	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Println(encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", decoded)
	fmt.Println(string(decoded))

	// URL safe avoid '/' and '+'; meaning any encoded will have no '/' and '+'
	urlSafeEncoding := base64.URLEncoding.EncodeToString(data)
	fmt.Println(urlSafeEncoding)
}
