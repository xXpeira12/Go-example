package main

import "fmt"

func sayHello(name string, round int) {
	for i := 0; i < round; i++ {
		fmt.Println("Hello " + name)
	}
}

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Function Example")

	sayHello("John", 2)
	fmt.Println(add(1, 2))
}