package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number", promt)
		panic(message)
	}

	return value
}

func getOperator() string {
	fmt.Println("operator is (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}
func subs(value1, value2 float64) float64 {
	return value1 - value2
}
func mul(value1, value2 float64) float64 {
	return value1 * value2
}
func devide(value1, value2 float64) float64 {
	return value1 / value2
}
func main() {
	value1 := getInput("Enter the value 1: ")
	value2 := getInput("Enter the value 2: ")

	var result float64
	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subs(value1, value2)
	case "*":
		result = mul(value1, value2)
	case "/":
		result = devide(value1, value2)
	default:
		panic("invalid operator")
	}
	fmt.Printf("Result: %v\n", result)
}
