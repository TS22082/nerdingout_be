package utils

import "fmt"

var whiteList = []string{"ts22082@gmail.com", "nurdragedevelopment@gmail.com"}

// ValidateEmail returns an error if the email does not match an email saved in the whitelist.
func ValidateEmail(email string) error {
	// @param email the email to validate
	// @return null if email is valid, or an error message otherwise
	if email == "" {
		return fmt.Errorf("invalid input: email cannot be blank")
	}

	for _, value := range whiteList {
		if value == email {
			return nil
		}
	}

	return fmt.Errorf("email %s is not valid", email)
}
