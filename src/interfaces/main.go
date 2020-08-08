package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	Name string
}
type spanishBot struct {
	Name string
}

func (eb englishBot) getGreeting() string {
	// Customized business logic for EnglishBot
	return "Hi There! " + eb.Name
}

func (sb spanishBot) getGreeting() string {
	// Customized business logic for SpanishBot
	return "Hola from " + sb.Name
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Golang doesn't support overloading
// func printGreeting(sb spanishBot) {
//	fmt.Println(sb.getGreeting())
//}

func main() {
	eb := englishBot{Name: "John"}
	sb := spanishBot{Name: "Bob"}

	printGreeting(eb)
	printGreeting(sb)
}
