package main

import (
	"errors"
	"fmt"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("Out of range")
	}
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := ""

	for _, numeral := range romanNumerals {
		for input >= numeral.Value {
			roman += numeral.Symbol
			input -= numeral.Value
		}
	}
	return roman, nil
}

func main() {
	num := 1223
	roman, err := ToRomanNumeral(num)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Roman number for %d is %s \n", num, roman)
}
