package validator

import "regexp"

var accountNumberRegex = regexp.MustCompile(`^\d{10}$`)

func ValidateAccountNumber(accountNumber string) bool {
	return accountNumberRegex.MatchString(accountNumber)
}

func ValidateAmount(amount float64) bool {
	return amount > 0
}
