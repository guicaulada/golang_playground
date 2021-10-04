package main

import "fmt"

// custom type called bot
// if you are a type in this program
// that calls the same functions defined in this interface
// you are also a member of this interface
//
// if you have a function getGreeting() that returns a string
// you are also a bot
//
// we cant create a value using an interface type
// interfaces are not generic types, Go doesn't support generic types
// interfaces are implicit we don't need any code to link an interface with it's members
// interfaces are a contract between different functions and different types
// interfaces are tough, it's hard to read them and hard to write them
// interfaces are not required but highly recommended
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}

// we can't declare a function with the same name with different argument types
// this why we need interfaces
//
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
//
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// we can omit the variable eb for (eb englishBot) since we are not using it
func (englishBot) getGreeting() string {
	return "Hi there!"
}

// we can omit the variable sb for (sb spanishBot) since we are not using it
func (spanishBot) getGreeting() string {
	return "Hola!"
}
