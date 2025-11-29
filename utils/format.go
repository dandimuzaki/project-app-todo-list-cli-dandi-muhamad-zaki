package utils

import (
	"errors"
	"time"
)

func StringToDatetime(s string) (time.Time, error) {
	resDate, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Now(), errors.New("invalid date format. use YYYY-MM-DD")
	}
	return resDate, nil
}