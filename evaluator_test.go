package bullsAndCows

import (
	"errors"
	"fmt"
	"testing"
)

func Test_ParseNumber(t *testing.T) {
	cases := []struct {
		string
		expected Number
		err      error
	}{
		{"1234", Number{1, 2, 3, 4}, nil},
		{"1123", Number{}, errors.New("1123 contains a duplicate digit")},
		{"8675309", Number{}, errors.New("8675309 contains too many digits. Expected: 4")},
		{"0497", Number{0, 4, 9, 7}, nil},
	}

	for _, tc := range cases {
		t.Run(tc.string, func(t *testing.T) {
			actual, actualErr := ParseNumber(tc.string)

			if actual != tc.expected {
				t.Logf("got: %q want: %q", actual, tc.expected)
				t.Fail()
			}

			if tc.err != nil {
				if actualErr.Error() != tc.err.Error() {
					t.Logf("got: %v want: %v", actualErr, tc.err)
					t.Fail()
				}
			} else if actualErr != nil {
				t.Error(actualErr)
			}
		})
	}
}

func TestNumber_Cows(t *testing.T) {
	cases := []struct {
		left     Number
		right    Number
		expected uint8
	}{
		{Number{1, 2, 3, 4}, Number{1, 2, 3, 4}, 0},
		{Number{1, 2, 3, 4}, Number{5, 6, 7, 8}, 0},
		{Number{5, 6, 1, 2}, Number{5, 1, 6, 2}, 2},
		{Number{1, 2, 5, 6}, Number{9, 7, 2, 1}, 2},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v %v", tc.left, tc.right), func(t *testing.T) {
			if cows := tc.left.Cows(tc.right); cows != tc.expected {
				t.Logf("got: %d want: %d", cows, tc.expected)
				t.Fail()
			}
		})
	}
}

func TestNumber_IsValid(t *testing.T) {
	cases := []struct {
		Number
		bool
	}{
		{Number{0, 9, 4, 7}, true},
		{Number{1, 2, 3, 4}, true},
		{Number{0, 0, 8, 9}, false},
		{Number{9, 0, 9, 1}, false},
		{Number{0, 4, 9, 7}, true},
	}

	for _, tc := range cases {
		t.Run(tc.Number.String(), func(t *testing.T) {
			if got := tc.Number.IsValid(); got != tc.bool {
				t.Fail()
			}
		})
	}
}
