// Given a string the program should check if the provided string is a valid ISBN-10. Putting this into place requires some thinking about preprocessing/parsing of the string prior to calculating the check digit for the ISBN.

// The program should be able to verify ISBN-10 both with and without separating dashes.

// Caveats
// Converting from strings to numbers can be tricky in certain languages. Now, it's even trickier since the check digit of an ISBN-10 may be 'X' (representing '10'). For instance 3-598-21507-X is a valid ISBN-10.

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	matches := strings.NewReplacer("-", "").Replace(isbn)
	matches = strings.ToLower(matches)
	if len(matches) != 10 {
		return false
	}

	sum := 0
	invalid := false
	for index, value := range matches {
		if (unicode.IsLetter(value) && value != 'x') || index > 9 {
			invalid = true
			break
		}

		num := int(value) - '0'
		if value == 'x' && index == 9 {
			num = 10
		}
		sum += ((10 - index) * num)
	}
	if invalid {
		return false
	}
	return sum%11 == 0
}

var testCases = []struct {
	description string
	isbn        string
	expected    bool
}{
	{
		description: "valid isbn",
		isbn:        "3-598-21508-8",
		expected:    true,
	},
	{
		description: "invalid isbn check digit",
		isbn:        "3-598-21508-9",
		expected:    false,
	},
	{
		description: "valid isbn with a check digit of 10",
		isbn:        "3-598-21507-X",
		expected:    true,
	},
	{
		description: "check digit is a character other than X",
		isbn:        "3-598-21507-A",
		expected:    false,
	},
	{
		description: "invalid check digit in isbn is not treated as zero",
		isbn:        "4-598-21507-B",
		expected:    false,
	},
	{
		description: "invalid character in isbn is not treated as zero",
		isbn:        "3-598-P1581-X",
		expected:    false,
	},
	{
		description: "X is only valid as a check digit",
		isbn:        "3-598-2X507-9",
		expected:    false,
	},
	{
		description: "valid isbn without separating dashes",
		isbn:        "3598215088",
		expected:    true,
	},
	{
		description: "isbn without separating dashes and X as check digit",
		isbn:        "359821507X",
		expected:    true,
	},
	{
		description: "isbn without check digit and dashes",
		isbn:        "359821507",
		expected:    false,
	},
	{
		description: "too long isbn and no dashes",
		isbn:        "3598215078X",
		expected:    false,
	},
	{
		description: "too short isbn",
		isbn:        "00",
		expected:    false,
	},
	{
		description: "isbn without check digit",
		isbn:        "3-598-21507",
		expected:    false,
	},
	{
		description: "check digit of X should not be used for 0",
		isbn:        "3-598-21515-X",
		expected:    false,
	},
	{
		description: "empty isbn",
		isbn:        "",
		expected:    false,
	},
	{
		description: "input is 9 characters",
		isbn:        "134456729",
		expected:    false,
	},
	{
		description: "invalid characters are not ignored after checking length",
		isbn:        "3132P34035",
		expected:    false,
	},
	{
		description: "invalid characters are not ignored before checking length",
		isbn:        "3598P215088",
		expected:    false,
	},
	{
		description: "input is too long but contains a valid isbn",
		isbn:        "98245726788",
		expected:    false,
	},
}

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	for index, _ := range testCases {
		fmt.Printf("----------- \n")
		fmt.Printf("Description: %s \n", testCases[index].description)
		got := IsValidISBN(testCases[index].isbn)

		if testCases[index].expected != got {
			fmt.Printf("%v Test Case Fail - Expected: %v and Got: %v \n", string(colorRed), testCases[index].expected, got)
		} else {
			fmt.Printf("%v Test Case Pass - Expected: %v and Got: %v \n", string(colorGreen), testCases[index].expected, got)
		}

	}
}
