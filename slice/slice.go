package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Golang", "Java", "Python", "C++"}
	fmt.Println("courseName = ", courseName)
	courseName = append(courseName, "Rust")
	fmt.Println("courseName = ", courseName)

	courseWeb := courseName[0:3]
	fmt.Println("courseWeb = ", courseWeb)

}
