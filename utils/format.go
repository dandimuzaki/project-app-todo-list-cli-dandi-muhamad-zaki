package utils

import (
	"strings"
)

func Uppercase(s string) string {
	words := strings.Split(s, " ")
	var capitalized []string
	for _, w := range words {
		newWord := strings.ToUpper(string(w[0])) + w[1:]
		capitalized = append(capitalized, newWord)
	}
	return strings.Join(capitalized, " ")
}

func Capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func IsValidStatus(s string) bool {
	valid := false
	if strings.ToLower(s) == "on progress" {
		valid = true
	}
	if strings.ToLower(s) == "on hold" {
		valid = true
	}
	if strings.ToLower(s) == "finished" {
		valid = true
	}
	return valid
}

func IsValidPriority(s string) bool {
	valid := false
	if strings.ToLower(s) == "low" {
		valid = true
	}
	if strings.ToLower(s) == "normal" {
		valid = true
	}
	if strings.ToLower(s) == "urgent" {
		valid = true
	}
	if strings.ToLower(s) == "critical" {
		valid = true
	}
	return valid
}