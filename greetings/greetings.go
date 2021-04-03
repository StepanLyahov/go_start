package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Add an init function to seed the rand package with the current time.
// Go executes init functions automatically at program startup,
// after global variables have been initialized.
// For more about init functions, see Effective Go.
// https://golang.org/doc/effective_go#init
func init() {
	log.SetPrefix("info: ")
	log.SetFlags(log.Lshortfile)
	time := time.Now().UnixNano()
	log.Print(time)

	rand.Seed(time)
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if name is void, return void string and error with description
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) //  more about maps (https://blog.golang.org/maps)

	for i, name := range names {
		log.Printf("Position: %v", i)
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
