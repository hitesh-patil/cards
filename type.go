package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type cardArr []string

func newCards() cardArr {
	cards := cardArr{}

	cardTypes := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, types := range cardTypes {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+types)
		}
	}

	return cards
}

func (c cardArr) printArr() {
	for i, cardName := range c {
		fmt.Println(i, cardName)
	}
}

func (c cardArr) toString() string {
	return strings.Join(c, ",")
}

func (c cardArr) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(c.toString()), 0644)
}

func byteToString(byteArr []byte) string {
	return string(byteArr)
}

func newDeckFromFile(fileName string) cardArr {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return cardArr(strings.Split(byteToString(content), ","))
}
