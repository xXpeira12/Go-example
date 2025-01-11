package main

import (
	"fmt"
)

func main() {
	var fullName string
	var age int
	fullName = "John"
	age = 25

	var lastName string = "Doe"

	// Go จะให้ค่าเริ่มต้นให้ตัวแปรตามชนิดของข้อมูล
	address := "New York"

	const prefix = "Mr."

	fmt.Println("Full Name: ", prefix, fullName, lastName)
	fmt.Println("Age: ", age)
	fmt.Println("Address: ", address)
}