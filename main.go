package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
			fmt.Println("Not implemented yet, sorry")
		case 3:
			fmt.Println("Not implemented yet, sorry")
		case 4:
			fmt.Println("Not implemented yet, sorry")
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
