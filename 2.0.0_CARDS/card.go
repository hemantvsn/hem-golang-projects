package main

var SUITS = []string{"clubs (♣)", "diamonds (♦)", "hearts (♥)", "spades (♠)"}
var VALUES = []string{"ACE", "TWO", "THREE", "FOUR", "FIVE"}

type Card struct {
	value string
	suit  string
}

/*
*
Similar to overriding toString method in JAVA
*/
func (c Card) String() string {
	return c.value + " OF " + c.suit
}

func getAllCards() []Card {
	var cards = []Card{}
	for _, suit := range SUITS {
		for _, value := range VALUES {
			cards = append(cards, Card{value: value, suit: suit})
		}
	}
	return cards
}
