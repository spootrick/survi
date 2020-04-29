package util

import "regexp"

var validEmailRegex = `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}\b`

func VerifyEmailFormat(email string) bool {
	isValid, _ := regexp.MatchString(validEmailRegex, email)

	return isValid
}
