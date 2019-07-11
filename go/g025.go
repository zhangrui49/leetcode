package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	index := 1
	nodeMap := make(map[int]*ListNode)
	next := head
	for {
		if next == nil {
			break
		}
		nodeMap[index] = next
		fmt.Printf("index:%d val: %d  \n", index, next.Val)
		next = next.Next
		index++
	}
	total := len(nodeMap)
	fmt.Println(total)
	for i := 1; i <= total; i++ {
		if i%k == 0 {
			nodeMap[i-k+1].Next = nodeMap[i+1]

			if i-k == 0 {
				head = nodeMap[i]
			} else {
				nodeMap[i-k*2+1].Next = nodeMap[i]
			}
			for j := 0; j < k-1; j++ {
				nodeMap[i-j].Next = nodeMap[i-j-1]
			}
		}
	}
	return head
}
