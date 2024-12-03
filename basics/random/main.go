package main

import (
	"fmt"
	"strings"
)

func main() {
	printChar()
}

func printChar() {
	start, stop := 'A', 'Z'

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "dec", "hex", "encoded", strings.Repeat("-", 45))

	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}
