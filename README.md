# Go Games

Commandline games for practice with programming in Go. 

## To Play
Make sure you have Go installed

2 ways to play:

  * Run `go build main.go` and then `./main` and experience utter awesomeness.

  * Runt the command `go run main.go` to compile and rune the program immediately.


## Games

1. Number Game 

  This game is good practice getting the basics down of getting user input and using Go's version of a while loop to keep the game running until completion.

2. Just tell me a joke

  This was to practice making not only http requests, but a request that returns data in JSON format. It is different than other languages when it comes to putting data into JSON format.

3. Just tell me a fun fact

  This was to practice simple http requests that return plain text. I used Go's built in `http.Get()` method and passed the results off to Go's `io.Copy()` method to print the results to the terminal.

4. Blackjack

  This game uses nested structs to simulate a game of Blackjack. I used receiver functions to manipulate the structs. I also put some of the game logic into a different package than the main to get used to importing local packages rather than standard Go packages.
