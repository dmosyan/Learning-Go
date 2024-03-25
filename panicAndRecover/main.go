package main

import "fmt"

func main() {
	dividend, divisor := 10, 5
	fmt.Printf("%v devided by %v is %v\n", dividend, divisor, devide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v devided by %v is %v\n", dividend, divisor, devide(dividend, divisor))

}

// create defer function and recover function to handle panic situation (like deviding by 0)

func devide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}
