package service

import "regexp"

func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return regex.MatchString(email)
}

func IsValidPassword(password string) bool {
	return len(password) >= 6
}

func ValidateAuthRegister(email, password string) string {
	if !IsValidEmail(email) {
		return "Email tidak valid"
	}
	if !IsValidPassword(password) {
		return "Password minimal 6 karakter"
	}
	return ""
}
