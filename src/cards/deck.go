package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of `deck` which is a slice of strings
type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) lastElement() (string, error) {
	if len(d) > 0 {
		return d[len(d)-1], nil
	} else {
		return "", nil
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	if bs, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else {
		aDeck := deck(strings.Split(string(bs), ","))
		return aDeck, nil
	}
}
func (d deck) shuffle() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPosition := rnd.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
