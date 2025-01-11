package main

import (
	"fmt"
)

func main() {
	// If else
	fmt.Println("If else")
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// Switch
	fmt.Println("Switch")
	y := 2
	switch y {
	case 1:
		fmt.Println("y is 1")
	case 2:
		fmt.Println("y is 2")
	default:
		fmt.Println("y is not 1 or 2")
	}

	// Pre proccess statement
	fmt.Println("Pre proccess statement")
	num1 := 10
	num2 := 20
	if sumNum := num1 + num2; sumNum > 20 {
		fmt.Println("sumNum is greater than 20")
	} else {
		fmt.Println("sumNum is less than 20")
	}

	// For
	fmt.Println("For")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For range
	fmt.Println("For range")
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// While
	fmt.Println("While")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Do while
	fmt.Println("Do while")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}
}