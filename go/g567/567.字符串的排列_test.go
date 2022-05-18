/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode.cn/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (43.85%)
 * Likes:    675
 * Dislikes: 0
 * Total Accepted:    188.9K
 * Total Submissions: 430.5K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 *
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 *
 *
 */

package g567

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.args.s1, tt.args.s2)
			t.Errorf("checkInclusion() = %v, want %v", got, tt.want)

		})
	}
}
