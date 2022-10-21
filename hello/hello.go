package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}