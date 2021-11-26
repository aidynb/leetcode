package lc

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
		{[]int{1}, 2, 1},
	}

	for _, tt := range tests {
		got := searchInsertPosition(tt.nums, tt.target)

		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
