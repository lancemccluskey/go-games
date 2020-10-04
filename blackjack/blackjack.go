package blackjack

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Card struct represents a single card in a 52 card deck
// It contains a name, ie Ace of Spades
// And a value, ie King == 13
type Card struct {
	name  string
	value int
}

// User struct represents the player and the dealer
// It contains a Hand which is Deck struct to hold the cards
// It contains a Total value which holds the total values of the cards
type User struct {
	Hand  Deck
	Total int
}

// Deck struct is a slice of Card structs
// and represents a 52 card playing deck
type Deck []Card

// AddCard appends a new card to the Deck slice
func (d *Deck) AddCard(c Card) {
	*d = append((*d), c)
}

// Shuffle uses the current time in nanoseconds as the seed
// to randomly switch the positions of the cards
func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// NewDeck returns a new Deck instance with 52 Cards in it
func (d *Deck) NewDeck() {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for j, value := range cardValues {
			c := Card{
				name:  value + " of " + suit,
				value: j + 1,
			}

			d.AddCard(c)
		}
	}
}

// Deal returns a new Deck instance of the specified handSize
// This is used to Deal initial starting hands to the player and dealer
// And also to deal 1 card a time when a player wants to hit
func (d *Deck) Deal(handSize int) Deck {
	hand, newDeck := (*d)[:handSize], (*d)[handSize:]
	*d = newDeck
	return hand
}

// CalculateTotal is method that loops through all the Cards
// in the given Users Hand and sums up the values
func (u *User) CalculateTotal() {
	for _, card := range u.Hand {
		if strings.Contains(card.name, "Ace") {
			u.Total += 11
		} else {
			u.Total += card.value
		}
	}

	if u.Total > 21 {
		u.Total -= 10
	}
}

// PrintHand prints the current Cards in the Users Hand to the terminal
func (u User) PrintHand() {
	fmt.Printf("Your hand: ")
	for _, card := range u.Hand {
		fmt.Printf("%v, ", card.name)
	}
	fmt.Println("")
}
