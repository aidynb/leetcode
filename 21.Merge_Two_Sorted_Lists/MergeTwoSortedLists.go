package lc

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PushFront(nums []int) {
	if l == nil || len(nums) == 0 {
		return
	}

	for l.Next != nil {
		l = l.Next
	}

	for _, num := range nums {
		newNode := &ListNode{Val: num}
		l.Next = newNode
		l = l.Next
	}
}

func (l *ListNode) Traverse() []int {
	result := []int{}

	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	trav := newList

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			trav.Next = l1
			l1 = l1.Next
		} else {
			trav.Next = l2
			l2 = l2.Next
		}
		trav = trav.Next
	}

	if l1 != nil {
		trav.Next = l1
	} else if l2 != nil {
		trav.Next = l2
	}

	return newList.Next
}
