// Create an implementation of the rotational cipher, also sometimes called the Caesar cipher.

// The Caesar cipher is a simple shift cipher that relies on transposing all the letters in the alphabet using an integer key between 0 and 26. Using a key of 0 or 26 will always yield the same output due to modular arithmetic. The letter is shifted for as many values as the value of the key.

// The general notation for rotational ciphers is ROT + <key>. The most commonly used rotational cipher is ROT13.

// A ROT13 on the Latin alphabet would be as follows:

// Plain:  abcdefghijklmnopqrstuvwxyz
// Cipher: nopqrstuvwxyzabcdefghijklm
// It is stronger than the Atbash cipher because it has 27 possible keys, and 25 usable keys.

// Ciphertext is written out in the same formatting as the input including spaces and punctuation.

// Examples
// ROT5 omg gives trl
// ROT0 c gives c
// ROT26 Cool gives Cool
// ROT13 The quick brown fox jumps over the lazy dog. gives Gur dhvpx oebja sbk whzcf bire gur ynml qbt.
// ROT13 Gur dhvpx oebja sbk whzcf bire gur ynml qbt. gives The quick brown fox jumps over the lazy dog.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var encrypted []string
	for _, char := range plain {
		if !unicode.IsLetter(char) {
			encrypted = append(encrypted, string(rune(char)))
			continue
		} else {
			encryptedChar := rune(char + int32(shiftKey)%26)
			if unicode.IsLower(char) && encryptedChar > 'z' {
				encryptedChar = encryptedChar - 26
			} else if unicode.IsUpper(char) && encryptedChar > 'Z' {
				encryptedChar = encryptedChar - 26
			}

			encrypted = append(encrypted, string(rune(encryptedChar)))
		}
	}
	return strings.Join(encrypted, "")
}

var testCases = []struct {
	description   string
	inputPlain    string
	inputShiftKey int
	expected      string
}{
	{
		description:   "rotate a by 0, same output as input",
		inputPlain:    "a",
		inputShiftKey: 0,
		expected:      "a",
	},
	{
		description:   "rotate a by 1",
		inputPlain:    "a",
		inputShiftKey: 1,
		expected:      "b",
	},
	{
		description:   "rotate a by 26, same output as input",
		inputPlain:    "a",
		inputShiftKey: 26,
		expected:      "a",
	},
	{
		description:   "rotate n by 13 with wrap around alphabet",
		inputPlain:    "n",
		inputShiftKey: 13,
		expected:      "a",
	},
	{
		description:   "rotate capital letters",
		inputPlain:    "OMG",
		inputShiftKey: 5,
		expected:      "TRL",
	},
	{
		description:   "rotate spaces",
		inputPlain:    "O M G",
		inputShiftKey: 5,
		expected:      "T R L",
	},
	{
		description:   "rotate numbers",
		inputPlain:    "Testing 1 2 3 testing",
		inputShiftKey: 4,
		expected:      "Xiwxmrk 1 2 3 xiwxmrk",
	},
	{
		description:   "rotate punctuation",
		inputPlain:    "Let's eat, Grandma!",
		inputShiftKey: 21,
		expected:      "Gzo'n zvo, Bmviyhv!",
	},
	{
		description:   "rotate all letters",
		inputPlain:    "The quick brown fox jumps over the lazy dog.",
		inputShiftKey: 13,
		expected:      "Gur dhvpx oebja sbk whzcf bire gur ynml qbt.",
	},
}

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	for _, testCase := range testCases {
		fmt.Printf("----------- \n")
		fmt.Printf("Description: %s \n", testCase.description)

		got := RotationalCipher(testCase.inputPlain, testCase.inputShiftKey)

		if testCase.expected != got {
			fmt.Printf("%v Test Case Fail - Expected: %v and Got: %v \n", string(colorRed), testCase.expected, got)
		} else {
			fmt.Printf("%v Test Case Pass - Expected: %v and Got: %v \n", string(colorGreen), testCase.expected, got)
		}

	}
}
