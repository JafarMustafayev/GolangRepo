package Validations

import (
	"awesomeProject/Enums"
	"unicode"
)

func ClassNameValidation(className string) bool {

	if len(className) < 5 {
		return false
	}
	var digits = []string{}
	var strings = []string{}
	var upperCase = []string{}

	var j int
	for j = 0; j < len(className); j++ {
		var char = className[j]

		if unicode.IsDigit(rune(char)) {
			digits = append(digits, string(char))
		} else if unicode.IsUpper(rune(char)) {
			upperCase = append(upperCase, string(char))
		} else {
			strings = append(strings, string(char))
		}

	}

	return true
}

func ClassTypeValidation(classType int) (bool, int) {

	if classType == Enums.BackEnd ||
		classType == Enums.FrontEnd ||
		classType == Enums.FullStack {

		var studentCount = 0
		switch classType {
		case Enums.BackEnd:
			println("Back end")
			studentCount = 20
		case Enums.FrontEnd:
			println("Front end")
			studentCount = 10
		default:
			println("Full stack")
			studentCount = 15
		}

		return true, studentCount
	}
	return false, 0
}
