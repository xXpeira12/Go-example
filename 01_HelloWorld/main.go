package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xXpeira12/Go-example/01_HelloWorld/sirasit"
)

func main() {
	fmt.Println("Hello, World")
	id := uuid.New()
	fmt.Println("%s", id)
	sirasit.SayHelloSirasit()
	sirasit.SayTest()
}