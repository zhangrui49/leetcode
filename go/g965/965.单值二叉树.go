/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
 *
 * https://leetcode.cn/problems/univalued-binary-tree/description/
 *
 * algorithms
 * Easy (68.75%)
 * Likes:    122
 * Dislikes: 0
 * Total Accepted:    46.3K
 * Total Submissions: 66.1K
 * Testcase Example:  '[1,1,1,1,1,null,1]'
 *
 * 如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
 *
 * 只有给定的树是单值二叉树时，才返回 true；否则返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[1,1,1,1,1,null,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：[2,2,2,5,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定树的节点数范围是 [1, 100]。
 * 每个节点的值都是整数，范围为 [0, 99] 。
 *
 *
 */
package g965

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	value := root.Val
	var r = true
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.Val != value {
			r = false
			return
		}
		if r {
			dfs(tn.Left)
			dfs(tn.Right)
		}

	}
	dfs(root)
	return r
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
