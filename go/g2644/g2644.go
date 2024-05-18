package g2644

func maxDivScore(nums []int, divisors []int) int {
	maxCount := 0
	minValue := 0
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		count := 0
		for _, number := range nums {
			if number%divisor == 0 {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			minValue = divisor
		} else if count == maxCount {
			if divisor < minValue || minValue == 0 {
				minValue = divisor
			}
		}
	}

	return minValue
}
