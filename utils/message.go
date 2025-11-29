package utils

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("task not found")
var ErrFile = errors.New("file error")
func ErrorMessage(err error) string {
	return fmt.Sprintf("\033[31m[FAILED] Something went wrong: %v\033[0m", err.Error())
}