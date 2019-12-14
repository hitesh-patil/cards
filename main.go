package main

func main() {
	// card := newCards()
	// // card.printArr()
	// card[13:26].printArr()
	// card.saveToFile("cardFile.txt")
	card := newDeckFromFile("cardFile.txt")
	// fmt.Println(card)
	card.shuffle()
	card.printArr()
}
