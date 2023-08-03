package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func main() {

	var eb englishBot
	var sb spanishBot
	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hello there"
}
