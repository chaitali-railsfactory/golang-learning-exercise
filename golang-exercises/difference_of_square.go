// Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.

// The square of the sum of the first ten natural numbers is (1 + 2 + ... + 10)² = 55² = 3025.

// The sum of the squares of the first ten natural numbers is 1² + 2² + ... + 10² = 385.

// Hence the difference between the square of the sum of the first ten natural numbers and the sum of the squares of the first ten natural numbers is 3025 - 385 = 2640.

package main

import "fmt"

func square(num int) int {
	return num * num
}

func sum(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

func squareOfSum(min int, max int) int {
	var slice []int
	for i := min; i <= max; i++ {
		slice = append(slice, i)
	}
	sumofSlice := sum(slice)
	return square(sumofSlice)
}

func sumOfSquare(min int, max int) int {
	var slice []int
	for i := min; i <= max; i++ {
		slice = append(slice, square(i))
	}
	return sum(slice)
}
func main() {
	min, max := 1, 10

	difference := squareOfSum(min, max) - sumOfSquare(min, max)
	fmt.Println("Difference is: ", difference)
}
