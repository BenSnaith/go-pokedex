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
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		// add more cases
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("lengths do not match: '%v' vs. '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, expectedWord)
			}
		}
	}
}
