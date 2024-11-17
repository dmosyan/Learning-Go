package main

import (
	"fmt"
	"slices"
)

func main() {

	people := map[string]int{
		"Alice":   22,
		"Bob":     18,
		"Charlie": 23,
		"Dave":    27,
		"Eve":     31,
	}

	fmt.Println(people)

	fmt.Println(people["Nemo"]) // 0 as the key doesn't exist

	age, ok := people["Nemo"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Printf("key doesn't exist, value: %v\n", age)
	}

	// leave out value
	for k := range people {
		fmt.Println(k)
	}

	// remove specific entry with the key
	delete(people, "Alice")
	fmt.Println(people)

	// sorting in descending order based on key (map itself doesn't have any order)
	var keys []string
	for k := range people {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a, b string) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	})
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, people[key])
	}
}
