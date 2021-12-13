package lc

import "testing"

func Test_BestTimeToBuyAndSellStock(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{1, 2}, 1},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 4, 1}, 2},
	}

	for _, tt := range tests {
		got := maxProfit(tt.prices)

		if got != tt.want {
			t.Errorf("got %v, want %v\n", got, tt.want)
		}
	}
}
