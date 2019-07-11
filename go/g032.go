package main

//func longestValidParentheses(s string) int {
//	len_s := len(s)
//	temp := ""
//	for j := len_s; j > 0; j-- {
//		for i := 0; i <= len_s-j; i++ {
//			temp = s[i : i+j]
//			if isValid(temp) {
//				return len(temp)
//			}
//		}
//	}
//	return 0
//}

func longestValidParentheses(s string) int {
	bytes := []byte(s)
	if len(bytes) < 2 {
		return 0
	}
	lengthList := make([]int, len(bytes))
	var max int
	for i := 1; i < len(bytes); i++ {
		if bytes[i] == ')' {
			j := i - lengthList[i-1] - 1
			if j >= 0 && bytes[j] == '(' {
				lengthList[i] = lengthList[i-1] + 2
				if j-1 >= 0 {
					lengthList[i] += lengthList[j-1]
				}
			}
		}
		if lengthList[i] > max {
			max = lengthList[i]
		}
	}
	return max
}

func isValid(str string) bool {
	array := make([]byte, len(str)+1)
	index := 1
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			array[index] = str[i]
			index++
		} else {
			index--
			if array[index] != '(' {
				return false
			}
		}
	}
	return index == 1
}
