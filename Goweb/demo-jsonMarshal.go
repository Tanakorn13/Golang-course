package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "John Doe", "123-456-7890", "jdoe@go.dev"})
	fmt.Println("Data: ", string(data))
}
