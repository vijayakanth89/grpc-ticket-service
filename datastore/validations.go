package datastore

import (
	"errors"
	"fmt"
	"regexp"
)

func IsZero(userId string) error {
	if userId == "0" {
		return errors.New(ERROR_INVALID_USER_ID)
	}
	return nil
}

func CheckEmpty(args ...string) error {
	for _, arg := range args {
		if arg == "" {
			return errors.New(arg + ERROR_EMPTY_STRING)
		}
	}
	return nil
}

func IsValidEmail(email string) error {
	// Regular expression for basic email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return errors.New(fmt.Sprintf(ERROR_INVALID_EMAIL_TEMPLATE, email))
	} else {
		return nil
	}
}
