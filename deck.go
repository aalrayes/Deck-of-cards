package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, numCards int) (deck, deck) {
	return d[:numCards], d[numCards:]
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeckFromFile(filename string) deck {

	fileContent, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Print("error", err)
		os.Exit(1)
	}
	deckFromFile := strings.Split(string(fileContent), ",")
	return deck(deckFromFile)
}

func (d deck) shuffle() deck {
	time := time.Now().UnixNano()
	source := rand.NewSource(time)
	r := rand.New(source)

	for i := range d {
		newInd := r.Intn(len(d) - 1)
		d[i], d[newInd] = d[newInd], d[i]
	}

	return d
}
