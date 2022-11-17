package main

type Employee struct {
	name        string
	leavesTaken int
	salary      int
}

func (emp Employee) getLeavesLeft() int {
	var leavesLeft = 20 - emp.leavesTaken
	return leavesLeft
}

/*
* If salary-bills > 1L, tax = 10%
* If salary-bills < 1L, no tax
 */
func (emp Employee) getInHandSalary(bills int) float64 {
	var taxableSalary = emp.salary - bills

	println("The total Taxable salary = ", taxableSalary)
	if taxableSalary > 100_000 {
		println("10% tax will be applied on taxable salary = ", taxableSalary)
		return float64(taxableSalary) * float64(0.9)
	} else {
		println("No tax is applicable on taxable salary =", taxableSalary)
		return float64(taxableSalary)
	}
}
