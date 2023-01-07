package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)


func main() {
	// Set properties of a predefined logger, including 
	// the log entry prefix and a flag to disable printing the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Max", "Go", "Rust", "Javascript"}
	// get a greeting message and print it
	messages, err := greetings.Hellos(names)
	// if an error was returned, print it to the console and exit the programj
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned map of messages to the console
	fmt.Println(messages)
}

