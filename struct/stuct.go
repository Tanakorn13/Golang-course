package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	// employee1 := employee{
	// 	employeeID:   "1",
	// 	employeeName: "John Doe",
	// 	phone:        "123-456-7890",
	// }
	// fmt.Println("employee1 = ", employee1)

	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "1",
		employeeName: "John Doe",
		phone:        "123-456-7890",
	}
	employee2 := employee{
		employeeID:   "2",
		employeeName: "Jane Smith",
		phone:        "987-654-3210",
	}
	employee3 := employee{
		employeeID:   "3",
		employeeName: "Mark Johnson",
		phone:        "555-555-5555",
	}
	employeeList = append(employeeList, employee1, employee2, employee3)
	// employeeList[0] = employee{
	// 	employeeID:   "1",
	// 	employeeName: "John Doe",
	// 	phone:        "123-456-7890",
	// }
	// employeeList[1].employeeID = "2"
	// employeeList[1].employeeName = "Jane Smith"
	// employeeList[1].phone = "987-654-3210"

	// employeeList[2] = employee{
	// 	employeeID:   "3",
	// 	employeeName: "Mark Johnson",
	// 	phone:        "555-555-5555",
	// }
	fmt.Println("employeeList = ", employeeList)
}
