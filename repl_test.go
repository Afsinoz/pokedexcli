package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello   world  ",
			expected:[]string{"hello","world"},
		},
		{
			input:" nabersin cinim ",
			expected:[]string{"nabersin","cinim"},
		},
		{
			input:" buralara kar yagiyor cinim",
			expected:[]string{"buralara","kar","yagiyor","cinim"},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length is wrong!")
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words are not maching")
				return
			}
		}

	}
}

