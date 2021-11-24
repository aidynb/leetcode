package lc

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"(]",
			false,
		},
		{
			"([)]",
			false,
		},
		{
			"{[]}",
			true,
		},
	}

	for _, tt := range tests {
		got := isValid(tt.s)

		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
