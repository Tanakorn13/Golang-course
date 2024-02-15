package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101, "EmployeeName": "John Doe", "Tel": "123-456-7890", "Email": "jdoe@go.dev"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	fmt.Println("ID = ", e.ID)
	fmt.Println("EmployeeName = ", e.EmployeeName)
	fmt.Println("Tel = ", e.Tel)
	fmt.Println("Email = ", e.Email)
}
