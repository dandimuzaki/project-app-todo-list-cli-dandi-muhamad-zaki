package utils

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("task not found")
var ErrFile = errors.New("file error")
var ErrInvalidStatus = errors.New("invalid status. please enter 'on progress', 'finished', or 'on hold'")
var ErrInvalidPriority = errors.New("invalid priority. please enter 'low', 'normal', 'urgent', or 'critical'")

func ErrorMessage(err error) string {
	return fmt.Sprintf("\033[31m[FAILED] Something went wrong: %v\033[0m", err.Error())
}