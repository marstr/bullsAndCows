package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/marstr/bullsAndCows"
)

func main() {
	answer := bullsAndCows.Number{
		uint8(rand.Intn(10)),
		uint8(rand.Intn(10)),
		uint8(rand.Intn(10)),
		uint8(rand.Intn(10)),
	}

	guesses := 0
	var empty bullsAndCows.Number

	for {
		guess := bullsAndCows.Number{}

		for guess == empty {
			var err error
			var raw string
			fmt.Print("Guess: ")
			fmt.Scanln(&raw)
			guess, err = bullsAndCows.ParseNumber(raw)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}

		guesses++

		fmt.Println("Cows: ", guess.Cows(answer))
		fmt.Println("Bulls: ", guess.Bulls(answer))

		if answer == guess {
			break
		}
	}

	fmt.Printf("You deduced the answer in %d guesses\n", guesses)
}
