package main

import "fmt"

func main() {
	fmt.Println("...........................New Project.............................")
	var cards deck = newDeck()
	var firstHand, secondHand deck = deal(cards, 5)
	firstHand.printCards()
	secondHand.printCards()
}
