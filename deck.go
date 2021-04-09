package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck
//which is a slice of strings
//extends and borrows all of the behavior of a slice of string
type deck []string

func newDeck() deck { //return value of type deck (array of string)
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, suits+" of "+values)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) { //return two values of type deck.
	return d[:handSize], d[handSize:] //return everything from beginning TO handsize AND everything from handsize to end.

}

func (d deck) print() { //reciever
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //permission 0666 anyone can read/write to file.

}

func (d deck) toString() string { //reciever function.

	return strings.Join([]string(d), ",")
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//Option 1: log the error and return a call to newDeck()
		//Option 2: log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)

	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

//shuffle
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) //generating a seed value //review Source package again.
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)           //get length of slice
		d[i], d[newPosition] = d[newPosition], d[i] //fancy way of swaping positions

	}
}

//for each index, card in cards

//generate a random number between 0 and len(cards)-1

//swap current card and the card at cards[randomnumber]
