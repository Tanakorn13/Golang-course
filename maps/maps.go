package main

import (
	"fmt"
)

var product = make(map[string]float64)

func main() {
	fmt.Println("product = ", product)

	//add
	product["Macbook"] = 40000
	product["dell"] = 30000
	product["hp"] = 20000

	fmt.Println("product = ", product)

	//delete
	delete(product, "dell")
	fmt.Println("product = ", product)

	//update
	product["Macbook"] = 50000
	fmt.Println("product = ", product)

	varlue1 := product["Macbook"]
	fmt.Println("value1 = ", varlue1)

	courseName := map[string]string{
		"key1": "Golang",
		"key2": "Java",
		"key3": "Python",
		"key4": "C++",
	}
	fmt.Println("courseName = ", courseName)
}
