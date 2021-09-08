package goffer10

//剑指 Offer 10- I. 斐波那契数列

var resultMap map[int]int = make(map[int]int)

func fib(n int) int {
	if n < 2 {
		return n
	}
	if v, ok := resultMap[n]; ok {
		return v
	}
	result := fib(n-1) + fib(n-2)
	resultMap[n] = result
	return result % 1000000007
}

var step1, step2, result = 0, 0, 1

func fib2(n int) int {
	if n < 2 {
		return n
	}
	for i := 2; i <= n; i++ {
		step1 = step2
		step2 = result
		result = (step1 + step2) % 1000000007
	}
	return result
}
