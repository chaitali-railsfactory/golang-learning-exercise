// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"

package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(DecToBase(3735928559, 11))
}
