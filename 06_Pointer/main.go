package main

import "fmt"

type Person struct {
	Name string
}

func changeValue(x *int) {
	*x = 10
}

func changeName(p *Person) {
	p.Name = "John"
}

func main() {
	// Declare a pointer variable
	var ptr *int

	// Declare a variable
	i := 10

	// Assign the address of i to ptr
	ptr = &i

	fmt.Println("Value of i:", i)
	fmt.Println("Address of i:", &i)
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Address of ptr:", &ptr)
	fmt.Println("Value of *ptr:", *ptr)

	// Change the value of i using pointer
	changeValue(ptr)
	fmt.Println("Value of i:", i)

	// Declare a pointer variable
	person := Person{Name: "Alice"}
	changeName(&person)
	fmt.Println("Name of person:", person.Name)
}