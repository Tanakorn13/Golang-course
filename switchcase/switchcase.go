package main

import "fmt"

var input string
var colorCode string

func main() {
	fmt.Println("Enter the color: ")
	fmt.Scanf("%s", &input)
	switch input {
	case "blue":
		fmt.Println("#0000ff")
	case "green":
		fmt.Println("#00ff00")
	case "pink":
		fmt.Println("#ff00ff")
	case "yellow":
		fmt.Println("#ffff00")
	default:
		fmt.Println("No color")
	}
}
