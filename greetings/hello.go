package main

import (
	"fmt"

	"greetings_GO/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Dours")
	fmt.Println(message)
}
