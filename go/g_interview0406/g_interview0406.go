package ginterview0406

/**

设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//中序遍历，在中间写判断逻辑即可
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	var r *TreeNode
	var prev *TreeNode
	var curr *TreeNode
	var f func(*TreeNode)
	f = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		f(tn.Left)

		if curr != nil {
			prev = curr
		}
		curr = tn
		if prev == p {
			r = curr
			return
		}
		if curr == nil {
			return
		}

		f(tn.Right)

	}
	f(root)
	return r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
