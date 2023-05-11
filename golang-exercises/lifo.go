// Implement a stack structure with pop, append, and print top functionalities.

package main

import "fmt"

func lifo(slice []int) []int {
	lenSlice := len(slice)
	var lifoSlice []int
	for i := (lenSlice - 1); i >= 0; i-- {
		fmt.Println(slice[i])
		lifoSlice = append(lifoSlice, slice[i])
	}
	return lifoSlice
}

func main() {
	array := []int{1, 2, 3, 6, 7, 8}
	fmt.Println(lifo(array))
}
