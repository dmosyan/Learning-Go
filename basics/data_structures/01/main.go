package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//playWithSlices()
	concSafeMod()

}

func playWithSlices() {
	// declare array
	array1 := [4]string{}
	fmt.Println(array1, cap(array1))

	// declare a slice
	slice1 := []int{}
	fmt.Println(len(slice1), cap(slice1))

	// declare slice with make with lenght 2 and capacity 5
	slice2 := make([]string, 2, 5)
	fmt.Println(slice2, len(slice2), cap(slice2))

	// append a new element in slice
	slice2 = append(slice2, "new")
	fmt.Println(slice2, len(slice2)) // [ new] 3

	//appen a new slice in slice
	slice3 := []string{"a", "new", "slice"}
	slice2 = append(slice2, slice3...)
	fmt.Println(slice2, len(slice2)) // [  new a new slice] 6

	// change a value in slice
	slice2[0] = "first"
	fmt.Println(slice2, len(slice2), cap(slice2)) // [first  new a new slice] 6 10

	// iterate over slice
	slice2[1] = "second"
	for i, v := range slice2 {
		fmt.Printf("index: %v value: %v\n", i, v)
	}

	// add a new value in a specific location in slice
	slice2 = append(slice2[:1+1], slice2[1:]...)
	fmt.Println(slice2) // [first second second new a new slice]
	slice2[2] = "third"
	fmt.Println(slice2) // [first second third new a new slice]

	// add a new element in the begning of a slice
	slice2 = append([]string{"Very First"}, slice2...)
	fmt.Println(slice2)

	// add a new slice in between another slice
	tail := append([]string{"something", "middle"}, slice2[2:]...)
	slice2 = append(slice2[:1], tail...)
	fmt.Println(slice2)

	// remove an element from a slice
	numbers := []int{3, 14, 159, 26, 53, 58}
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)

}

var (
	shared []int = []int{1, 2, 3, 4, 5, 6, 7}
	mutex  sync.Mutex
	wg     sync.WaitGroup
)

// concurrency safe modifications of slices
func concSafeMod() {

	wg.Add(2)
	go increaseValue(shared, &wg)
	go decreaseValue(shared, &wg)
	wg.Wait()

	fmt.Println("modification is done")

}

func increaseValue(slice []int, wg *sync.WaitGroup) {

	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < len(slice); i++ {
		time.Sleep(20 * time.Microsecond)
		slice[i]++
	}

	fmt.Println(slice)
}

func decreaseValue(slice []int, wg *sync.WaitGroup) {

	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < len(slice); i++ {
		time.Sleep(20 * time.Microsecond)
		slice[i]--
	}

	fmt.Println(slice)
}
