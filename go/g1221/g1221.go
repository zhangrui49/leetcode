/*
在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。
给你一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。
注意：分割得到的每个字符串都必须是平衡字符串。
返回可以通过分割得到的平衡字符串的 最大数量
*/

package g1221

//分割平衡字符串
func balancedStringSplit(s string) int {
	leftCount, rightCount, total := 0, 0, 0
	L, R := 'L', 'R'

	for _, v := range s {
		if L == v {
			leftCount++
		}
		if R == v {
			rightCount++
		}

		if leftCount == rightCount {
			total++
			leftCount = 0
			rightCount = 0
		}
	}
	return total
}
