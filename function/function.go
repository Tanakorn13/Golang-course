package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func plus(value1 int, value2 int) int{
	var sum int = value1 + value2
	return sum
}

func plus3value(value1 , value2 , value3 int) int{
	return value1 + value2 + value3
}
func main() {
	hello()
	result := plus(5, 3)
	fmt.Println("5 + 3 =", result)
	result2 := plus3value(1, 2, 3)
	fmt.Println("1 + 2 + 3=", result2)
}
