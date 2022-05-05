package main

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	list := &ListNode{}
	list.Val = 1
	list.Next = &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	//list.Next = &ListNode{2, nil}
	result := reverseKGroup(list, 2)

	next := result
	for {
		if next == nil {
			break
		}
		t.Log(next.Val)
		next = next.Next
	}
}
