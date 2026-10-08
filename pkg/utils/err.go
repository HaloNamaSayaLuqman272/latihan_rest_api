package utils

import "strings"

func GetValidationErrorMessage(message string) []string {
	separator := ","
	errors := strings.Split(message, separator)

	return errors[:len(errors)-1]
}
