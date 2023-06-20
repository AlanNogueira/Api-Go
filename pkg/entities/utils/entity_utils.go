package utils

import (
	"net/mail"
	"regexp"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateName(name string) bool {
	match, _ := regexp.MatchString(`^((\b[A-zÀ-ú']{2,40}\b)\s*){2,}$`, name)
	return match
}
