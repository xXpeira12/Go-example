package main

import (
	"fmt"
)

type Student struct {
	name string
	age int
	weght float32
	height float32
	grade string
	address Address
}

type Address struct {
	street string
	city string
	zipCode string
}

func main() {
	var student1 Student
	student1.name = "John Doe"
	student1.age = 25
	student1.weght = 70.5
	student1.height = 180.5
	student1.grade = "A"
	student1.address = Address{ 
		street: "123 Main St", 
		city: "New York", 
		zipCode: "10001" ,
	}

	fmt.Println(student1)
}