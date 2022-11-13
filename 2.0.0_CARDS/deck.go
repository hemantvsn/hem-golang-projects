package main

import "fmt"

type Deck []string

func newDeck() Deck {
	return Deck{"ONE", "TWO"}
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
