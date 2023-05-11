package main

import "fmt"

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//

func Fibonacci(n int) int {
	num1, num2 := 0, 1
	var sum int
	for i := 2; i <= n; i++ {
		sum = num1 + num2
		num1, num2 = num2, sum
	}
	return sum
}

func main() {
	fmt.Println(Fibonacci(6))
}
