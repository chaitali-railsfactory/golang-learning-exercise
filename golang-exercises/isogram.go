// Instructions
// Determine if a word or phrase is an isogram.

// An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

// Examples of isograms:

// lumberjacks
// background
// downstream
// six-year-old
// The word isograms, however, is not an isogram, because the s repeats.

package main

import (
	"fmt"
	"strings"
)

func wordCount(words string) map[string]int {
	count := make(map[string]int)
	for i := 0; i < len(words); i++ {
		char := strings.ToUpper(string(words[i]))
		_, exist := count[char]
		if exist {
			count[char] += 1
		} else {
			count[char] = 1
		}
	}
	return count
}
func IsIsogram(word string) bool {
	wordCount := wordCount(word)
	isIso := true
	for char, count := range wordCount {
		if count > 1 && char != "-" && char != " " {
			isIso = false
		}
	}
	return isIso

}
func main() {
	fmt.Printf("'%s' is Isogram %v \n", "lumberjacks", IsIsogram("lumberjacks"))
	fmt.Printf("'%s' is not Isogram %v \n", "zzyzx", IsIsogram("zzyzx"))
}
