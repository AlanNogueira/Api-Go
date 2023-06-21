package utils

import (
	"net/mail"
	"regexp"
	"time"
	"unicode"
)

const layout = "02-01-2006"

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateName(name string) bool {
	match, _ := regexp.MatchString(`^((\b[A-zÀ-ú']{2,40}\b)\s*){2,}$`, name)
	return match
}

func ValidatePassword(password string) bool {
	letters := 0
	var sevenOrMore, number, upper, special bool
	for _, character := range password {
		switch {
		case unicode.IsNumber(character):
			number = true
		case unicode.IsUpper(character):
			upper = true
			letters++
		case unicode.IsPunct(character) || unicode.IsSymbol(character):
			special = true
		case unicode.IsLetter(character) || character == ' ':
			letters++
		}
	}
	sevenOrMore = letters >= 7
	return sevenOrMore && number && upper && special
}

func ValidateDeliveryDate(date CustomTime) bool {
	dateNow := time.Now().Format(layout)
	maxDeliveryDate, err := time.Parse(layout, dateNow)
	if err != nil {
		return false
	}
	maxDeliveryDate = maxDeliveryDate.AddDate(0, 0, 30)

	return !date.Time.IsZero() && !date.Time.After(maxDeliveryDate)
}
