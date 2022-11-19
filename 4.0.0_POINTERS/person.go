package main

import "fmt"

type Person struct {
	id     int
	name   string
	salary float32
}

type Address struct {
	city    string
	country string
}

func (p Person) updateNameOLD(newName string) {
	fmt.Println("Updating ", p, " name to : ", newName)
	var ptr = &p
	fmt.Println("The address of p =", &ptr)
	p.name = newName
}

// To access the field X of a struct when we have the struct pointer p
// we could write (*p).X.
//
//	However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
func (ptr *Person) updateNameNEW(newName string) {
	var x = &ptr
	fmt.Println(x)

	var y = *ptr
	fmt.Println(y)

	(*ptr).name = newName
	// This is a shortcut
	// To access the field X of a struct when we have the struct pointer p
	// we could write (*p).X. However, that notation is cumbersome,
	//so the language permits us instead to  write just p.X, without the explicit dereference.
	ptr.name = newName

}

func (p Person) print() {
	fmt.Printf("PERSON = %+v", p)
	fmt.Println()
}
