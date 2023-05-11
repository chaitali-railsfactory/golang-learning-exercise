// Implement the SumOfSquares function which takes an integer, c and returns the sum of all squares between 1 and c. You’ll need to use select statements, goroutines, and channels.

// For example, entering 5 would return 55 because

package main

import "fmt"

func SumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}

func main() {
	sumOfSquaresChannel := make(chan int)
	quitChannel := make(chan int)

	sum := 0

	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-sumOfSquaresChannel
		}
		fmt.Println(sum)
		quitChannel <- 0
	}()
	SumOfSquares(sumOfSquaresChannel, quitChannel)
}
