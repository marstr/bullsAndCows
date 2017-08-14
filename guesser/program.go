package main

import (
	"fmt"

	bc "github.com/marstr/bullsAndCows"
)

func main() {
	remaining := []bc.Number{}

	for i := 0; i < 10000; i++ {
		parsed, err := bc.ParseNumber(fmt.Sprint(i))
		if err != nil {
			remaining = append(remaining, parsed)
		}
	}

	for {
		guess := remaining[0]

		fmt.Println("My guess: ", guess)

	}
}
