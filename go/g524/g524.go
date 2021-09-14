/*

给你一个字符串 s 和一个字符串数组 dictionary 作为字典，
找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串

*/

package g524

//通过删除字母匹配到字典里最长单词

func findLongestWord(s string, dictionary []string) string {
	max := 0
	result := ""
	for _, element := range dictionary {
		isChild := isChild(element, s)
		if isChild {
			size := len(element)
			//这里有点迷,字典序最小的,竟然不是字典数组的下标最小,是字符串的大小
			if size > max {
				max = size
				result = element
			}
			//比较字符串的大小
			if size == max {
				if element < result {
					result = element
				}

			}

		}
	}
	return result
}

func isChild(s, parent string) bool {
	sizeP := len(parent)
	sizeC := len(s)
	if sizeC > sizeP {
		return false
	}

	//子序列下标
	index := 0

	for i := 0; i < sizeP; i++ {
		charP := parent[i]
		charC := s[index]
		if charC == charP {
			index++
			if index == sizeC {
				return true
			}
		}

	}

	return false
}
