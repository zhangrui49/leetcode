package g144

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (71.02%)
 * Likes:    820
 * Dislikes: 0
 * Total Accepted:    595.7K
 * Total Submissions: 838.8K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,2,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：root = [1,2]
 * 输出：[1,2]
 *
 *
 * 示例 5：
 *
 *
 * 输入：root = [1,null,2]
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	r := []int{}
	var f func(*TreeNode)
	f = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		r = append(r, tn.Val)
		f(tn.Left)
		f(tn.Right)

	}
	f(root)
	return r
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//go数组值传递
func preorder(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}
	*r = append(*r, root.Val)
	preorder(root.Left, r)
	preorder(root.Right, r)
}

//迭代法
func preorderTraversalV2(root *TreeNode) []int {
	r := []int{}
	stack := []*TreeNode{}
	tn := root

	for tn != nil || len(stack) > 0 {
		for tn != nil {
			r = append(r, tn.Val)
			stack = append(stack, tn)
			tn = tn.Left
		}
		tn = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return r
}
