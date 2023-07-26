package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Surya")
	fmt.Println(message)
}
