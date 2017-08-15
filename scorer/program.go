package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/marstr/bullsAndCows"
)

func main() {
	answer := generateNoRepeats()

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

func generateRepeats() (generated bullsAndCows.Number) {
	generated = bullsAndCows.Number{
		uint8(rand.Intn(10)),
		uint8(rand.Intn(10)),
		uint8(rand.Intn(10)),
		uint8(rand.Intn(10)),
	}
	return
}

func generateNoRepeats() (generated bullsAndCows.Number) {
	options := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := len(generated) - 1; i >= 0; i-- {
		selected := rand.Intn(len(options))
		generated[i] = options[selected]
		options = append(options[:selected], options[selected+1:]...)
	}
	return
}
