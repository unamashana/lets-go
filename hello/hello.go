package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")

	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}