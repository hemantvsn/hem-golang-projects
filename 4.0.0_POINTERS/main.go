package main

import "fmt"

func main() {

	var p = Person{
		id:     123,
		name:   "Hemant",
		salary: 1000,
	}

	p.print()

	// OLD WAY
	var ptr = &p
	fmt.Println("The address of p =", &p)
	ptr.updateNameNEW("HEMYAAAA")
	p.print()

	// Other way, concise
	p.updateNameNEW("KAKAKAKAKAKAKA")
	p.print()

}
