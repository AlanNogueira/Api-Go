package utils

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const expiryDateLayout = "02-01-2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(expiryDateLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(expiryDateLayout))), nil
}

func (ct *CustomTime) Validate() string {
	if ct.Time.IsZero() {
		return "ReleaseDate, "
	}

	return ""
}

func (ct *CustomTime) Format() (err error) {
	stringDate := ct.Time.Format(layout)
	ct.Time, err = time.Parse(layout, stringDate)
	if err != nil {
		return
	}

	return nil
}
