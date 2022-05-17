/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 *
 * https://leetcode.cn/problems/verifying-an-alien-dictionary/description/
 *
 * algorithms
 * Easy (55.37%)
 * Likes:    144
 * Dislikes: 0
 * Total Accepted:    23.9K
 * Total Submissions: 42.2K
 * Testcase Example:  '["hello","leetcode"]\n"hlabcdefgijkmnopqrstuvwxyz"'
 *
 * 某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
 *
 * 给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回
 * false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
 * 输出：true
 * 解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。
 *
 * 示例 2：
 *
 *
 * 输入：words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
 * 输出：false
 * 解释：在该语言的字母表中，'d' 位于 'l' 之后，那么 words[0] > words[1]，因此单词序列不是按字典序排列的。
 *
 * 示例 3：
 *
 *
 * 输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
 * 输出：false
 * 解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"，因为 'l' > '∅'，其中
 * '∅' 是空白字符，定义为比任何其他字符都小（更多信息）。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * order.length == 26
 * 在 words[i] 和 order 中的所有字符都是英文小写字母。
 *
 *
 */
package g953

import (
	"strings"
)

// @lc code=start

var orderMap map[byte]int //key: 字符 value: 下标

func isAlienSorted(words []string, order string) bool {
	length := len(words)
	orderMap = make(map[byte]int)

outer:
	for i := 1; i < length; i++ {
		front := words[i-1]
		after := words[i]
		lf := len(front)
		la := len(after)

		var innerLen int
		if lf > la {
			innerLen = la
		} else {
			innerLen = lf
		}

		for j := 0; j < innerLen; j++ {
			cf := front[j]
			ca := after[j]
			vf := getIndex(cf, order)
			va := getIndex(ca, order)
			if vf < va {
				continue outer
			} else if vf > va {
				return false
			}
		}

		if lf > la {
			return false
		}
	}
	return true
}

func getIndex(c byte, order string) int {
	if v, ok := orderMap[c]; ok {
		return v
	} else {
		p := strings.IndexByte(order, c)
		orderMap[c] = p
		return p
	}
}

// @lc code=end
