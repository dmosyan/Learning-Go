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

var menu = []menuItem{
	{name: "Coffee", price: map[string]float64{"small": 1.65, "medium": 1.8, "large": 1.95}},
	{name: "Tea", price: map[string]float64{"single": 1.8, "double": 2.1, "triple": 2.95}},
}

func main() {

loop:
	for {
		in := bufio.NewReader(os.Stdin)

		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.price {
					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the name of the new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, price: make(map[string]float64)})
		case "q":
			break loop
		}
	}
}
