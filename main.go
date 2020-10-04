package main

import (
	"encoding/json"
	"fmt"
	"go-games/blackjack"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type joke struct {
	ID     string
	Joke   string
	Status int
}

func main() {
	fmt.Println("Welcome to Go Games!")
	fmt.Println("Please select a game below to start playing")

	displayGameOptions()
	var selection int
	fmt.Scanln(&selection)

	for selection != 5 {
		switch selection {
		case 1:
			numberGame()
		case 2:
			displayJoke()
		case 3:
			displayFact()
		case 4:
			playBlackjack()
		}

		displayGameOptions()
		fmt.Scanln(&selection)
	}
}

func displayGameOptions() {
	fmt.Println("-----------------------")
	fmt.Println("1) Guess the Number")
	fmt.Println("2) I hate games, tell me a joke")
	fmt.Println("3) I hate games, tell me a fun fact")
	fmt.Println("4) Blackjack")
	fmt.Println("5) Exit")
}

func displayFact() {
	res, err := http.Get("http://numbersapi.com/random/trivia")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("-----------------------")
	fmt.Println("FUN FACT")
	io.Copy(os.Stdout, res.Body)
	fmt.Println("")
}

func displayJoke() {
	client := &http.Client{Timeout: 10 * time.Second}

	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}

	j := new(joke)
	json.NewDecoder(res.Body).Decode(j)
	fmt.Println("-----------------------")
	fmt.Println("JOKE")
	fmt.Println(j.Joke)

	res.Body.Close()
}

func numberGame() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100)
	fmt.Println("NUMBER GAME")
	fmt.Println("GOAL: Guess correct number between 0 and 100")
	fmt.Println("-----------------------")
	fmt.Printf("Enter a number: ")
	var userInput int
	fmt.Scanln(&userInput)

	for userInput != n {
		if userInput < n {
			fmt.Printf("\nToo low! Try again: ")
		} else {
			fmt.Printf("\nToo high! Try again: ")
		}
		fmt.Scanln(&userInput)
	}

	fmt.Println("Success! You guessed the correct number!")
}

func playBlackjack() {
	playingDeck := blackjack.Deck{}
	playingDeck.NewDeck()
	playingDeck.Shuffle()
	playingDeck.Shuffle()
	playingDeck.Shuffle()

	dealer := blackjack.User{}
	player := blackjack.User{}

	player.Hand = playingDeck.Deal(2)
	dealer.Hand = playingDeck.Deal(2)

	fmt.Println("-----------------------")
	fmt.Println("BLACKJACK")
	fmt.Println("Enter 1 to hit or 2 to stay")
	fmt.Println("-----------------------")
	player.PrintHand()

	dealer.CalculateTotal()
	if dealer.Total >= 21 {
		if dealer.Total > 21 {
			fmt.Println("Dealer busted!")
		} else {
			fmt.Println("Dealer has 21!")
		}

		fmt.Printf("Play again? (Y/N): ")
		var playAgain string
		switch playAgain {
		case "Y", "y", "Yes", "yes":
			playBlackjack()
		case "N", "n", "No", "no":
			return
		}
	}

	fmt.Printf("Hit(1) or Stay(2): ")
	var input int

	fmt.Scanln(&input)

	for input != 2 {
		player.Hand = append(player.Hand, playingDeck.Deal(1)...)
		player.PrintHand()

		fmt.Printf("Hit(1) or Stay(2): ")
		fmt.Scanln(&input)
	}

	player.CalculateTotal()
	dealer.CalculateTotal()

	if player.Total > 21 {
		fmt.Println("You busted!")
	} else if dealer.Total > 21 {
		fmt.Println("The dealer busted! You won!")
	} else if dealer.Total > player.Total {
		fmt.Println("You lost! Dealer won!")
	} else {
		fmt.Println("You won!")
	}

}
