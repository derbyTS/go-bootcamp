package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func VerifyPassword(encodedHash, password string) error {
	parts := strings.Split(encodedHash, ".")
	if len(parts) != 2 {
		return ErrorHandler(
			errors.New("Invalid encoded hash format"),
			"Invalid encoded hash format",
		)
	}
	saltBase64 := parts[0]
	passwordHashBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return ErrorHandler(err, "Failed to decode salt")
	}
	passwordHash, err := base64.StdEncoding.DecodeString(passwordHashBase64)
	if err != nil {
		return ErrorHandler(err, "Failed to decode password")
	}

	inputPasswordHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if len(inputPasswordHash) != len(passwordHash) {
		return ErrorHandler(errors.New("Incorect password"), "Incorect password")
	}

	if subtle.ConstantTimeCompare(inputPasswordHash, passwordHash) == 1 {
	} else {
		return ErrorHandler(errors.New("Incorect password"), "Incorect password")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", ErrorHandler(errors.New("Password is blank"), "Please enter password")
	}
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", ErrorHandler(
			errors.New("Failed to generate salt"),
			"Internal error",
		)
	}
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)
	return encodedHash, nil
}
