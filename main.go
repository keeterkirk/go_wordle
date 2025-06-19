// create account
// login with your username and password
// read word list
// opening screen
// guess
// guess check
// post guess render
// winning
// losing
// quitting
// statistics
// leaderboard

package main

import (
	"fmt"

	"github.com/keeterkirk/go_wordle/dictionary"
)

func main() {
	fmt.Println("Welcome to the Wordle game!")

	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Hello, " + name + "! Let's start the game.")
	fmt.Println(dictionary.RandomWord())
}
