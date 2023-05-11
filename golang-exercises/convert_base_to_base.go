// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//
package main

import "fmt"

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

func DecToBase(dec, base int) string {
	var res string
	var rem int
	const charset = "0123456789ABCDEF"
	for dec > 0 {
		rem = dec % base
		dec = dec / base
		res = fmt.Sprintf("%s%s", string(charset[rem]), res)
	}
	return res
}

func BaseToBase(value string, base, newBase int) string {
	decimal := BaseToDec(value, base)
	newValue := DecToBase(decimal, newBase)
	return newValue
}

func main() {
	fmt.Println(BaseToBase("3735928559", 10, 4))
}
