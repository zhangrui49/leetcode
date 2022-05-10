package g557

import (
	"bytes"
	"strings"
)

/**

给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序

*/

//API大法
func reverseWords(s string) string {
	split := strings.Split(s, " ")
	var bt bytes.Buffer

	length := len(split)

	for i := 0; i < length; i++ {
		v := split[i]
		bytes := String2Bytes(v)
		reverseString(bytes)
		bt.WriteString(string(bytes))
		if i != length-1 {
			bt.WriteByte(' ')
		}

	}
	return bt.String()
}

//见g344
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

func String2Bytes(data string) []byte {
	return []byte(data)
}

//纯手打
func reverseWordsV2(s string) string {
	length := len(s)
	r := make([]byte, 0)
	start := 0
	for i := 0; i < length; i++ {
		c := s[i]

		if c == ' ' {

			for j := i - 1; j >= start; j-- {
				r = append(r, s[j])
			}
			r = append(r, ' ')
			start = i + 1
		}

		if i == length-1 {
			for j := i; j >= start; j-- {
				r = append(r, s[j])
			}
		}
	}
	return string(r)
}
