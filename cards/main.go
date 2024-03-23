package main

func main() {

	//cards := newDeck()
	//err := cards.saveToFile("cards")
	//if err != nil {
	//	return
	//}

	cards := newDeck()
	cards.shuffleCards()
	cards.print()
}
