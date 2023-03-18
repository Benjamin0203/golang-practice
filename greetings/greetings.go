package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    // message := fmt.Sprintf("Hi, %v. Welcome!", name)
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

//init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}
// randomFormat returns one of a set of greeting messages.
//randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).

// The returned message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}
