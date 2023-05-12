package main

import (
	"fmt"
	"sort"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	len := len(list)
	for i := 0; i < len; i++ {
		sweep := false
		for j := 0; j < len-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				sweep = true
			}
		}
		if !sweep {
			break
		}
	}
	fmt.Println("Sorted list", list)
}

// BubbleSortString is a bubble sort for string slices.
// Alphabetically sort
func BubbleSortString(list []string) {
	len := len(list)
	for i := 0; i < len; i++ {
		sweep := false
		for j := 0; j < len-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				sweep = true
			}
		}
		if !sweep {
			break
		}
	}
	fmt.Println("Sorted list", list)
}

type Person struct {
	Age       int
	LastName  string
	FirstName string
}

type People []Person

func (p People) Len() int { return len(p) }
func (p People) Less(i, j int) bool {
	a, b := p[i], p[j]
	if a.Age != b.Age {
		return a.Age < b.Age
	}
	if a.LastName != b.LastName {
		return a.LastName < b.LastName
	}
	return a.FirstName < b.FirstName
}
func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people People) {
	sort.Sort(People(people))
	for _, person := range people {
		fmt.Printf("First Name: %s, Last Name: %s and Age %d\n", person.FirstName, person.LastName, person.Age)
	}
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		sweep := false
		for j := 0; j < list.Len()-1-i; j++ {
			if list.Less(j+1, j) {
				list.Swap(j, j+1)
				sweep = true
			}
		}
		if !sweep {
			break
		}
	}
	fmt.Println("Sorted list", list)
}

func main() {
	list := []int{2, 5, 6, 3, 1}
	listOfString := []string{"hi", "hello", "how", "are", "you"}
	people := []Person{
		{71, "Thomas", "Train"},
		{65, "Harry", "Hippo"},
		{53, "Percy", "Engine"},
		{45, "Johnny", "Testuser"},
		{33, "Bob", "Smilesalot"},
		{21, "Frank", "Smith"},
		{12, "Alex", "Zero"},
		{12, "Jordan", "Tables"},
		{12, "Bobby", "Tables"},
		{12, "Billy", "Tables"},
	}
	fmt.Println("---Sorting of Integer---")
	BubbleSortInt(list)
	fmt.Println("---Sorting of String---")
	BubbleSortString(listOfString)
	fmt.Println("Sorting of Struct")
	BubbleSortPerson(people)
	fmt.Println("---Sorting by sort Interface with int slice---")
	BubbleSort(sort.IntSlice(list))
	fmt.Println("---Sorting by sort Interface with string slice---")
	BubbleSort(sort.StringSlice(listOfString))
	fmt.Println("---Sorting by sort Interface with custom slice---")
	BubbleSort(People(people))
}
