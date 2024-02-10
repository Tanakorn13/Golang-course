package main

import "fmt"

var productName [4]string

func main() {
	productName[0] = "Macbook"
	productName[1] = "Dell"
	productName[2] = "HP"
	productName[3] = "Asus"

	fmt.Println(productName)

	price := [4]float32{10000}

	price[1] = 2000
	price[2] = 3000
	price[3] = 4000

	fmt.Println(price)
}