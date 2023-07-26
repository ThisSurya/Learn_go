package greetings

import "fmt"

func Hello(namamu string) string {
	message := fmt.Sprintf("Good morning, %v.", namamu)
	return message
}
