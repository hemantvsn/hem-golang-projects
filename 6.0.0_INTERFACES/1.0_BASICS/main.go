package main

import "fmt"

// This is an interface with method declaration
type Shape interface {
	// Implementing class must implement this method
	// where there are no args and returns int
	getArea() int
}

// No extends or implements interface like JAVA
type Rectange struct {
	length  int
	breadth int
}

// No extends or implements interface like JAVA
type Square struct {
	side int
}

type Circle struct {
	radius int
}

// SQUARE class is having getArea method with same signature
// ie EMPTY ARGS and RETURN INT
// Hence it automatically extends interface
func (s Square) getArea() int {
	return s.side * s.side
}

// RECTANGE class is having getArea method with same signature
// ie EMPTY ARGS and RETURN INT
// Hence it automatically extends interface
func (r Rectange) getArea() int {
	return r.breadth * r.length
}

// Please observe that getArea method of Circle is different
// SHAPE needs getArea method with 0 args and int return type
// But
// Circle has getArea method with 0 args and FLOAT return type
// Hence circle class doesn't implements SHAPE
func (c Circle) getArea() float32 {
	return float32(c.radius) * 3.14
}

// This is a generic method - which accepts SHAPE ARG
// As both square and rectange qualify for extending interface,
// we can call printShape() on both square and rectange variables
// Since circle doesn't implement SHAPE, we cannot call printShape on cirlce var
func printShape(s Shape) {
	fmt.Printf("%+v", s)
	fmt.Println()
}

func main() {
	var rect = Rectange{length: 10, breadth: 5}
	var sqr = Square{side: 8}
	var cir = Circle{radius: 10}

	fmt.Println("The area of rectange = ", rect, "is ", rect.getArea())
	fmt.Println("The area of square = ", sqr, "is ", sqr.getArea())
	fmt.Println("The area of circle = ", cir, "is ", cir.getArea())

	// Can call printShape()
	printShape(rect)
	printShape(sqr)

	// GIVING COMPILATION ERROR - as circle doesn't implement SHAPE
	printShape(cir)
}
