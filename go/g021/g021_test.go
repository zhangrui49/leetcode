package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{}
	list1.Val = 1
	list1.Next = &ListNode{4, &ListNode{5, nil}}
	list2 := &ListNode{}
	list2.Val = 1
	list2.Next = &ListNode{3, &ListNode{4, nil}}

	result := mergeTwoListsV2(list1, list2)
	next := result
	for {
		if next == nil {
			break
		}
		t.Error(next.Val)
		next = next.Next
	}
}
