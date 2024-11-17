package main

import (
	"fmt"
	"slices"
)

var (
	integers = []int{3, 14, 156, 24, 46}
	floats   = []float64{3.14, 1.41, 1.73, 2.72, 4.53}
	strings  = []string{"the", "quick", "brown", "fox", "jumped"}
)

func main() {
	sortingAscending()
	sortingDescending()
	stableSortDescending()

}

func sortingAscending() {
	slices.Sort(integers)
	fmt.Println(integers)

	slices.Sort(floats)
	fmt.Println(floats)

	slices.Sort(strings)
	fmt.Println(strings)
}

func sortingDescending() {

	for i := len(integers)/2 - 1; i >= 0; i-- {
		opp := len(integers) - 1 - i
		integers[i], integers[opp] = integers[opp], integers[i]
	}

	fmt.Println(integers)

	slices.SortFunc(floats, func(a, b float64) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	})

	fmt.Println(floats)

	slices.SortFunc(strings, func(a, b string) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	})

	fmt.Println(strings)
}

type Person struct {
	Name string
	Age  int
}

func stableSortDescending() {

	people := []Person{
		{"Alice", 22},
		{"Bob", 18},
		{"Charlie", 23},
		{"Dave", 27},
		{"Eve", 31},
	}

	slices.SortFunc(people, func(a, b Person) int {
		if a.Age > b.Age {
			return -1
		}
		if a.Age < b.Age {
			return 1
		}
		return 0
	})

	fmt.Println(people)
}
