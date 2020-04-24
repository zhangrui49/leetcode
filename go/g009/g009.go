package g009

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x > 0) {
		return false
	}
	var temp int = 0
	for x > temp {
		temp = temp*10 + x%10
		x /= 10
	}
	return temp == x || temp/10 == x
	//return IsPalindrome(strconv.Itoa(x))
}

func IsPalindrome(s string) bool {
	size := len(s)

	if size == 0 {
		return false
	}

	left, right, mid := 0, 0, 0

	if size%2 == 0 {
		mid = size / 2
		left = mid - 1
		right = mid
	} else {
		mid = (size - 1) / 2
		left = mid - 1
		right = mid + 1
	}

	for left >= 0 && right < size && s[left] == s[right] {
		left--
		right++
	}
	if right == size {
		return true
	}
	return false
}
