// Write a program that accepts a sentence and calculate the number of letters and digits.

package main

import (
	"fmt"
	"unicode"
)

func calculateDigitAndLetter(sentence string) {
	digits, letters := 0, 0
	for _, char := range sentence {
		if unicode.IsDigit(char) {
			digits++
		} else if unicode.IsLetter(char) {
			letters++
		} else {
			continue
		}
	}
	fmt.Printf("Letter: %d and digits: %d \n", letters, digits)
}
func main() {
	sentence := "hello world! 123"
	calculateDigitAndLetter(sentence)
}
