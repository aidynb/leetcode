package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := &ListNode{}
	l1.PushFront([]int{1, 2, 4})
	l2 := &ListNode{}
	l2.PushFront([]int{1, 3, 4})

	l3 := &ListNode{}
	l4 := &ListNode{}

	l5 := &ListNode{}
	l6 := &ListNode{}
	l6.PushFront([]int{0})

	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want []int
	}{
		{l1, l2, []int{1, 1, 2, 3, 4, 4}},
		{l3, l4, []int{}},
		{l5, l6, []int{0}},
	}

	for _, tt := range tests {
		res := mergeTwoLists(tt.l1.Next, tt.l2.Next)

		got := res.Traverse()

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
