package g160

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pt1 := headA
	pt2 := headB
	exchanged := false
	for {
		if pt1 == pt2 {
			return pt1
		}
		pt1 = pt1.Next
		if pt1 == nil {
			if !exchanged {
				pt1 = headB
				exchanged = true
			} else {
				return nil
			}
		}
		pt2 = pt2.Next
		if pt2 == nil {
			pt2 = headA
		}

	}
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}
