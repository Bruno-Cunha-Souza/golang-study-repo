package luhn

import "errors"

var ErrInvalidInput = errors.New("invalid input: must contain only digits and be at least 2 characters")

func LuhnAlgorithm(cardNumber string) (bool, error) {
	if len(cardNumber) < 2 {
		return false, ErrInvalidInput
	}

	total := 0
	isSecondDigit := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		char := cardNumber[i]
		if char < '0' || char > '9' {
			return false, ErrInvalidInput
		}
		digit := int(char - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit
		isSecondDigit = !isSecondDigit
	}

	return total%10 == 0, nil
}
