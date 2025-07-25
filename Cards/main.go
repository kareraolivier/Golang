package main

import "fmt"

func main() {
	fmt.Println("new project")
	// var card string = newCard()
	var cards = deck{"Olivier", newCard()}
	cards = append(cards, "karera")
	cards.print()
}
func newCard() string {
	return "five of diamond"
}
