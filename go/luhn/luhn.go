package luhn

import (
	"strconv"
	"strings"
)

// Valid determines whether credit card number is valid
func Valid(creditCard string) bool {
	strippedCreditCard := strings.Replace(creditCard, " ", "", -1)
	cardLen := len(strippedCreditCard)
	if cardLen <= 1 {
		return false
	}
	checksum := 0
	for i, c := range strippedCreditCard {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
		if (cardLen-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		checksum += digit
	}
	return checksum%10 == 0
}
