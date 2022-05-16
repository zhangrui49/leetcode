/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] 序列化和反序列化二叉搜索树
 *
 * https://leetcode.cn/problems/serialize-and-deserialize-bst/description/
 *
 * algorithms
 * Medium (56.63%)
 * Likes:    356
 * Dislikes: 0
 * Total Accepted:    40.7K
 * Total Submissions: 67.1K
 * Testcase Example:  '[2,1,3]'
 *
 * 序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
 *
 * 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。
 * 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
 *
 * 编码的字符串应尽可能紧凑。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：[2,1,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数范围是 [0, 10^4]
 * 0 <= Node.val <= 10^4
 * 题目数据 保证 输入的树是一棵二叉搜索树。
 *
 *
 */
package g449

import (
	"math"
	"strconv"
	"strings"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	//前序遍历
	array := []string{}
	//var s string
	var f func(*TreeNode)
	f = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		//s = s + strconv.Itoa(tn.Val) + " "
		array = append(array, strconv.Itoa(tn.Val))
		f(tn.Left)
		f(tn.Right)

	}
	f(root)
	return strings.Join(array, " ")
	//return s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	splits := strings.Split(data, " ")
	var f func(int, int) *TreeNode
	f = func(left, right int) *TreeNode {
		if len(splits) == 0 {
			return nil
		}
		v, _ := strconv.Atoi(splits[0])

		if v < left || v > right {
			return nil
		}
		splits = splits[1:]

		return &TreeNode{Val: v, Left: f(left, v), Right: f(v, right)}
	}
	return f(math.MinInt, math.MaxInt)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
