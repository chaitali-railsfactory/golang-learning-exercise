/*
Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5,
between 2000 and 3200 (both included). The numbers obtained should be printed in a comma-separated sequence on a single line.

*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var nums []string

	for num := 2000; num <= 3200; num++ {
		if (num%7 == 0) && (num%5 != 0) {
			nums = append(nums, strconv.Itoa(num))
		}
	}
	fmt.Println(strings.Join(nums, " , "))
}
