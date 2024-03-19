package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name  string
	price map[string]float64
}

var menue = []menuItem{
	{name: "Coffee", price: map[string]float64{"small": 1.65, "medium": 1.8, "large": 1.95}},
	{name: "Tea", price: map[string]float64{"single": 1.8, "double": 2.1, "triple": 2.95}},
}

func main() {

	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)

	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice) // we don't know what to do with this yet

	for _, item := range menue {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.price {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
