/**
005最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
*/
package g005

import (
	"fmt"
)

func longestPalindrome(s string) string {
	lenStr := len(s)
	start, lenSub := 0, 0
	for i := 0; i < lenStr; i++ {
		println(i)
		tempStart1, tempLen1 := ripple(s, i, i)
		tempStart2, tempLen2 := ripple(s, i, i+1)
		if tempLen1 > 0 || tempLen2 > 0 {
			if tempLen1 > tempLen2 {
				if tempLen1 > lenSub {
					start = tempStart1
					lenSub = tempLen1
				}
			} else {

				if tempLen2 > lenSub {
					start = tempStart2
					lenSub = tempLen2
				}

			}
			fmt.Printf("len:%d start:%d \n", lenSub, start)
		}
		// if (lenStr-i)*2 < lenSub {
		// 	break
		// }

	}

	return s[start : start+lenSub]
}

func ripple(s string, left int, right int) (int, int) {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	subLen := right - left - 1
	if left == -1 {
		left = 0
	} else {
		left = left + 1
	}
	fmt.Printf("len:%d left:%d \n", subLen, left)
	return left, subLen
}
