package controllers

import (
	"fmt"
	"strings"
)

func validatePassword(password string) error {
	num := "1234567890"
	simbol := "!@#$%^&*"

	if len(password) <= 6 {
		return fmt.Errorf("password harus lebih dari 6 karakter")
	}

	hasSymbol := false
	hasNumber := false

	for _, char := range password {
		if strings.ContainsRune(simbol, char) {
			hasSymbol = true
		}
		if strings.ContainsRune(num, char) {
			hasNumber = true
		}
	}

	if !hasSymbol {
		return fmt.Errorf("password harus mengandung setidaknya satu simbol")
	}
	if !hasNumber {
		return fmt.Errorf("password harus mengandung setidaknya satu angka")
	}

	return nil
}
