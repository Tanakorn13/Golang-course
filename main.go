package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var msg string = "Hello"
func main() {
	numberfloat := 10.5
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberInt) + numberfloat)
	fmt.Println(msg + " Go")
	fmt.Println("My money =", numberInt + numberInt2)
}