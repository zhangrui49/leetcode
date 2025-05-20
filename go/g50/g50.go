package g50

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	} else {
		return 1.0 / quickMul(x, -n)
	}
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}

// 老规矩，先暴力法 304/307 cases passed (N/A)
func myPowV1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1.0 / myPowV1(x, -n)
	}
	if x == 1.0 {
		return 1.0
	}

	result := 1.0
	for n > 0 {
		result *= x
		n--
	}
	return result
}

// @lc code=end
