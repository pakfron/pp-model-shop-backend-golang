package validate

import "regexp"

func ValidateRegister(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	println(emailRegex.MatchString(email))
	return emailRegex.MatchString(email)
}
