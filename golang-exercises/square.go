// With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

// Suppose the following input is supplied to the program: 8

// Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

package main

import "fmt"

func squareMap(num int) map[int]int {
	square := make(map[int]int)
	for index := 1; index <= num; index++ {
		square[index] = index * index
	}
	return square
}

func main() {
	fmt.Println(squareMap(8))
}
