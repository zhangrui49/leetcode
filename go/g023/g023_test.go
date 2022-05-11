package main

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	list1 := &ListNode{}
	list1.Val = 1
	list1.Next = &ListNode{4, &ListNode{5, nil}}
	list2 := &ListNode{}
	list2.Val = 1
	list2.Next = &ListNode{3, &ListNode{4, nil}}
	list3 := &ListNode{}
	list3.Val = 2
	list3.Next = &ListNode{6, nil}
	result := mergeKLists([]*ListNode{list1, list2, list3})
	next := result
	for {
		if next == nil {
			break
		}
		t.Log(next.Val)
		next = next.Next
	}
}
