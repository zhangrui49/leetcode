/*
	给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
	你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
	要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
	文本的最后一行应为左对齐，且单词之间不插入额外的空格。

	说明:
	单词是指由非空格字符组成的字符序列。
	每个单词的长度大于 0，小于等于 maxWidth。
	输入单词数组 words 至少包含一个单词
*/

package g068

//文本左右对齐
func fullJustify(words []string, maxWidth int) []string {
	size := len(words)
	result := []string{""}
	index := 0
	for {
		sumLen := 0
		start := index
		//当前单次长度+空格个数
		for start < size && sumLen+len(words[start])+start-index <= maxWidth {
			sumLen += len(words[start])
			start++
		}
		

	}
	return result
}
