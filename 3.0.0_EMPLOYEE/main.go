package main

import "fmt"

/*
TO run - go run main.go employee.go
*/
func main() {
	var emp1 = Employee{
		name:        "Hemant",
		leavesTaken: 4,
		salary:      230000,
		address: []Address{
			Address{city: "Pune", country: "India"},
			Address{city: "Noida", country: "IND"},
		},
	}

	fmt.Println(emp1)
	var left = emp1.getLeavesLeft()
	fmt.Println("Leaves left = ", left)

	var bills = 45000
	var inHandSalary = emp1.getInHandSalary(bills)
	fmt.Println("InHand salary = ", inHandSalary)

	// when you declare a variable and not initialize it
	// go asssigns it 0 VALUE
	var emp2 Employee
	fmt.Printf("%+v", emp2)
	fmt.Printf("%+v", emp1)
}
