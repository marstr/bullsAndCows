package bullsAndCows

import (
	"bytes"
	"fmt"
)

// Number represents a valid guess in the game of Bulls and Cows
type Number [4]uint8

// IsValid ensures there are no repeated digits in the Number.
func (n Number) IsValid() bool {
	seen := make(map[uint8]struct{})

	for i := 0; i < len(n); i++ {
		current := n[i]
		if current >= 10 {
			return false
		}

		if _, ok := seen[current]; ok {
			return false
		}

		seen[current] = struct{}{}
	}

	return true
}

// Bulls counts how many digits are present and in the same location.
func (n Number) Bulls(other Number) (count uint8) {
	for i := 0; i < len(n); i++ {
		if n[i] == other[i] {
			count++
		}
	}
	return
}

// Cows counts how many digits are present in another number, but in a different place.
func (n Number) Cows(other Number) (count uint8) {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			if i == j {
				continue
			}

			if n[i] == other[j] {
				count++
			}
		}
	}
	return
}

// ParseNumber will read a number out of a guess
func ParseNumber(raw string) (result Number, err error) {
	var x int

	defer func() {
		if err != nil {
			result = Number{}
		}
	}()

	_, err = fmt.Sscan(raw, &x)
	if err != nil {
		return
	}

	for i := len(result) - 1; i >= 0; i-- {
		result[i] = uint8(x % 10)
		x /= 10
	}

	if x != 0 {
		err = fmt.Errorf("%s contains too many digits. Expected: %d", raw, len(result))
		return
	}

	if !result.IsValid() {
		err = fmt.Errorf("%s contains a duplicate digit", raw)
		return
	}

	return
}

func (n Number) String() string {
	results := &bytes.Buffer{}

	for i := 0; i < len(n); i++ {
		fmt.Fprint(results, n[i])
	}
	return results.String()
}
