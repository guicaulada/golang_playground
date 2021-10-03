package main

import "fmt"

func main() {
	// a slice is made by 3 elements
	//  - pointer to first element of an array
	//  - capacity
	//  - length
	// so when passing a slice as an argument to a function
	// Go still makes a copy, but the pointer is the same
	// this happens to other types as well
	// those are called Reference Types
	//
	// Reference Types: slices, maps, channels, pointers, functions
	// Value Types: int, float, string, bool, struct

	words := []string{"Hi", "There", "How", "Are", "You"}
	updateWords(words)
	fmt.Println(words)
}

func updateWords(s []string) {
	s[0] = "Bye"
}
