package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := make(map[int]*ListNode)
	count := 0
	next := head
	for {
		temp[count] = next
		if next == nil {
			if count-n == 0 {
				return head.Next
			}
			prv := temp[count-n-1]
			prv.Next = temp[count-n+1]
			break
		}
		next = next.Next
		count++
	}
	return head
}

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	temp := make([]*ListNode, 0)
	next := head
	for next != nil {
		temp = append(temp, next)
		next = next.Next
	}
	length := len(temp)

	if length == n {
		return head.Next
	}

	prev := temp[len(temp)-n-1]

	if length-n == length-1 {
		prev.Next = nil
	} else {
		prev.Next = temp[len(temp)-n+1]
	}
	return head

}

func removeNthFromEndV3(head *ListNode, n int) *ListNode {
	temp := make([]*ListNode, 0)
	dummy := &ListNode{0, head}
	// next := dummy
	// for next != nil {
	// 	temp = append(temp, next)
	// 	next = next.Next
	// }

	for node := dummy; node != nil; node = node.Next {
		temp = append(temp, node)
	}

	prev := temp[len(temp)-n-1]
	prev.Next = prev.Next.Next

	return dummy.Next
}

func removeNthFromEndV4(head *ListNode, n int) *ListNode {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	dummy := &ListNode{0, head}
	curr := dummy
	for i := 0; i < length-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
