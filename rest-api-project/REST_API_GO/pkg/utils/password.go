package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
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
