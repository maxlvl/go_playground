package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return error message
	if name == "" {
		return "", errors.New("empty name")
	} // Returns a greeting that embeds the name in the message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)
	// loop through the received slice of names, calling
	// the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		
		// In the map, associate the retrieved message with the name
		messages[name] = message
	}
	
	return messages, nil
}

// init sets initial value for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random
func randomFormat() string {
	// A slice of message formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v. Well met!",
	}
	
	// return a randomly selected greeting each time	
	return formats[rand.Intn(len(formats))]
	
}