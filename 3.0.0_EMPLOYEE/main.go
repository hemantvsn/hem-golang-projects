package main

import "fmt"

/*
TO run - go run main.go employee.go
*/
func main() {
	var emp1 = Employee{name: "Hemant", leavesTaken: 4, salary: 230000}

	fmt.Println(emp1)
	var left = emp1.getLeavesLeft()
	fmt.Println("Leaves left = ", left)

	var bills = 45000
	var inHandSalary = emp1.getInHandSalary(bills)
	fmt.Println("InHand salary = ", inHandSalary)

}
