package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type "deck" which is a slice of type "string"
type deck []string

// receiver fn that should create and return a new deck (list) of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// add a new card of value, suit
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// create receiver fn (ie method to the calling type instance)
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// create a receiver fn to deal a hand of cards by selecting a range of the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func WriteFile(filename string, data []byte, perm os.FileMode) error
func (d deck) saveToFile(filename string) error {
	// type conversion
	// 0666 anyone can read and write to this file
	data := []byte(d.toString())
	return ioutil.WriteFile(filename, data, 0666)
}

func readFromFile(filename string) deck {
	// func ReadFile(filename string) ([]byte, error) --> returns []byte
	bs, e := ioutil.ReadFile(filename)
	if e != nil {
		// return the error and exit the program
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
	// it is better to think of the byte slice as a string
	// func Split(s, sep string) []string
	strSlice := strings.Split(string(bs), ", ") // conversion to []string
	return deck(strSlice)
}

// for each index, card in cards
// generate a random number between 0 and (cards.length - 1)
// swap current card with card at index

// helper fns
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) shuffle() {
	// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
	// since January 1, 1970 UTC. The result is undefined if the Unix time in
	// nanoseconds cannot be represented by an int64
	t := time.Now().UnixNano()

	// create seed value for randomness using current local time
	source := rand.NewSource(t)
	r := rand.New(source)

	for i := range d {
		// pass the source seed to the rand fn
		randomIndex := r.Intn(len(d) - 1)
		// swap the cards
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}
