package main

import (
	"fmt"
	"math"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// func main() {
// 	list1 := &ListNode{}
// 	list1.Val = 1
// 	list1.Next = &ListNode{4, &ListNode{5, nil}}
// 	list2 := &ListNode{}
// 	list2.Val = 1
// 	list2.Next = &ListNode{3, &ListNode{4, nil}}
// 	list3 := &ListNode{}
// 	list3.Val = 2
// 	list3.Next = &ListNode{6, nil}
// 	result := mergeKLists([]*ListNode{list1, list2, list3})
// 	next := result
// 	for {
// 		if next == nil {
// 			break
// 		}
// 		fmt.Println(next.Val)
// 		next = next.Next
// 	}
// }

/**
题目其实先用一个数组存所有的链表Val,然后对数组排序，再返回一个新的链表，最简单
*/

/*
 @see leecode 21
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	node := &ListNode{}
	if l1.Val <= l2.Val {
		node = l1
		node.Next = mergeTwoLists(l1.Next, l2)
	} else if l1.Val > l2.Val {
		node = l2
		node.Next = mergeTwoLists(l1, l2.Next)
	}
	return node
}

func mergeKListsBetter(lists []*ListNode) *ListNode {
	node := &ListNode{}
	if lists == nil {
		return node
	}
	size := len(lists)
	if size == 0 {
		return node
	}

	if size == 1 {
		return lists[0]
	}

	mid := size / 2
	//二分法 递归
	left := mergeKListsBetter(lists[:mid])
	right := mergeKListsBetter(lists[mid:])
	// 最终执行两两合并
	return mergeTwoLists(left, right)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	size := len(lists)
	if size == 0 {
		return nil
	}
	var first *ListNode = nil
	var now *ListNode = nil
	indexMap := make(map[int]*ListNode)
	for i := 0; i < size; i++ {
		if lists[i] != nil {
			indexMap[i] = lists[i]
		}
	}
	for {
		min := math.MaxInt32
		innerMap := make(map[int]int)

		if len(indexMap) == 0 {
			break
		}

		for k, value := range indexMap {
			if value != nil {
				v := value.Val
				innerMap[v] = k
				if v < min {
					min = v
				}
			}
		}

		index := innerMap[min]
		indexNode := indexMap[index]
		if now == nil {
			now = indexNode
			first = indexNode
		} else {
			now.Next = indexNode
			now = indexNode
		}
		if indexNode != nil {
			if indexNode.Next != nil {
				indexMap[index] = indexNode.Next
			} else {
				delete(indexMap, index)
			}
		}

	}
	return first
}
