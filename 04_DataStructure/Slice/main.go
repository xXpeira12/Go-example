package main

import "fmt"

func main() {
	fmt.Println("Slice Example")

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	// Slicing a slice
	subSlice := mySlice[1:3]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice))
	fmt.Println(cap(subSlice))

	// Appending data to slice
	mySlice = append(mySlice, 6, 7, 8)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}