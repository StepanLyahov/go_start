package main

import (
	"fmt"
	"go_start/greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Lshortfile)

	// Get a greeting message and print it.
	message, error := greetings.Hello("")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
