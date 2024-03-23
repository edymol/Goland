package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck
// which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// This function reads as: A function that is of type deck in instance d
// with a data type integer that will have two return values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize],
		d[handSize:]
}

// this function has a receiver with type of deck with the name of toString and returns a
// the strings,Join imported and serves to join all the separated string in deck to one
// string separated by ,
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//func (d deck) saveToFile(filename string) error {
//	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
//}

func (d deck) WriteTo(w io.Writer) (n int64, err error) {
	data := d.toString()
	nBytes, err := io.WriteString(w, data)
	if err != nil {
		return 0, err
	}
	return int64(nBytes), nil
}

func (d deck) saveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = d.WriteTo(file)
	return err
}

func (d *deck) ReadFrom(r io.Reader) (n int64, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	*d = strings.Split(string(data), ",")
	return int64(len(data)), nil
}

func newDeckFromFile(filename string) deck {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var d deck
	_, err = d.ReadFrom(file)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return d
}

func (d deck) shuffleCards() {
	source := rand.NewSource(time.Now().UnixNano()) // this creates a random number returns an int64 to be able to use truly random numbers
	r := rand.New(source)                           // creates a new random number using the source passed above
	for i := range d {
		newPosition := r.Intn(len(d) - 1) // calling the r.Intn function from the source to create a new position

		d[i], d[newPosition] = d[newPosition], d[i] // swapping same as Python
	}
}
