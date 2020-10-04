package blackjack

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Card struct {
	name  string
	value int
}

type User struct {
	Hand  Deck
	Total int
}

type Deck []Card

func (d *Deck) AddCard(c Card) {
	*d = append((*d), c)
}

func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

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

func (d *Deck) Deal(handSize int) Deck {
	hand, newDeck := (*d)[:handSize], (*d)[handSize:]
	*d = newDeck
	return hand
}

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

func (u User) PrintHand() {
	fmt.Printf("Your hand: ")
	for _, card := range u.Hand {
		fmt.Printf("%v, ", card.name)
	}
	fmt.Println("")
}
