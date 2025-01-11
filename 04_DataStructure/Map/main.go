package main

import "fmt"

func main() {
	fmt.Println("Map Example")

	myMap := make(map[string]int)

	// Add key-value pair
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	// Access and print value
	fmt.Println("Value of key 'one' is", myMap["one"])
	fmt.Println("Value of key 'two' is", myMap["two"])
	fmt.Println("Value of key 'three' is", myMap["three"])

	// Iterate over map
	fmt.Println("Iterating over map")
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Delete key-value pair
	delete(myMap, "two")
	fmt.Println("After deleting key 'two'")
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Check if key exists
	fmt.Println("Checking if key 'two' exists")
	value, exists := myMap["two"]
	if exists {
		fmt.Println("Value of key 'two' is", value)
	} else {
		fmt.Println("Key 'two' does not exist")
	}
}