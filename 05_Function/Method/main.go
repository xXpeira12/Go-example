package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

// Method with a receiver of type Student
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

type IntSlice []int

func (is IntSlice) Sum() int {
	sum := 0
	for _, v := range is {
		sum += v
	}
	return sum
}

func main() {
	student := Student{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	fmt.Println(student.FullName())

	intSlice := IntSlice{1, 2, 3, 4, 5}
	fmt.Println(intSlice.Sum())
}