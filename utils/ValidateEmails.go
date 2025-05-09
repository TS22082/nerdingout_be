package utils

import "fmt"

var whiteList = []string{"ts22082@gmail.com", "nurdragedevelopment@gmail.com"}

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("Invalid input: email cannot be blank")
	}

	for _, value := range whiteList {
		if value == email {
			return nil
		}
	}

	return fmt.Errorf("Email %s is not valid", email)
}
