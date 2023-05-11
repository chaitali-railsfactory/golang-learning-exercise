// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//

package main

import (
	"fmt"
)

func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	res := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for innerIndex, char := range charset {
			if rune(value[i]) == char {
				index = innerIndex
				break
			}
		}
		res += index * multiplier
		multiplier = multiplier * base
	}
	return res
}

func main() {
	fmt.Println(BaseToDec("8831A383B", 12))
}
