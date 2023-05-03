/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8

Then, the output should be: 40320
*/
package main

import "fmt"

func factorial(num int) int {
	var fact int = 1
	for value := num; value >= 1; value-- {
		fact = fact * value
	}
	return fact
}

func main() {
	num := 8
	fmt.Printf("Factorial of %d: %d", num, factorial(num))
}
