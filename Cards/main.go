package main

import "fmt"

func main() {
	fmt.Println("...........................New Project.............................")
	var cards deck = newDeckFromFile("mycards")
	cards.printCards()
	// cards.saveToFile("mycards")
	// var firstHand, secondHand deck = deal(cards, 5)
	// firstHand.printCards()
	// secondHand.printCards()
	// var gretings string = "Hello"
	// fmt.Println(cards.toString())
}
