package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Printf("This is password %s\n", password)
	fmt.Printf("This is hash %d\n", hash)

	fmt.Printf("hash hex value is %x\n", hash)
	fmt.Printf("hash512 hex value is %x\n", hash512)

	// To avoid Rainbow Table Attack vs Dictionary Attack we use Salting
	fmt.Println("------------------salt--------------------")
	salt, err := generateSalt()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("salt hex value is %x\n", salt)
	fmt.Printf("Salt String %s\n", salt)

	// Simulating that we insert it in database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	generateHashPass := hashPassword(password, salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Hash: ", generateHashPass)

	// Simulating verify password
	signUpHash := hashPassword("password1234", salt)
	if err != nil {
		fmt.Println(err)
		return
	}
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	loginHash := hashPassword(password, decodedSalt)
	if signUpHash == loginHash {
		fmt.Println("Passsword is correct")
	} else {
		fmt.Println("Login failed")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	// _, err := io.ReadFull(rand.Reader, salt)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
