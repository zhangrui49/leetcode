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
