package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// functions must have return value type after the arguments
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suits := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suits)
		}
	}

	return cards
}

// (d deck) is the receiver
// we are basically adding a method to the type deck
// d is equivalent to this
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// we can use slices like on python
// functions can return multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

// we can convert types by using <type>(<var>)
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log the error and return a call to newDeck()
		// or log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// convert the slice of strings to a deck
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1) // len() gets the length of a slice
		d[i], d[newPos] = d[newPos], d[i]
	}
}
