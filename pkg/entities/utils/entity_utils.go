package utils

import (
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"go.mongodb.org/mongo-driver/mongo/options"
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
	characters := 0
	var eighth, number, upper, special bool
	for _, character := range password {
		switch {
		case unicode.IsNumber(character):
			number = true
			characters++
		case unicode.IsUpper(character):
			upper = true
			characters++
		case unicode.IsPunct(character) || unicode.IsSymbol(character):
			special = true
			characters++
		case unicode.IsLetter(character) || character == ' ':
			characters++
		}
	}
	eighth = characters >= 7
	return eighth && number && upper && special
}

func ValidateDeliveryDate(date CustomTime) bool {
	dateNow := time.Now().Format(layout)
	minDeliveryDate, err := time.Parse(layout, dateNow)
	if err != nil {
		return false
	}
	maxDeliveryDate := minDeliveryDate.AddDate(0, 0, 30)

	return !date.Time.After(maxDeliveryDate) && !date.Time.Before(minDeliveryDate)
}

func ReturnMessageOrValue(value interface{}, valueType string) interface{} {
	if value == nil {
		if valueType == "out" {
			return "no books out of the limit date"
		} else if valueType == "within" {
			return "no books within the limit date"
		}
	}

	return value
}

func Pagination(r *http.Request, FindOptions *options.FindOptions) (int64, int64) {
	if r.URL.Query().Get("page") != "" && r.URL.Query().Get("limit") != "" {
		page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
		limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
		if page == 1 || page == 0 {
			FindOptions.SetSkip(0)
			FindOptions.SetLimit(limit)
			return page, limit
		}

		FindOptions.SetSkip((page - 1) * limit)
		FindOptions.SetLimit(limit)
		return page, limit

	}
	FindOptions.SetSkip(0)
	FindOptions.SetLimit(0)
	return 0, 0
}
