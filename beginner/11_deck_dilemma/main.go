package main

// There is a math riddle that says:
// Your friend choses at random a card from a standard deck of 52 cards, and keeps this card concealed.
// You have to guess which of the 52 cards it is.
//
// Before you guess, you can ask your friend one of the following three questions:
// - is the card red?
// - is the card a face card? (Jack, Queen or King)
// - is the card the ace of spades?
//
// Your friend will answer truthfully.
// Which question would you ask that gives you the best chance of guessing the correct card?
//
// This program simulates that interaction 100.000 times for each question to prove that
// the probability for guessing the correct card is the same, doesn't matter what question you ask.
// You could calculate this manually using mathematical probability, but I found this method much more fun.
// And it helps you stop doubting yourself since the results are counter-intuitive.

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type card struct {
	value string
	suit  string
}

type deck []card

type fraction struct {
	n int
	d int
}

func main() {
	ch := make(chan string)
	d := newDeck()

	guessesTotal := 300000
	guessesEach := guessesTotal / 3

	count := map[string]int{
		"isRed":  0,
		"isFace": 0,
		"isAce":  0,
	}

	for i := 0; i < guessesEach; i++ {
		go guessByIsRed(d, ch)
	}

	for i := 0; i < guessesEach; i++ {
		go guessByIsFace(d, ch)
	}

	for i := 0; i < guessesEach; i++ {
		go guessByIsAce(d, ch)
	}

	for i := 0; i < guessesTotal; i++ {
		key := <-ch
		if v, ok := count[key]; ok {
			count[key] = v + 1
		}
	}

	isRedProb := float64(count["isRed"]) / float64(guessesEach)
	isFaceProb := float64(count["isFace"]) / float64(guessesEach)
	isAceProb := float64(count["isAce"]) / float64(guessesEach)
	isRedFraction := roundToFraction(isRedProb)
	isFaceFraction := roundToFraction(isFaceProb)
	isAceFraction := roundToFraction(isAceProb)
	fmt.Println("probabilities:")
	fmt.Printf("  isRed:  %.4f ~ %d/%d \n", isRedProb, isRedFraction.n, isRedFraction.d)
	fmt.Printf("  isFace: %.4f ~ %d/%d \n", isFaceProb, isFaceFraction.n, isFaceFraction.d)
	fmt.Printf("  isAce:  %.4f ~ %d/%d \n", isAceProb, isAceFraction.n, isAceFraction.d)
	fmt.Printf("expected: %.4f ~ 1/26 \n", 1.0/26.0)
}

func roundToFraction(x float64) fraction {
	precision := 0.000001
	n := math.Floor(x)
	x = x - n
	if x < precision {
		return fraction{n: int(n), d: 1}
	} else if 1-precision < x {
		return fraction{n: int(n + 1), d: 1}
	}

	lower := fraction{n: 0, d: 1}
	upper := fraction{n: 1, d: 1}

	for {
		middle := fraction{
			n: lower.n + upper.n,
			d: lower.d + upper.d,
		}

		if float64(middle.d)*(x-precision) < float64(middle.n) {
			upper.n = middle.n
			upper.d = middle.d
		} else if float64(middle.d)*(x-precision) > float64(middle.n) {
			lower.n = middle.n
			lower.d = middle.d
		} else {
			return fraction{
				n: 1, d: middle.d / (int(n)*middle.d + middle.n),
			}
		}
	}
}

func guessByIsRed(d deck, ch chan string) {
	c := d.random()
	if c.isRed() {
		d = d.reds()
	} else {
		d = d.blacks()
	}
	sendIfTrue(ch, "isRed", d.random() == c)
}

func guessByIsFace(d deck, ch chan string) {
	c := d.random()
	if c.isFace() {
		d = d.faces()
	} else {
		d = d.nonFaces()
	}
	sendIfTrue(ch, "isFace", d.random() == c)
}

func guessByIsAce(d deck, ch chan string) {
	c := d.random()
	if c.isAce() {
		d = d.aces()
	} else {
		d = d.nonAces()
	}
	sendIfTrue(ch, "isAce", d.random() == c)
}

func sendIfTrue(ch chan string, s string, b bool) {
	if b {
		ch <- s
		return
	}
	ch <- ""
}

func (c card) isAce() bool {
	return c.value == "Ace"
}

func (c card) isFace() bool {
	return c.value == "Jack" || c.value == "Queen" || c.value == "King"
}

func (c card) isRed() bool {
	return c.suit == "Diamonds" || c.suit == "Hearts"
}

func (d deck) reds() deck {
	cards := deck{}
	for _, c := range d {
		if c.isRed() {
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) blacks() deck {
	cards := deck{}
	for _, c := range d {
		if !c.isRed() {
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) aces() deck {
	cards := deck{}
	for _, c := range d {
		if c.isAce() {
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) nonAces() deck {
	cards := deck{}
	for _, c := range d {
		if !c.isAce() {
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) faces() deck {
	cards := deck{}
	for _, c := range d {
		if c.isFace() {
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) nonFaces() deck {
	cards := deck{}
	for _, c := range d {
		if !c.isFace() {
			cards = append(cards, c)
		}
	}
	return cards
}

func (d deck) random() card {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(d))
	return d[r]
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, card{value, suit})
		}
	}

	return cards
}
