package main

import "fmt"

/*
TO run - go run main.go employee.go
*/
func main() {
	var emp1 = Employee{name: "Hemant", leavesTaken: 4}

	fmt.Println(emp1)
	var left = emp1.getLeavesLeft()

	fmt.Println(left)

}
