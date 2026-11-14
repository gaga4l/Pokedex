package main

import (
	"testing"

	"github.com/gaga4l/pokedex/internal/config"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hi there     i'm waiting      for you    ",
			expected: []string{"hi", "there", "i'm", "waiting", "for", "you"},
		},
	}

	for _, c := range cases {
		actual := config.CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}

	}
}
