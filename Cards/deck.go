package main

import "fmt"

type deck []string

func newDeck() deck {
	var cards deck
	var cardSuits = []string{"spades", "hearts", "diamonds", "clubs"}
	var cardValues = []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	var firstHandSize deck = d[:handSize]
	var secondHandSize deck = d[handSize:]

	return firstHandSize, secondHandSize
}

func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
