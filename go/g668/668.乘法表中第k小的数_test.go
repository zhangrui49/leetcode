/*
 * @lc app=leetcode.cn id=668 lang=golang
 *
 * [668] 乘法表中第k小的数
 *
 * https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/description/
 *
 * algorithms
 * Hard (52.11%)
 * Likes:    264
 * Dislikes: 0
 * Total Accepted:    18.2K
 * Total Submissions: 32.3K
 * Testcase Example:  '3\n3\n5'
 *
 * 几乎每一个人都用 乘法表。但是你能在乘法表中快速找到第k小的数字吗？
 *
 * 给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。
 *
 * 例 1：
 *
 *
 * 输入: m = 3, n = 3, k = 5
 * 输出: 3
 * 解释:
 * 乘法表:
 * 1    2    3
 * 2    4    6
 * 3    6    9
 *
 * 第5小的数字是 3 (1, 2, 2, 3, 3).
 *
 *
 * 例 2：
 *
 *
 * 输入: m = 2, n = 3, k = 6
 * 输出: 6
 * 解释:
 * 乘法表:
 * 1    2    3
 * 2    4    6
 *
 * 第6小的数字是 6 (1, 2, 2, 3, 4, 6).
 *
 *
 * 注意：
 *
 *
 * m 和 n 的范围在 [1, 30000] 之间。
 * k 的范围在 [1, m * n] 之间。
 *
 *
 */

package g668

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				m: 3,
				n: 3,
				k: 6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthNumberV2(tt.args.m, tt.args.n, tt.args.k)
			t.Errorf("findKthNumber() = %v, want %v", got, tt.want)

		})
	}
}
