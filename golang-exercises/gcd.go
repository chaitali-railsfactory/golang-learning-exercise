package main

import "fmt"

// Euclidean algorithms
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main() {
	fmt.Println(GCD(30, 9))
}
