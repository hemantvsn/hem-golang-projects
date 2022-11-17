package main

import (
	"fmt"
	"io/ioutil"
)

type Deck struct {
	cards []Card
}

func newDeck() Deck {
	return Deck{cards: getAllCards()}
}

// TOSTRING method of object
func (d Deck) String() string {
	var toString = ""
	for i, card := range d.cards {
		toString = toString + fmt.Sprintln(i, card)
	}
	return toString
}

func (d Deck) splitInTwo(splitIndex int) (Deck, Deck) {
	var split1 = Deck{cards: d.cards[:splitIndex]}
	var split2 = Deck{cards: d.cards[splitIndex:]}
	return split1, split2
}

func (d Deck) saveToFile(fileName string) {
	var byteArr = []byte(d.String())
	var err = ioutil.WriteFile(fileName, byteArr, 0777)
	if err != nil {
		fmt.Println("Couldn't write DECK TO FILE", err)
	} else {
		fmt.Println("Deck successfully written to file ", fileName)
	}
}
