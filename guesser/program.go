package main

import (
	"fmt"

	bc "github.com/marstr/bullsAndCows"
)

func main() {
	remaining := []bc.Number{}

	for i := 0; i < 10000; i++ {
		parsed, err := bc.ParseNumber(fmt.Sprint(i))
		if err == nil {
			remaining = append(remaining, parsed)
		}
	}

	fmt.Printf("Starting off, there are %d valid numbers you may have chosen.\n", len(remaining))

	guesses := 0

	for {
		if len(remaining) == 0 {
			fmt.Println("I give up!")
			return
		}

		guess := remaining[0]

		fmt.Println("My guess: ", guess)

		var bulls, cows uint8
		fmt.Print("Cows: ")
		fmt.Scan(&cows)
		fmt.Print("Bulls: ")
		fmt.Scan(&bulls)

		if len(guess) == int(bulls) {
			fmt.Printf("It took me %d tries to guess your number.\n", guesses)
			return
		}

		var next []bc.Number
		for _, entry := range remaining {
			if entry.Bulls(guess) == bulls && entry.Cows(guess) == cows {
				next = append(next, entry)
			}
		}

		fmt.Printf("I've narrowed it down to %d options:\n", len(next))
		remaining = next
		guesses++
	}
}
