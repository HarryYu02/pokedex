package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error: Length not match")
			t.Errorf("\tExpected: %s (%d)", c.expected, len(c.expected))
			t.Errorf("\tActual: %s (%d)", actual, len(actual))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word[i] != expectedWord[i] {
				t.Errorf("Error: Word not match at index %d", i)
				return
			}
		}
	}
}
