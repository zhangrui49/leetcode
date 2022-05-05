package g70

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

var temp = make(map[int]int) //避免重复计算

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	one, two := 0, 0
	if v, ok := temp[n-1]; ok {
		one = v
	} else {
		one = climbStairs(n - 1)
	}
	if v, ok := temp[n-2]; ok {
		two = v
	} else {
		two = climbStairs(n - 2)
	}

	temp[n] = one + two
	return one + two

}

func climbStairsV2(n int) int {
	p1, p2, r := 0, 0, 1

	for i := 1; i <= n; i++ {
		p2 = p1
		p1 = r
		r = p1 + p2
	}

	return r
}

func climbStairsV3(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	temp := make([]int, n)
	temp[0] = 1
	temp[1] = 2

	for i := 2; i < n; i++ {
		temp[i] = temp[i-2] + temp[i-1]
	}

	return temp[n-1]
}
