package utils

import (
	"errors"
)

type ContextString string

func AuthorizeUser(userRole string, allowedRoles ...string) (bool, error) {
	for _, allowedRole := range allowedRoles {
		if userRole == allowedRole {
			return true, nil
		}
	}

	return false, errors.New("User not authorized")
}
