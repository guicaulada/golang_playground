package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }
	// var colors map[string]string // same as above but with empty values
	// colors := make(map[string]string) // same as above
	// colors["white"] = "#ffffff" // how to declare an element
	// fmt.Println(colors)

	// delete(colors, "white") // how to delete an element
	// fmt.Println(colors)

	// numbers := make(map[int]string) // it also works for ints
	// numbers[10] = "Ten"
	// fmt.Println(numbers)

	// colors.white // this cant be used must use square braces

	// Map:
	//   - All keys and values must be the same type
	//   - Keys are indexed and we can iterate over them
	//   - Represents a collection of related properties
	//   - Don't need to know all the keys to compile
	//   - REFERENCE TYPE
	// Strict:
	//   - Values can have different types
	//   - Keys are not indexed and don't support indexing
	//   - Represents an "object" with a lot of different properties
	//   - Need to know all the properties to compile
	//   - VALUE TYPE

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
