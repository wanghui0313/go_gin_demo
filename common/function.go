package common

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func Sum(a, b int) int {
	return a + b
}

func CheckMobile(f validator.FieldLevel) bool {
	field := f.Field().Interface().(string)
	exp := "^1[3|5|7|8|9][0-9]\\d{8}$"
	if match, _ := regexp.MatchString(exp, field); match {
		return true
	}
	return false
}
