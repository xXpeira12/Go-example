package main

import "fmt"

func main() {
	fmt.Println("Array Example")

	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println(myArray[1])
	fmt.Println(myArray)

	// Looping through array
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}