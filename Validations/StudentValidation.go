package Validations

import (
	"unicode"
)

func StudentNameValidation(studentName string) bool {

	var nameLength = len(studentName)
	if nameLength < 3 {
		return false
	}

	if !unicode.IsUpper(rune(studentName[0])) {
		return false
	}
	return true
}

func StudentAgeValidation(age int) bool {
	if age < 18 || 40 < age {
		return false
	}
	return true
}

func StudentValidation(name string, surname string, age int) bool {

	if StudentAgeValidation(age) &&
		StudentNameValidation(name) &&
		StudentNameValidation(surname) {
		return true
	}
	return false
}
