package service

import "regexp"

func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return regex.MatchString(email)
}
