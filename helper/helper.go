package helper

import "strings"

func ValidateUserInput(fname string, lname string, email string, userTix uint, remaininTix uint) (bool, bool, bool) {
	isValidName := len(fname) >= 2 && len(lname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTix := userTix <= remaininTix && userTix > 0

	return isValidEmail, isValidUserTix, isValidName
}
