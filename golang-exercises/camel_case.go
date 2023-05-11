// There is a sequence of words in CamelCase as a string of letters, , having the following properties:

// It is a concatenation of one or more words consisting of English letters.
// All letters in the first word are lowercase.
// For each of the subsequent words, the first letter is uppercase and rest of the letters are lowercase.
// Given , determine the number of words in .

// Example

// There are  words in the string: 'one', 'Two', 'Three'.

// Function Description

// Complete the camelcase function in the editor below.

// camelcase has the following parameter(s):

// string s: the string to analyze
// Returns

// Hint: the number of words in S

package main

import (
	"fmt"
	"unicode"
)

func camelCaseCount(functionName string) int {
	count := 0
	for index, char := range functionName {
		if index == 0 || unicode.IsUpper(char) {
			count++
		}
	}
	return count
}
func main() {
	functionName := "saveChangesInTheEditor"
	fmt.Printf("Camel case count is: %d \n", camelCaseCount(functionName))
}
