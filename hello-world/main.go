package main

import (
	"fmt"
	controlFlow "hello-world/src"
	furtherexplored "hello-world/src/furtherexplored"
)

type person struct {
	fullName string
	dob      string
}

func add(x int, y int) int {
	return x + y
}

func (p person) sayHello() {
	fmt.Println("Hi, my name is", p.fullName)
	fmt.Println("\t my age is", p.dob)
}

func init() {
	fmt.Println("Begin initialization")
}

func main() {
	var fullName string = "Abdullah"
	age := 24
	x := 50
	const m float32 = 42.43
	myName := person{
		"Abdullah Umar Babsail",
		"24 Feb 1989",
	}
	// fmt.Printf("Hello My Name %s! \t type %T \n", fullName, fullName)
	// fmt.Printf("My age %b\n", age)
	// fmt.Printf("X %d and %v \n", x, m)
	fmt.Println("Hello My Name", fullName, "my age", age)
	fmt.Printf("\t X %d and %v \n", x, m)
	fmt.Println("\n", add(42, 13))
	fmt.Println("test edit by nano \n")

	myName.sayHello()
	furtherexplored.Fascinating()
	controlFlow.ControlFlow()
}

